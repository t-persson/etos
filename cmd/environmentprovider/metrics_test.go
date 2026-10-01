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

package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/eiffel-community/etos/api/v1alpha1"
	"github.com/eiffel-community/etos/api/v1alpha2"
	"github.com/eiffel-community/etos/pkg/provider"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	tracenoop "go.opentelemetry.io/otel/trace/noop"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"
)

const (
	testNamespace = "rt6t-test"
	testRequestID = "6a5b1c2d-3e4f-4a5b-8c6d-7e8f9a0b1c2d"
)

var (
	baseTime    = time.Date(2026, 10, 1, 12, 0, 0, 0, time.UTC)
	requestTime = baseTime.Add(-30 * time.Second)
)

// testClock returns a clock that starts at baseTime and advances 5 seconds on each read.
func testClock() func() time.Time {
	current := baseTime
	first := true
	return func() time.Time {
		if !first {
			current = current.Add(5 * time.Second)
		}
		first = false
		return current
	}
}

// providerResources returns the given number of IUT, LogArea and ExecutionSpace resources.
func providerResources(iuts, logAreas, executionSpaces int) []client.Object {
	labels := map[string]string{"etos.eiffel-community.github.io/environment-request-id": testRequestID}
	meta := func(kind string, i int) metav1.ObjectMeta {
		return metav1.ObjectMeta{Name: fmt.Sprintf("%s-%d", kind, i), Namespace: testNamespace, Labels: labels}
	}
	objects := make([]client.Object, 0, iuts+logAreas+executionSpaces)
	for i := range iuts {
		objects = append(objects, &v1alpha2.Iut{
			ObjectMeta: meta("iut", i),
			Spec:       v1alpha2.IutSpec{ID: fmt.Sprintf("iut-%d", i), Deadline: 100},
		})
	}
	for i := range logAreas {
		objects = append(objects, &v1alpha2.LogArea{
			ObjectMeta: meta("logarea", i),
			Spec:       v1alpha2.LogAreaSpec{ID: fmt.Sprintf("logarea-%d", i), Deadline: 100},
		})
	}
	for i := range executionSpaces {
		objects = append(objects, &v1alpha2.ExecutionSpace{
			ObjectMeta: meta("executionspace", i),
			Spec:       v1alpha2.ExecutionSpaceSpec{ID: fmt.Sprintf("executionspace-%d", i), Deadline: 100},
		})
	}
	return objects
}

// testEnvironmentRequest returns an EnvironmentRequest with the given schema version and amounts.
func testEnvironmentRequest(schemaVersion string, minimum, maximum int) *v1alpha1.EnvironmentRequest {
	return &v1alpha1.EnvironmentRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "rt6t-secret-request-name",
			Namespace:         testNamespace,
			CreationTimestamp: metav1.NewTime(requestTime),
		},
		Spec: v1alpha1.EnvironmentRequestSpec{
			ID:            testRequestID,
			Name:          "rt6t-suite",
			SchemaVersion: schemaVersion,
			MinimumAmount: minimum,
			MaximumAmount: maximum,
			Splitter: v1alpha1.Splitter{Tests: []v1alpha1.Test{
				{ID: "test-1"}, {ID: "test-2"},
			}},
		},
	}
}

// provision runs Provision against a fake Kubernetes client and returns the collected metrics.
func provision(
	t *testing.T,
	ctx context.Context,
	environmentRequest *v1alpha1.EnvironmentRequest,
	funcs interceptor.Funcs,
	objects ...client.Object,
) (metricdata.ResourceMetrics, error) {
	t.Helper()
	provider.SetKubernetesClient(fake.NewClientBuilder().
		WithScheme(provider.Scheme).
		WithObjects(objects...).
		WithInterceptorFuncs(funcs).
		Build())
	t.Cleanup(func() { provider.SetKubernetesClient(nil) })

	reader := sdkmetric.NewManualReader()
	meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
	p := &environmentProvider{meter: meterProvider.Meter(meterName), now: testClock()}
	err := p.Provision(ctx, provider.ProvisionConfig{
		Namespace:          testNamespace,
		EnvironmentRequest: environmentRequest,
		Tracer:             tracenoop.NewTracerProvider().Tracer("test"),
	})
	var rm metricdata.ResourceMetrics
	if collectErr := reader.Collect(context.Background(), &rm); collectErr != nil {
		t.Fatalf("failed to collect metrics: %v", collectErr)
	}
	return rm, err
}

