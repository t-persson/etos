// Copyright Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package provider

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"testing"
	"time"

	"github.com/eiffel-community/etos/api/v1alpha1"
	"github.com/eiffel-community/etos/api/v1alpha2"
	"go.jetify.com/sse"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	tracenoop "go.opentelemetry.io/otel/trace/noop"
)

// newTestMeterProvider returns a meter provider backed by a manual reader.
func newTestMeterProvider(t *testing.T) (*sdkmetric.MeterProvider, *sdkmetric.ManualReader) {
	t.Helper()
	reader := sdkmetric.NewManualReader()
	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
	t.Cleanup(func() { _ = meterProvider.Shutdown(context.Background()) })
	return meterProvider, reader
}

// useGlobalMeterProvider installs a manual-reader meter provider as the global meter provider
// for the duration of the test.
func useGlobalMeterProvider(t *testing.T) *sdkmetric.ManualReader {
	t.Helper()
	meterProvider, reader := newTestMeterProvider(t)
	previous := otel.GetMeterProvider()
	otel.SetMeterProvider(meterProvider)
	t.Cleanup(func() { otel.SetMeterProvider(previous) })
	return reader
}

// collect returns the metrics with the given names, keyed by name.
func collect(t *testing.T, reader *sdkmetric.ManualReader) map[string]metricdata.Metrics {
	t.Helper()
	var resourceMetrics metricdata.ResourceMetrics
	if err := reader.Collect(context.Background(), &resourceMetrics); err != nil {
		t.Fatalf("Collect() error = %v", err)
	}
	metrics := map[string]metricdata.Metrics{}
	for _, scope := range resourceMetrics.ScopeMetrics {
		if scope.Scope.Name != meterName {
			t.Errorf("scope = %q, want %q", scope.Scope.Name, meterName)
		}
		for _, m := range scope.Metrics {
			metrics[m.Name] = m
		}
	}
	return metrics
}

// attributeMap converts an attribute set to a map for comparison.
func attributeMap(set attribute.Set) map[string]string {
	result := map[string]string{}
	for _, kv := range set.ToSlice() {
		result[string(kv.Key)] = kv.Value.String()
	}
	return result
}

// counterValues returns the counter data points keyed by their outcome, verifying their attributes.
func counterValues(
	t *testing.T, metrics map[string]metricdata.Metrics, name, unit string, attrs map[string]string,
) map[string]int64 {
	t.Helper()
	m, ok := metrics[name]
	if !ok {
		t.Fatalf("metric %q not recorded", name)
	}
	if m.Unit != unit {
		t.Errorf("%s unit = %q, want %q", name, m.Unit, unit)
	}
	sum, ok := m.Data.(metricdata.Sum[int64])
	if !ok || !sum.IsMonotonic {
		t.Fatalf("%s is not a monotonic int64 counter: %T", name, m.Data)
	}
	values := map[string]int64{}
	for _, point := range sum.DataPoints {
		got := attributeMap(point.Attributes)
		outcome := got[string(outcomeKey)]
		delete(got, string(outcomeKey))
		if fmt.Sprint(got) != fmt.Sprint(attrs) {
			t.Errorf("%s attributes = %v, want %v plus outcome", name, got, attrs)
		}
		values[outcome] = point.Value
	}
	return values
}

// histogramPoint returns the single histogram data point of a metric.
func histogramPoint(
	t *testing.T, metrics map[string]metricdata.Metrics, name string,
) metricdata.HistogramDataPoint[float64] {
	t.Helper()
	m, ok := metrics[name]
	if !ok {
		t.Fatalf("metric %q not recorded", name)
	}
	if m.Unit != secondsUnit {
		t.Errorf("%s unit = %q, want %q", name, m.Unit, secondsUnit)
	}
	histogram, ok := m.Data.(metricdata.Histogram[float64])
	if !ok {
		t.Fatalf("%s is not a float64 histogram: %T", name, m.Data)
	}
	if len(histogram.DataPoints) != 1 {
		t.Fatalf("%s has %d data points, want 1", name, len(histogram.DataPoints))
	}
	if !slices.Equal(histogram.DataPoints[0].Bounds, durationBuckets) {
		t.Errorf("%s bounds = %v, want %v", name, histogram.DataPoints[0].Bounds, durationBuckets)
	}
	return histogram.DataPoints[0]
}

