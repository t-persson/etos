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
	"slices"
	"time"

	"github.com/eiffel-community/etos/api/v1alpha1"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

// Metric conventions:
//   - Metrics are exported over OTLP through the global meter provider, which the provider
//     framework flushes on shutdown on every exit path. The ETOS release is the etos.version
//     resource attribute, never a metric attribute.
//   - Attributes are bounded: API version, normalized outcome and, for non-successful
//     provisioning, the failed stage. Resource names, provider values and errors are never used.
//   - Each provisioning attempt is recorded exactly once, including when provisioning panics.
const (
	meterName = "github.com/eiffel-community/etos/cmd/environmentprovider"

	apiVersionKey = attribute.Key("etos.api_version")
	outcomeKey    = attribute.Key("etos.outcome")
	stageKey      = attribute.Key("etos.stage")

	outcomeSuccess = "success"
	outcomeFailure = "failure"
	outcomeTimeout = "timeout"

	stageResourceLookup    = "resource_lookup"
	stageCapacityCheck     = "capacity_check"
	stageEnvironmentCreate = "environment_creation"

	unknownAttributeValue = "unknown"

	environmentUnit = "{environment}"
	secondsUnit     = "s"

	provisioningDurationName = "etos.environment_provider.provisioning.duration"
	readinessDurationName    = "etos.environment_provider.readiness.duration"
	creationsName            = "etos.environment_provider.environment.creations"
	requestedName            = "etos.environment_provider.environments.requested"
	readyName                = "etos.environment_provider.environments.ready"
)

var (
	supportedAPIVersions = []string{"v1alpha1", "v1beta1"}
	creationOutcomes     = []string{outcomeSuccess, outcomeFailure, outcomeTimeout}

	// durationBuckets cover provisioning from sub-second up to the longest supported timeouts.
	durationBuckets = []float64{0.5, 1, 2.5, 5, 10, 30, 60, 120, 300, 600, 900, 1800, 3600}
)

// provisioningMetrics holds the Environment Provider metric instruments.
type provisioningMetrics struct {
	provisioningDuration metric.Float64Histogram
	readinessDuration    metric.Float64Histogram
	creations            metric.Int64Counter
	requested            metric.Int64Counter
	ready                metric.Int64Counter
}

// newProvisioningMetrics creates the Environment Provider metric instruments from a meter.
func newProvisioningMetrics(meter metric.Meter) (*provisioningMetrics, error) {
	provisioningDuration, provisioningDurationErr := meter.Float64Histogram(
		provisioningDurationName,
		metric.WithUnit(secondsUnit),
		metric.WithDescription("Duration of one Environment Provider provisioning attempt, from collecting "+
			"provider resources until all Environment resources are created or provisioning fails."),
		metric.WithExplicitBucketBoundaries(durationBuckets...),
	)
	readinessDuration, readinessDurationErr := meter.Float64Histogram(
		readinessDurationName,
		metric.WithUnit(secondsUnit),
		metric.WithDescription("Time from EnvironmentRequest creation until all of its Environment resources "+
			"were created by the Environment Provider. Recorded for successful provisioning only."),
		metric.WithExplicitBucketBoundaries(durationBuckets...),
	)
	creations, creationsErr := meter.Int64Counter(
		creationsName,
		metric.WithUnit(environmentUnit),
		metric.WithDescription("Environment resource creation attempts by outcome."),
	)
	requested, requestedErr := meter.Int64Counter(
		requestedName,
		metric.WithUnit(environmentUnit),
		metric.WithDescription("Minimum number of environments required by provisioning attempts."),
	)
	ready, readyErr := meter.Int64Counter(
		readyName,
		metric.WithUnit(environmentUnit),
		metric.WithDescription("Environments handed off by successful provisioning attempts. "+
			"Failed attempts add zero."),
	)
	if err := errors.Join(
		provisioningDurationErr, readinessDurationErr, creationsErr, requestedErr, readyErr,
	); err != nil {
		return nil, err
	}
	return &provisioningMetrics{
		provisioningDuration: provisioningDuration,
		readinessDuration:    readinessDuration,
		creations:            creations,
		requested:            requested,
		ready:                ready,
	}, nil
}

// noopProvisioningMetrics returns metric instruments that record nothing.
func noopProvisioningMetrics() *provisioningMetrics {
	m, _ := newProvisioningMetrics(noop.NewMeterProvider().Meter(meterName))
	return m
}