// findMetric returns the metric with the given name, failing the test if it is missing.
func findMetric(t *testing.T, rm metricdata.ResourceMetrics, name string) metricdata.Metrics {
	t.Helper()
	for _, scope := range rm.ScopeMetrics {
		if scope.Scope.Name != meterName {
			continue
		}
		for _, m := range scope.Metrics {
			if m.Name == name {
				return m
			}
		}
	}
	t.Fatalf("metric %q not found", name)
	return metricdata.Metrics{}
}

// hasMetric reports whether the metric with the given name was exported.
func hasMetric(rm metricdata.ResourceMetrics, name string) bool {
	for _, scope := range rm.ScopeMetrics {
		for _, m := range scope.Metrics {
			if m.Name == name {
				return true
			}
		}
	}
	return false
}

// sums returns the monotonic sum data points of a counter keyed by their encoded attributes.
func sums(t *testing.T, m metricdata.Metrics) map[string]int64 {
	t.Helper()
	data, ok := m.Data.(metricdata.Sum[int64])
	if !ok || !data.IsMonotonic {
		t.Fatalf("metric %q is not a monotonic int64 sum: %T", m.Name, m.Data)
	}
	points := map[string]int64{}
	for _, dp := range data.DataPoints {
		points[dp.Attributes.Encoded(attribute.DefaultEncoder())] = dp.Value
	}
	return points
}

// histogram returns the histogram data points of a metric keyed by their encoded attributes.
func histogram(t *testing.T, m metricdata.Metrics) map[string]metricdata.HistogramDataPoint[float64] {
	t.Helper()
	data, ok := m.Data.(metricdata.Histogram[float64])
	if !ok {
		t.Fatalf("metric %q is not a float64 histogram: %T", m.Name, m.Data)
	}
	points := map[string]metricdata.HistogramDataPoint[float64]{}
	for _, dp := range data.DataPoints {
		points[dp.Attributes.Encoded(attribute.DefaultEncoder())] = dp
	}
	return points
}

// attrs encodes attributes the same way as the data point maps.
func attrs(kv ...attribute.KeyValue) string {
	set := attribute.NewSet(kv...)
	return set.Encoded(attribute.DefaultEncoder())
}

// assertSafeAttributes verifies that only bounded attributes are used by every data point.
func assertSafeAttributes(t *testing.T, rm metricdata.ResourceMetrics) {
	t.Helper()
	allowed := map[attribute.Key][]string{
		apiVersionKey: {"v1alpha1", "v1beta1", unknownAttributeValue},
		outcomeKey:    {outcomeSuccess, outcomeFailure, outcomeTimeout},
		stageKey:      {stageResourceLookup, stageCapacityCheck, stageEnvironmentCreate},
	}
	check := func(name string, set attribute.Set) {
		for _, kv := range set.ToSlice() {
			values, ok := allowed[kv.Key]
			if !ok {
				t.Errorf("metric %q has unexpected attribute %q", name, kv.Key)
				continue
			}
			found := false
			for _, v := range values {
				found = found || v == kv.Value.AsString()
			}
			if !found {
				t.Errorf("metric %q has unbounded %q value %q", name, kv.Key, kv.Value.AsString())
			}
		}
	}
	for _, scope := range rm.ScopeMetrics {
		for _, m := range scope.Metrics {
			switch data := m.Data.(type) {
			case metricdata.Sum[int64]:
				for _, dp := range data.DataPoints {
					check(m.Name, dp.Attributes)
				}
			case metricdata.Histogram[float64]:
				for _, dp := range data.DataPoints {
					check(m.Name, dp.Attributes)
				}
			default:
				t.Errorf("metric %q has unexpected data type %T", m.Name, m.Data)
			}
		}
	}
}

