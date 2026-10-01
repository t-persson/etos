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
	"slices"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// Metric conventions:
//   - Metrics are exported over OTLP through the global meter provider, which run flushes on
//     shutdown on every exit path. The ETOS release is the etos.version resource attribute,
//     never a metric attribute.
//   - Each provider process makes exactly one provisioning or release attempt, recorded once.
//     Providers do not retry and their Jobs do not restart, so there is no retry metric.
//   - Attributes are bounded: provider type, API version, normalized outcome and, on the
//     duration histograms of unsuccessful attempts, the failed stage or a normalized error type.
//     Resource names, provider values and error messages are never used.
//   - Attempt counters add zero for the outcomes that did not occur, so dashboards see zero
//     rather than missing data while telemetry is arriving.
//   - The Environment Provider is not instrumented here; it records its own metrics.
const (
	meterName = "github.com/eiffel-community/etos/pkg/provider"

	providerTypeKey = attribute.Key("etos.provider.type")
	apiVersionKey   = attribute.Key("etos.api_version")
	outcomeKey      = attribute.Key("etos.outcome")
	stageKey        = attribute.Key("etos.stage")
	errorTypeKey    = attribute.Key("error.type")

	outcomeSuccess = "success"
	outcomeFailure = "failure"
	outcomeTimeout = "timeout"

	// stageStartup is a failure before the provider operation started, such as fetching the
	// EnvironmentRequest or connecting to the message bus.
	stageStartup = "startup"
	// stageOperation is a failure of the provisioning or release operation itself.
	stageOperation = "operation"
	// stageResult is a failure to write the result to the termination-log.
	stageResult = "result"

	errorTypeTestRunner   = "test_runner_error"
	errorTypeShutdown     = "shutdown"
	errorTypeStreamClosed = "stream_closed"
	errorTypeStreamError  = "stream_error"
	errorTypeCanceled     = "canceled"
	errorTypeTimeout      = "timeout"

	unknownAttributeValue = "unknown"
	attemptUnit           = "{attempt}"
	secondsUnit           = "s"

	provisioningAttemptsName   = "etos.provider.provisioning.attempts"
	provisioningDurationName   = "etos.provider.provisioning.duration"
	releaseAttemptsName        = "etos.provider.release.attempts"
	releaseDurationName        = "etos.provider.release.duration"
	testRunnerReadinessAttName = "etos.provider.test_runner.readiness.attempts"
	testRunnerReadinessDurName = "etos.provider.test_runner.readiness.duration"
)

var (
	supportedAPIVersions = []string{"v1alpha1", "v1beta1"}
	outcomes             = []string{outcomeSuccess, outcomeFailure, outcomeTimeout}

	// instrumentedProviderTypes maps the provider types run by this package to bounded
	// etos.provider.type values.
	instrumentedProviderTypes = map[string]string{
		"Iut":            "iut",
		"LogArea":        "log_area",
		"ExecutionSpace": "execution_space",
	}

	// durationBuckets cover provider work from sub-second up to the longest supported timeouts.
	durationBuckets = []float64{0.5, 1, 2.5, 5, 10, 30, 60, 120, 300, 600, 900, 1800, 3600}
)

// operationMetrics holds the instruments for one provider operation (provisioning or release).
type operationMetrics struct {
	attempts metric.Int64Counter
	duration metric.Float64Histogram
}

// newOperationMetrics creates the attempt counter and duration histogram of an operation.
func newOperationMetrics(
	meter metric.Meter, attemptsName, durationName, description string,
) (operationMetrics, error) {
	attempts, attemptsErr := meter.Int64Counter(
		attemptsName,
		metric.WithUnit(attemptUnit),
		metric.WithDescription(description+" attempts by outcome."),
	)
	duration, durationErr := meter.Float64Histogram(
		durationName,
		metric.WithUnit(secondsUnit),
		metric.WithDescription("Duration of "+description+" attempts, from provider start until the result "+
			"is known."),
		metric.WithExplicitBucketBoundaries(durationBuckets...),
	)
	return operationMetrics{attempts: attempts, duration: duration}, errors.Join(attemptsErr, durationErr)
}

// record records exactly one attempt. Outcomes that did not occur are added as zero.
// The extra attributes are added to the duration histogram only.
func (m operationMetrics) record(
	ctx context.Context,
	attrs []attribute.KeyValue,
	outcome string,
	seconds float64,
	extra ...attribute.KeyValue,
) {
	for _, candidate := range outcomes {
		var value int64
		if candidate == outcome {
			value = 1
		}
		m.attempts.Add(ctx, value, metric.WithAttributes(append(slices.Clone(attrs), outcomeKey.String(candidate))...))
	}
	durationAttrs := append(slices.Clone(attrs), outcomeKey.String(outcome))
	m.duration.Record(ctx, seconds, metric.WithAttributes(append(durationAttrs, extra...)...))
}

// providerAttempt records the single provisioning or release attempt of a provider process.
// A nil providerAttempt records nothing.
type providerAttempt struct {
	metrics      operationMetrics
	providerType string
	apiVersion   string
	stage        string
	start        time.Time
	now          func() time.Time
}