// TestProviderAttemptRecordsOneAttemptPerOutcome verifies that each provider type records exactly
// one provisioning or release attempt with bounded attributes, a zero for the other outcomes, the
// duration in seconds, and the failed stage on unsuccessful attempts only.
func TestProviderAttemptRecordsOneAttemptPerOutcome(t *testing.T) {
	tests := []struct {
		name          string
		providerType  string
		wantType      string
		release       bool
		schemaVersion string
		stage         string
		err           error
		wantOutcome   string
		wantStage     string
	}{
		{
			name: "iut provisioning success", providerType: "Iut", wantType: "iut",
			schemaVersion: "v1beta1", stage: stageOperation, wantOutcome: outcomeSuccess,
		},
		{
			name: "log area provisioning failure", providerType: "LogArea", wantType: "log_area",
			schemaVersion: "v1beta1", stage: stageOperation, err: errors.New("raw error iut-name"),
			wantOutcome: outcomeFailure, wantStage: stageOperation,
		},
		{
			name: "execution space provisioning timeout", providerType: "ExecutionSpace",
			wantType: "execution_space", schemaVersion: "v1alpha1", stage: stageOperation,
			err:         fmt.Errorf("wait: %w", context.DeadlineExceeded),
			wantOutcome: outcomeTimeout, wantStage: stageOperation,
		},
		{
			name: "startup failure before the EnvironmentRequest is known", providerType: "Iut",
			wantType: "iut", stage: stageStartup, err: errors.New("not found"),
			wantOutcome: outcomeFailure, wantStage: stageStartup,
		},
		{
			name: "termination-log failure", providerType: "LogArea", wantType: "log_area",
			schemaVersion: "v1beta1", stage: stageOperation,
			err:         &terminationLogError{err: errors.New("read-only file system")},
			wantOutcome: outcomeFailure, wantStage: stageResult,
		},
		{
			name: "iut release success", providerType: "Iut", wantType: "iut", release: true,
			schemaVersion: "v1beta1", stage: stageOperation, wantOutcome: outcomeSuccess,
		},
		{
			name:         "execution space release failure with unbounded schema version",
			providerType: "ExecutionSpace", wantType: "execution_space", release: true,
			schemaVersion: "my-custom-version", stage: stageOperation, err: errors.New("forbidden"),
			wantOutcome: outcomeFailure, wantStage: stageOperation,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			meterProvider, reader := newTestMeterProvider(t)
			start := time.Unix(1000, 0)
			attempt := newProviderAttempt(meterProvider, tt.providerType, tt.release, start)
			attempt.now = func() time.Time { return start.Add(1500 * time.Millisecond) }
			if tt.schemaVersion != "" {
				attempt.setAPIVersion(tt.schemaVersion)
			}
			attempt.setStage(tt.stage)
			err := tt.err
			attempt.finish(&err)

			wantAPIVersion := tt.schemaVersion
			if wantAPIVersion != "v1alpha1" && wantAPIVersion != "v1beta1" {
				wantAPIVersion = unknownAttributeValue
			}
			attempts, duration := provisioningAttemptsName, provisioningDurationName
			otherAttempts, otherDuration := releaseAttemptsName, releaseDurationName
			if tt.release {
				attempts, otherAttempts = otherAttempts, attempts
				duration, otherDuration = otherDuration, duration
			}
			metrics := collect(t, reader)
			for _, name := range []string{otherAttempts, otherDuration} {
				if _, ok := metrics[name]; ok {
					t.Errorf("metric %q recorded, want only the %s operation", name, attempts)
				}
			}
			base := map[string]string{
				string(providerTypeKey): tt.wantType,
				string(apiVersionKey):   wantAPIVersion,
			}
			values := counterValues(t, metrics, attempts, attemptUnit, base)
			for _, outcome := range outcomes {
				want := int64(0)
				if outcome == tt.wantOutcome {
					want = 1
				}
				if got, ok := values[outcome]; !ok || got != want {
					t.Errorf("%s{outcome=%s} = %d (present %t), want %d", attempts, outcome, got, ok, want)
				}
			}
			point := histogramPoint(t, metrics, duration)
			if point.Count != 1 || point.Sum != 1.5 {
				t.Errorf("%s count = %d sum = %v, want 1 and 1.5", duration, point.Count, point.Sum)
			}
			wantAttrs := map[string]string{
				string(providerTypeKey): tt.wantType,
				string(apiVersionKey):   wantAPIVersion,
				string(outcomeKey):      tt.wantOutcome,
			}
			if tt.wantStage != "" {
				wantAttrs[string(stageKey)] = tt.wantStage
			}
			if got := attributeMap(point.Attributes); fmt.Sprint(got) != fmt.Sprint(wantAttrs) {
				t.Errorf("%s attributes = %v, want %v", duration, got, wantAttrs)
			}
		})
	}
}