// TestProvisionMetricsSuccess verifies names, units, counts and durations of a successful provisioning.
func TestProvisionMetricsSuccess(t *testing.T) {
	rm, err := provision(t, context.Background(), testEnvironmentRequest("v1beta1", 1, 2), interceptor.Funcs{},
		providerResources(2, 2, 2)...)
	if err != nil {
		t.Fatalf("Provision() error = %v", err)
	}
	assertSafeAttributes(t, rm)
	beta := apiVersionKey.String("v1beta1")

	units := map[string]string{
		provisioningDurationName: "s",
		readinessDurationName:    "s",
		creationsName:            "{environment}",
		requestedName:            "{environment}",
		readyName:                "{environment}",
	}
	for name, unit := range units {
		if got := findMetric(t, rm, name).Unit; got != unit {
			t.Errorf("metric %q unit = %q, want %q", name, got, unit)
		}
	}

	creations := sums(t, findMetric(t, rm, creationsName))
	wantCreations := map[string]int64{
		attrs(beta, outcomeKey.String(outcomeSuccess)): 2,
		attrs(beta, outcomeKey.String(outcomeFailure)): 0,
		attrs(beta, outcomeKey.String(outcomeTimeout)): 0,
	}
	if fmt.Sprint(creations) != fmt.Sprint(wantCreations) {
		t.Errorf("creations = %v, want %v", creations, wantCreations)
	}
	if got := sums(t, findMetric(t, rm, requestedName)); got[attrs(beta)] != 1 || len(got) != 1 {
		t.Errorf("requested = %v, want 1", got)
	}
	if got := sums(t, findMetric(t, rm, readyName)); got[attrs(beta)] != 2 || len(got) != 1 {
		t.Errorf("ready = %v, want 2", got)
	}

	provisioning := histogram(t, findMetric(t, rm, provisioningDurationName))
	dp, ok := provisioning[attrs(beta, outcomeKey.String(outcomeSuccess))]
	if !ok || len(provisioning) != 1 || dp.Count != 1 || dp.Sum != 5 {
		t.Errorf("provisioning duration = %+v, want one 5s success observation", provisioning)
	}
	readiness := histogram(t, findMetric(t, rm, readinessDurationName))
	dp, ok = readiness[attrs(beta)]
	if !ok || len(readiness) != 1 || dp.Count != 1 || dp.Sum != 35 {
		t.Errorf("readiness duration = %+v, want one 35s observation", readiness)
	}
}

// TestProvisionMetricsFailures verifies outcomes, failed stages and explicit zeros for failed provisioning.
func TestProvisionMetricsFailures(t *testing.T) {
	createCalls := 0
	failSecondCreate := interceptor.Funcs{
		Create: func(ctx context.Context, c client.WithWatch, obj client.Object, opts ...client.CreateOption) error {
			createCalls++
			if createCalls == 2 {
				return errors.New("admission webhook denied rt6t-secret-request-name")
			}
			return c.Create(ctx, obj, opts...)
		},
	}
	failList := func(err error) interceptor.Funcs {
		return interceptor.Funcs{
			List: func(context.Context, client.WithWatch, client.ObjectList, ...client.ListOption) error {
				return err
			},
		}
	}

	tests := []struct {
		name          string
		schemaVersion string
		funcs         interceptor.Funcs
		resources     []client.Object
		wantVersion   string
		wantOutcome   string
		wantStage     string
		wantCreations map[string]int64
	}{
		{
			name:          "resource lookup failure",
			schemaVersion: "v1beta1",
			funcs:         failList(errors.New("connection refused")),
			wantVersion:   "v1beta1",
			wantOutcome:   outcomeFailure,
			wantStage:     stageResourceLookup,
			wantCreations: map[string]int64{outcomeSuccess: 0, outcomeFailure: 0, outcomeTimeout: 0},
		},
		{
			name:          "resource lookup timeout",
			schemaVersion: "v1alpha1",
			funcs:         failList(fmt.Errorf("list: %w", context.DeadlineExceeded)),
			wantVersion:   "v1alpha1",
			wantOutcome:   outcomeTimeout,
			wantStage:     stageResourceLookup,
			wantCreations: map[string]int64{outcomeSuccess: 0, outcomeFailure: 0, outcomeTimeout: 0},
		},
		{
			name:          "insufficient resources",
			schemaVersion: "",
			resources:     providerResources(1, 2, 2),
			wantVersion:   unknownAttributeValue,
			wantOutcome:   outcomeFailure,
			wantStage:     stageCapacityCheck,
			wantCreations: map[string]int64{outcomeSuccess: 0, outcomeFailure: 0, outcomeTimeout: 0},
		},
		{
			name:          "environment creation failure",
			schemaVersion: "v1beta1",
			funcs:         failSecondCreate,
			resources:     providerResources(2, 2, 2),
			wantVersion:   "v1beta1",
			wantOutcome:   outcomeFailure,
			wantStage:     stageEnvironmentCreate,
			wantCreations: map[string]int64{outcomeSuccess: 1, outcomeFailure: 1, outcomeTimeout: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createCalls = 0
			rm, err := provision(t, context.Background(), testEnvironmentRequest(tt.schemaVersion, 2, 2), tt.funcs,
				tt.resources...)
			if err == nil {
				t.Fatal("Provision() succeeded, want error")
			}
			assertSafeAttributes(t, rm)
			version := apiVersionKey.String(tt.wantVersion)

			creations := sums(t, findMetric(t, rm, creationsName))
			for outcome, want := range tt.wantCreations {
				if got, ok := creations[attrs(version, outcomeKey.String(outcome))]; !ok || got != want {
					t.Errorf("creations{%s} = %d (present %t), want %d", outcome, got, ok, want)
				}
			}
			if got := sums(t, findMetric(t, rm, requestedName)); got[attrs(version)] != 2 {
				t.Errorf("requested = %v, want 2", got)
			}
			ready, ok := sums(t, findMetric(t, rm, readyName))[attrs(version)]
			if !ok || ready != 0 {
				t.Errorf("ready = %d (present %t), want explicit 0", ready, ok)
			}

			provisioning := histogram(t, findMetric(t, rm, provisioningDurationName))
			dp, ok := provisioning[attrs(version, outcomeKey.String(tt.wantOutcome), stageKey.String(tt.wantStage))]
			if !ok || len(provisioning) != 1 || dp.Count != 1 || dp.Sum != 5 {
				t.Errorf("provisioning duration = %+v, want one 5s %s/%s observation",
					provisioning, tt.wantOutcome, tt.wantStage)
			}
			if hasMetric(rm, readinessDurationName) {
				t.Errorf("readiness duration recorded on failure")
			}
		})
	}
}