// provisioningRecorder records the metrics of a single provisioning attempt.
type provisioningRecorder struct {
	metrics    *provisioningMetrics
	apiVersion attribute.KeyValue
	requestAt  time.Time
	start      time.Time
	now        func() time.Time
	stage      string
	finished   bool
}

// start begins recording a provisioning attempt for an EnvironmentRequest.
// Creation outcome series are initialized so that dashboards see zero rather than missing data.
func (m *provisioningMetrics) start(
	ctx context.Context,
	environmentRequest *v1alpha1.EnvironmentRequest,
	requested int,
	now func() time.Time,
) *provisioningRecorder {
	r := &provisioningRecorder{
		metrics:    m,
		apiVersion: apiVersionKey.String(apiVersion(environmentRequest.Spec.SchemaVersion)),
		requestAt:  environmentRequest.CreationTimestamp.Time,
		start:      now(),
		now:        now,
		stage:      stageResourceLookup,
	}
	attrs := metric.WithAttributes(r.apiVersion)
	m.requested.Add(ctx, int64(max(requested, 0)), attrs)
	for _, outcome := range creationOutcomes {
		m.creations.Add(ctx, 0, metric.WithAttributes(r.apiVersion, outcomeKey.String(outcome)))
	}
	return r
}

// enter marks the provisioning stage that is currently executing.
func (r *provisioningRecorder) enter(stage string) {
	r.stage = stage
}

// finishOnPanic must be deferred directly after start. If provisioning panics before the attempt
// was recorded, it records a failure in the current stage and then re-panics, so that the
// framework's deferred shutdown can still flush the outcome.
func (r *provisioningRecorder) finishOnPanic(ctx context.Context) {
	if v := recover(); v != nil {
		r.finish(ctx, r.stage, outcomeFailure)
		panic(v)
	}
}

// environmentCreated records the outcome of one Environment resource creation attempt.
func (r *provisioningRecorder) environmentCreated(ctx context.Context, err error) {
	r.metrics.creations.Add(ctx, 1, metric.WithAttributes(r.apiVersion, outcomeKey.String(outcome(ctx, err))))
}

// succeeded records a successful provisioning attempt that handed off the given number of environments.
func (r *provisioningRecorder) succeeded(ctx context.Context, ready int) {
	if r.finished {
		return
	}
	r.finished = true
	now := r.now()
	r.metrics.ready.Add(ctx, int64(ready), metric.WithAttributes(r.apiVersion))
	r.metrics.provisioningDuration.Record(ctx, secondsBetween(r.start, now),
		metric.WithAttributes(r.apiVersion, outcomeKey.String(outcomeSuccess)))
	if !r.requestAt.IsZero() {
		r.metrics.readinessDuration.Record(ctx, secondsBetween(r.requestAt, now),
			metric.WithAttributes(r.apiVersion))
	}
}

// failed records a provisioning attempt that failed in the given stage.
func (r *provisioningRecorder) failed(ctx context.Context, stage string, err error) {
	r.finish(ctx, stage, outcome(ctx, err))
}

// finish records an unsuccessful provisioning attempt with the given stage and outcome.
func (r *provisioningRecorder) finish(ctx context.Context, stage, outcome string) {
	if r.finished {
		return
	}
	r.finished = true
	r.metrics.ready.Add(ctx, 0, metric.WithAttributes(r.apiVersion))
	r.metrics.provisioningDuration.Record(ctx, secondsBetween(r.start, r.now()),
		metric.WithAttributes(r.apiVersion, outcomeKey.String(outcome), stageKey.String(stage)))
}

// apiVersion returns the bounded API version attribute value.
func apiVersion(schemaVersion string) string {
	if slices.Contains(supportedAPIVersions, schemaVersion) {
		return schemaVersion
	}
	return unknownAttributeValue
}

// outcome normalizes an error into a bounded outcome. The provisioning deadline is a timeout.
func outcome(ctx context.Context, err error) string {
	switch {
	case err == nil:
		return outcomeSuccess
	case errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded):
		return outcomeTimeout
	default:
		return outcomeFailure
	}
}

// secondsBetween returns the non-negative number of seconds between start and end.
func secondsBetween(start, end time.Time) float64 {
	return max(end.Sub(start).Seconds(), 0)
}