// TestProviderAttemptRecordsPanicAsFailure verifies that a panicking provider records one failed
// attempt and that the panic continues to propagate.
func TestProviderAttemptRecordsPanicAsFailure(t *testing.T) {
	meterProvider, reader := newTestMeterProvider(t)
	attempt := newProviderAttempt(meterProvider, "Iut", false, time.Now())
	attempt.setAPIVersion("v1beta1")
	attempt.setStage(stageOperation)

	panicking := func() (err error) {
		defer attempt.finish(&err)
		panic("provider bug")
	}
	func() {
		defer func() {
			if recovered := recover(); recovered != "provider bug" {
				t.Errorf("recovered = %v, want the original panic", recovered)
			}
		}()
		_ = panicking()
	}()

	metrics := collect(t, reader)
	values := counterValues(t, metrics, provisioningAttemptsName, attemptUnit, map[string]string{
		string(providerTypeKey): "iut",
		string(apiVersionKey):   "v1beta1",
	})
	if values[outcomeFailure] != 1 || values[outcomeSuccess] != 0 {
		t.Errorf("attempts = %v, want one failure", values)
	}
	point := histogramPoint(t, metrics, provisioningDurationName)
	if got := attributeMap(point.Attributes)[string(stageKey)]; got != stageOperation {
		t.Errorf("stage = %q, want %q", got, stageOperation)
	}
}

// TestProviderAttemptSkipsUninstrumentedProviders verifies that the Environment Provider, which
// records its own metrics, records nothing here and that a nil attempt still propagates panics.
func TestProviderAttemptSkipsUninstrumentedProviders(t *testing.T) {
	meterProvider, reader := newTestMeterProvider(t)
	attempt := newProviderAttempt(meterProvider, "Environment", false, time.Now())
	if attempt != nil {
		t.Fatalf("newProviderAttempt(Environment) = %v, want nil", attempt)
	}
	attempt.setAPIVersion("v1beta1")
	attempt.setStage(stageOperation)
	err := errors.New("failure")
	attempt.finish(&err)

	func() {
		defer func() {
			if recover() == nil {
				t.Error("panic was swallowed by a nil attempt")
			}
		}()
		func() {
			defer attempt.finish(&err)
			panic("environment provider bug")
		}()
	}()
	if metrics := collect(t, reader); len(metrics) != 0 {
		t.Errorf("recorded metrics %v, want none", metrics)
	}
}

// TestWriteTerminationLogMarksResultFailures verifies that a failure to write a successful result
// is marked for the result stage while keeping its message.
func TestWriteTerminationLogMarksResultFailures(t *testing.T) {
	t.Setenv("KUBERNETES_SERVICE_HOST", "kubernetes.default")
	previous := terminationLog
	terminationLog = filepath.Join(t.TempDir(), "missing", "termination-log")
	t.Cleanup(func() { terminationLog = previous })

	params := Parameters{providerType: "Iut", tracer: tracenoop.NewTracerProvider().Tracer("test")}
	succeed := func(context.Context, Provider, Parameters, *v1alpha1.EnvironmentRequest) error { return nil }
	err := writeTerminationLog(context.Background(), succeed, nil, params, &v1alpha1.EnvironmentRequest{})

	var resultErr *terminationLogError
	if !errors.As(err, &resultErr) {
		t.Fatalf("writeTerminationLog() error = %v, want a terminationLogError", err)
	}
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("writeTerminationLog() error = %v, want it to wrap the write error", err)
	}
}