// TestProvisionMetricsDeadlineContext verifies that an expired provisioning deadline is a timeout.
func TestProvisionMetricsDeadlineContext(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
	defer cancel()
	failCreate := interceptor.Funcs{
		Create: func(context.Context, client.WithWatch, client.Object, ...client.CreateOption) error {
			return errors.New("client rate limiter Wait returned an error")
		},
	}
	rm, err := provision(t, ctx, testEnvironmentRequest("v1beta1", 1, 1), failCreate, providerResources(1, 1, 1)...)
	if err == nil {
		t.Fatal("Provision() succeeded, want error")
	}
	beta := apiVersionKey.String("v1beta1")
	if got := sums(t, findMetric(t, rm, creationsName))[attrs(beta, outcomeKey.String(outcomeTimeout))]; got != 1 {
		t.Errorf("creations{timeout} = %d, want 1", got)
	}
	provisioning := histogram(t, findMetric(t, rm, provisioningDurationName))
	key := attrs(beta, outcomeKey.String(outcomeTimeout), stageKey.String(stageEnvironmentCreate))
	if dp, ok := provisioning[key]; !ok || dp.Count != 1 {
		t.Errorf("provisioning duration = %+v, want one timeout observation", provisioning)
	}
}

// TestProvisioningRecorderSingleCount verifies that a provisioning attempt is recorded only once.
func TestProvisioningRecorderSingleCount(t *testing.T) {
	ctx := context.Background()
	reader := sdkmetric.NewManualReader()
	m, err := newProvisioningMetrics(sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader)).Meter(meterName))
	if err != nil {
		t.Fatalf("newProvisioningMetrics() error = %v", err)
	}
	recorder := m.start(ctx, testEnvironmentRequest("v1beta1", 1, 1), 1, testClock())
	recorder.succeeded(ctx, 1)
	recorder.failed(ctx, stageEnvironmentCreate, errors.New("late failure"))
	recorder.succeeded(ctx, 1)

	var rm metricdata.ResourceMetrics
	if err := reader.Collect(ctx, &rm); err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	beta := apiVersionKey.String("v1beta1")
	provisioning := histogram(t, findMetric(t, rm, provisioningDurationName))
	if dp := provisioning[attrs(beta, outcomeKey.String(outcomeSuccess))]; len(provisioning) != 1 || dp.Count != 1 {
		t.Errorf("provisioning duration = %+v, want exactly one success observation", provisioning)
	}
	if got := sums(t, findMetric(t, rm, readyName))[attrs(beta)]; got != 1 {
		t.Errorf("ready = %d, want 1", got)
	}
	if got := histogram(t, findMetric(t, rm, readinessDurationName))[attrs(beta)]; got.Count != 1 {
		t.Errorf("readiness count = %d, want 1", got.Count)
	}
}