// newProviderAttempt starts recording the attempt of a provider process using the given meter
// provider. It returns nil for provider types that are not instrumented by this package.
func newProviderAttempt(
	meterProvider metric.MeterProvider, providerType string, release bool, start time.Time,
) *providerAttempt {
	metricType, ok := instrumentedProviderTypes[providerType]
	if !ok {
		return nil
	}
	meter := meterProvider.Meter(meterName)
	var metrics operationMetrics
	var err error
	if release {
		metrics, err = newOperationMetrics(meter, releaseAttemptsName, releaseDurationName, "provider release")
	} else {
		metrics, err = newOperationMetrics(
			meter, provisioningAttemptsName, provisioningDurationName, "provider provisioning",
		)
	}
	if err != nil {
		otel.Handle(err)
	}
	return &providerAttempt{
		metrics:      metrics,
		providerType: metricType,
		apiVersion:   unknownAttributeValue,
		stage:        stageStartup,
		start:        start,
		now:          time.Now,
	}
}

// setAPIVersion sets the API version attribute from the EnvironmentRequest schema version.
func (a *providerAttempt) setAPIVersion(schemaVersion string) {
	if a != nil {
		a.apiVersion = apiVersion(schemaVersion)
	}
}

// setStage sets the stage that a failure from now on is attributed to.
func (a *providerAttempt) setStage(stage string) {
	if a != nil {
		a.stage = stage
	}
}

// finish records the attempt with the result stored in err. It must be deferred directly so that
// a panic is recorded as a failure before it continues to propagate.
func (a *providerAttempt) finish(err *error) {
	if a == nil {
		return
	}
	if recovered := recover(); recovered != nil {
		a.record(errors.New("provider panicked"))
		panic(recovered)
	}
	a.record(*err)
}

// record records the attempt outcome and duration for the given result.
func (a *providerAttempt) record(err error) {
	ctx := context.Background()
	attrs := []attribute.KeyValue{
		providerTypeKey.String(a.providerType),
		apiVersionKey.String(a.apiVersion),
	}
	seconds := secondsBetween(a.start, a.now())
	if err == nil {
		a.metrics.record(ctx, attrs, outcomeSuccess, seconds)
		return
	}
	stage := a.stage
	if errors.As(err, new(*terminationLogError)) {
		stage = stageResult
	}
	a.metrics.record(ctx, attrs, errorOutcome(err), seconds, stageKey.String(stage))
}

// terminationLogError marks a failure to write a successful result to the termination-log.
// Its message is the message of the wrapped error.
type terminationLogError struct {
	err error
}

// Error returns the message of the wrapped error.
func (e *terminationLogError) Error() string {
	return e.err.Error()
}

// Unwrap returns the wrapped error.
func (e *terminationLogError) Unwrap() error {
	return e.err
}

// recordTestRunnerReadiness records one Test Runner readiness wait of the Execution Space
// Provider. errorType is a bounded failure classification and is ignored on success.
func recordTestRunnerReadiness(
	ctx context.Context, schemaVersion string, start time.Time, err error, errorType string,
) {
	meter := otel.GetMeterProvider().Meter(meterName)
	attempts, attemptsErr := meter.Int64Counter(
		testRunnerReadinessAttName,
		metric.WithUnit(attemptUnit),
		metric.WithDescription("Test Runner readiness waits by the Execution Space Provider by outcome."),
	)
	duration, durationErr := meter.Float64Histogram(
		testRunnerReadinessDurName,
		metric.WithUnit(secondsUnit),
		metric.WithDescription("Duration of Test Runner readiness waits by the Execution Space Provider, "+
			"until the Test Runner reports its status or the wait fails."),
		metric.WithExplicitBucketBoundaries(durationBuckets...),
	)
	if joined := errors.Join(attemptsErr, durationErr); joined != nil {
		otel.Handle(joined)
	}
	metrics := operationMetrics{attempts: attempts, duration: duration}
	attrs := []attribute.KeyValue{apiVersionKey.String(apiVersion(schemaVersion))}
	seconds := secondsBetween(start, time.Now())
	// Record with an independent context: ctx may already be canceled when the wait fails.
	recordCtx := context.WithoutCancel(ctx)
	if err == nil {
		metrics.record(recordCtx, attrs, outcomeSuccess, seconds)
		return
	}
	metrics.record(recordCtx, attrs, errorOutcome(err), seconds, errorTypeKey.String(errorType))
}

// errorOutcome normalizes a failure into the timeout or failure outcome.
func errorOutcome(err error) string {
	if errors.Is(err, context.DeadlineExceeded) {
		return outcomeTimeout
	}
	return outcomeFailure
}

// apiVersion returns the bounded API version attribute value.
func apiVersion(schemaVersion string) string {
	if slices.Contains(supportedAPIVersions, schemaVersion) {
		return schemaVersion
	}
	return unknownAttributeValue
}

// secondsBetween returns the non-negative number of seconds between start and end.
func secondsBetween(start, end time.Time) float64 {
	return max(end.Sub(start).Seconds(), 0)
}