// TestWaitForTestRunnerRecordsReadiness verifies the Test Runner readiness attempt metrics for each
// normalized outcome and error type, without names, instances, or messages as attributes.
func TestWaitForTestRunnerRecordsReadiness(t *testing.T) {
	tests := []struct {
		name          string
		payload       string
		event         string
		block         bool
		cancel        bool
		timeout       time.Duration
		wantOutcome   string
		wantErrorType string
	}{
		{
			name:        "ready",
			event:       "status",
			payload:     `{"instance":"etr-instance","status":"running","message":"secret detail"}`,
			wantOutcome: outcomeSuccess,
		},
		{
			name:          "test runner error",
			event:         "status",
			payload:       `{"instance":"etr-instance","status":"error","message":"secret detail"}`,
			wantOutcome:   outcomeFailure,
			wantErrorType: errorTypeTestRunner,
		},
		{
			name:          "shutdown",
			event:         "shutdown",
			payload:       `{}`,
			wantOutcome:   outcomeFailure,
			wantErrorType: errorTypeShutdown,
		},
		{
			name:          "stream closed",
			wantOutcome:   outcomeFailure,
			wantErrorType: errorTypeStreamClosed,
		},
		{
			name:          "timeout",
			block:         true,
			timeout:       200 * time.Millisecond,
			wantOutcome:   outcomeTimeout,
			wantErrorType: errorTypeTimeout,
		},
		{
			name:          "canceled",
			block:         true,
			cancel:        true,
			wantOutcome:   outcomeFailure,
			wantErrorType: errorTypeCanceled,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := useGlobalMeterProvider(t)
			done := make(chan struct{})
			server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				writer.Header().Set("Content-Type", "text/event-stream")
				writer.(http.Flusher).Flush()
				if tt.block {
					select {
					case <-request.Context().Done():
					case <-done:
					}
					return
				}
				if tt.event == "" {
					return
				}
				if err := sse.NewEncoder(writer).EncodeEvent(&sse.Event{
					Event: tt.event,
					Data:  sse.Raw(tt.payload),
				}); err != nil {
					t.Errorf("encoding SSE event: %v", err)
				}
				writer.(http.Flusher).Flush()
			}))
			defer server.Close()
			defer close(done)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if tt.timeout > 0 {
				ctx, cancel = context.WithTimeout(ctx, tt.timeout)
				defer cancel()
			}
			if tt.cancel {
				time.AfterFunc(200*time.Millisecond, cancel)
			}

			executionSpace := &ExecutionSpace{ExecutionSpace: &v1alpha2.ExecutionSpace{
				Spec: v1alpha2.ExecutionSpaceSpec{
					Instructions: v1alpha2.Instructions{Environment: map[string]string{"ENVIRONMENT_ID": "etr-instance"}},
				},
			}}
			environmentRequest := &v1alpha1.EnvironmentRequest{
				Spec: v1alpha1.EnvironmentRequestSpec{
					SchemaVersion: "v1beta1",
					Identifier:    "test-identifier",
					Config:        v1alpha1.EnvironmentProviderJobConfig{EtosSse: server.URL},
				},
			}
			err := executionSpace.WaitForTestRunner(ctx, environmentRequest)
			if (err == nil) != (tt.wantOutcome == outcomeSuccess) {
				t.Fatalf("WaitForTestRunner() error = %v, want outcome %s", err, tt.wantOutcome)
			}

			metrics := collect(t, reader)
			values := counterValues(t, metrics, testRunnerReadinessAttName, attemptUnit, map[string]string{
				string(apiVersionKey): "v1beta1",
			})
			for _, outcome := range outcomes {
				want := int64(0)
				if outcome == tt.wantOutcome {
					want = 1
				}
				if values[outcome] != want {
					t.Errorf("readiness attempts{outcome=%s} = %d, want %d", outcome, values[outcome], want)
				}
			}
			point := histogramPoint(t, metrics, testRunnerReadinessDurName)
			if point.Count != 1 || point.Sum < 0 {
				t.Errorf("readiness duration count = %d sum = %v, want one non-negative value", point.Count, point.Sum)
			}
			wantAttrs := map[string]string{
				string(apiVersionKey): "v1beta1",
				string(outcomeKey):    tt.wantOutcome,
			}
			if tt.wantErrorType != "" {
				wantAttrs[string(errorTypeKey)] = tt.wantErrorType
			}
			if got := attributeMap(point.Attributes); fmt.Sprint(got) != fmt.Sprint(wantAttrs) {
				t.Errorf("readiness duration attributes = %v, want %v", got, wantAttrs)
			}
		})
	}
}