// TestProvisionMetricsGlobalMeter verifies that, without an override, metrics are recorded through
// the global meter provider that the provider framework configures and flushes.
func TestProvisionMetricsGlobalMeter(t *testing.T) {
	previous := otel.GetMeterProvider()
	t.Cleanup(func() { otel.SetMeterProvider(previous) })
	reader := sdkmetric.NewManualReader()
	otel.SetMeterProvider(sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader)))

	provider.SetKubernetesClient(fake.NewClientBuilder().
		WithScheme(provider.Scheme).
		WithObjects(providerResources(1, 1, 1)...).
		Build())
	t.Cleanup(func() { provider.SetKubernetesClient(nil) })

	p := &environmentProvider{}
	if err := p.Provision(context.Background(), provider.ProvisionConfig{
		Namespace:          testNamespace,
		EnvironmentRequest: testEnvironmentRequest("v1beta1", 1, 1),
		Tracer:             tracenoop.NewTracerProvider().Tracer("test"),
	}); err != nil {
		t.Fatalf("Provision() error = %v", err)
	}
	var rm metricdata.ResourceMetrics
	if err := reader.Collect(context.Background(), &rm); err != nil {
		t.Fatalf("failed to collect metrics: %v", err)
	}
	provisioning := histogram(t, findMetric(t, rm, provisioningDurationName))
	key := attrs(apiVersionKey.String("v1beta1"), outcomeKey.String(outcomeSuccess))
	if dp, ok := provisioning[key]; !ok || dp.Count != 1 || dp.Sum < 0 {
		t.Errorf("provisioning duration = %+v, want one success observation", provisioning)
	}
}

// TestProvisionMetricsPanic verifies that a panic during provisioning records exactly one failure
// in the stage that panicked and that the panic is propagated.
func TestProvisionMetricsPanic(t *testing.T) {
	tests := []struct {
		name      string
		funcs     interceptor.Funcs
		wantStage string
	}{
		{
			name: "panic during resource lookup",
			funcs: interceptor.Funcs{
				List: func(context.Context, client.WithWatch, client.ObjectList, ...client.ListOption) error {
					panic("list exploded for rt6t-secret-request-name")
				},
			},
			wantStage: stageResourceLookup,
		},
		{
			name: "panic during environment creation",
			funcs: interceptor.Funcs{
				Create: func(context.Context, client.WithWatch, client.Object, ...client.CreateOption) error {
					panic("create exploded for rt6t-secret-request-name")
				},
			},
			wantStage: stageEnvironmentCreate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			provider.SetKubernetesClient(fake.NewClientBuilder().
				WithScheme(provider.Scheme).
				WithObjects(providerResources(1, 1, 1)...).
				WithInterceptorFuncs(tt.funcs).
				Build())
			t.Cleanup(func() { provider.SetKubernetesClient(nil) })

			reader := sdkmetric.NewManualReader()
			meterProvider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
			p := &environmentProvider{meter: meterProvider.Meter(meterName), now: testClock()}

			recovered := func() (v any) {
				defer func() { v = recover() }()
				_ = p.Provision(context.Background(), provider.ProvisionConfig{
					Namespace:          testNamespace,
					EnvironmentRequest: testEnvironmentRequest("v1beta1", 1, 1),
					Tracer:             tracenoop.NewTracerProvider().Tracer("test"),
				})
				return nil
			}()
			if recovered == nil {
				t.Fatal("Provision() did not re-panic")
			}

			var rm metricdata.ResourceMetrics
			if err := reader.Collect(context.Background(), &rm); err != nil {
				t.Fatalf("failed to collect metrics: %v", err)
			}
			assertSafeAttributes(t, rm)
			beta := apiVersionKey.String("v1beta1")

			provisioning := histogram(t, findMetric(t, rm, provisioningDurationName))
			key := attrs(beta, outcomeKey.String(outcomeFailure), stageKey.String(tt.wantStage))
			if dp, ok := provisioning[key]; !ok || len(provisioning) != 1 || dp.Count != 1 || dp.Sum != 5 {
				t.Errorf("provisioning duration = %+v, want one 5s failure/%s observation", provisioning, tt.wantStage)
			}
			ready, ok := sums(t, findMetric(t, rm, readyName))[attrs(beta)]
			if !ok || ready != 0 {
				t.Errorf("ready = %d (present %t), want explicit 0", ready, ok)
			}
			if hasMetric(rm, readinessDurationName) {
				t.Error("readiness duration recorded on panic")
			}
		})
	}
}
