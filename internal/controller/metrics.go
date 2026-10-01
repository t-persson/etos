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

package controller

import (
	"slices"
	"time"

	etosv1alpha1 "github.com/eiffel-community/etos/api/v1alpha1"
	"github.com/eiffel-community/etos/internal/controller/status"
	"github.com/eiffel-community/etos/pkg/version"
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

// Metric conventions:
//   - Counters and histograms record transitions only after the transition has been persisted
//     to the Kubernetes API, so a status update conflict or retry never counts an event twice.
//     They are process-local and reset on controller restart; use rate()/increase() to query them.
//   - Current-state gauges are not maintained in memory. They are computed at scrape time from
//     the informer cache by stateCollector, so they are correct after a controller restart.
//   - Labels are bounded: API version, cluster name, stage, and normalized outcome only.
//   - The ETOS release version is exposed once through etos_controller_build_info.

const (
	outcomeSuccess = "success"
	outcomeFailure = "failure"
	outcomeTimeout = "timeout"

	stageProvisioning = "provisioning"
	stageReadiness    = "readiness"

	unknownLabel = "unknown"
	clusterLabel = "etos.eiffel-community.github.io/cluster"
)

var (
	supportedAPIVersions = []string{"v1alpha1", "v1beta1"}
	testRunOutcomes      = []string{outcomeSuccess, outcomeFailure, outcomeTimeout}
	provisioningOutcomes = []string{outcomeSuccess, outcomeFailure}
	releaseOutcomes      = []string{outcomeSuccess, outcomeFailure}

	// lifecycleBuckets cover controller lifecycles from seconds up to eight hours.
	lifecycleBuckets = []float64{1, 5, 10, 30, 60, 120, 300, 600, 900, 1800, 3600, 7200, 14400, 28800}
)

var (
	buildInfo = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:        "etos_controller_build_info",
			Help:        "Build information about the running ETOS controller. Always 1.",
			ConstLabels: prometheus.Labels{"etos_version": version.Version},
		},
	)
	testRunsStarted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "etos_controller_testruns_started_total",
			Help: "Number of TestRuns whose controller lifecycle started.",
		},
		[]string{"api_version", "cluster"},
	)
	testRunsTerminal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "etos_controller_testruns_terminal_total",
			Help: "Number of TestRuns reaching a terminal controller outcome.",
		},
		[]string{"api_version", "cluster", "outcome"},
	)
	testRunsDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "etos_controller_testrun_duration_seconds",
			Help:    "Duration of TestRuns from status start to terminal transition.",
			Buckets: lifecycleBuckets,
		},
		[]string{"api_version", "cluster", "outcome"},
	)
	environmentRequestsStarted = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "etos_controller_environment_requests_started_total",
			Help: "Number of EnvironmentRequests whose controller lifecycle started.",
		},
		[]string{"api_version", "cluster"},
	)
	environmentRequestJobOutcomes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "etos_controller_environment_request_jobs_total",
			Help: "Number of EnvironmentRequest provisioning outcomes persisted by the controller.",
		},
		[]string{"api_version", "cluster", "stage", "outcome"},
	)
	environmentRequestStageDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "etos_controller_environment_request_stage_duration_seconds",
			Help:    "Duration of EnvironmentRequest provisioning and TestRun environment readiness stages.",
			Buckets: lifecycleBuckets,
		},
		[]string{"api_version", "cluster", "stage", "outcome"},
	)
	environmentReleaseAttempts = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "etos_controller_environment_release_attempts_total",
			Help: "Number of environment release cleanup attempts by outcome.",
		},
		[]string{"api_version", "cluster", "outcome"},
	)
	environmentReleaseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "etos_controller_environment_release_duration_seconds",
			Help:    "Duration of completed environment releases from deletion request to finalizer removal.",
			Buckets: lifecycleBuckets,
		},
		[]string{"api_version", "cluster"},
	)
)

func init() {
	buildInfo.Set(1)
	metrics.Registry.MustRegister(
		buildInfo,
		testRunsStarted,
		testRunsTerminal,
		testRunsDuration,
		environmentRequestsStarted,
		environmentRequestJobOutcomes,
		environmentRequestStageDuration,
		environmentReleaseAttempts,
		environmentReleaseDuration,
	)
}

// apiVersion returns the bounded API version label used by controller metrics.
func apiVersion(schemaVersion string) string {
	if slices.Contains(supportedAPIVersions, schemaVersion) {
		return schemaVersion
	}
	return unknownLabel
}

// clusterName returns a cluster label value with an explicit fallback for missing values.
func clusterName(cluster string) string {
	if cluster == "" {
		return unknownLabel
	}
	return cluster
}

// terminalOutcome maps a controller condition reason to a bounded metric outcome.
func terminalOutcome(reason string) string {
	switch reason {
	case status.ReasonTimedOut:
		return outcomeTimeout
	case status.ReasonCompleted:
		return outcomeSuccess
	default:
		return outcomeFailure
	}
}

// isTerminalReason reports whether a status reason represents a terminal lifecycle transition.
func isTerminalReason(reason string) bool {
	return reason == status.ReasonCompleted || reason == status.ReasonFailed || reason == status.ReasonTimedOut
}

// secondsSince returns the non-negative number of seconds between start and now.
func secondsSince(start, now time.Time) float64 {
	return max(now.Sub(start).Seconds(), 0)
}

// testRunLabels returns the bounded API version and cluster labels of a TestRun.
func testRunLabels(testrun *etosv1alpha1.TestRun) (string, string) {
	return apiVersion(testrun.Spec.SchemaVersion), clusterName(testrun.Spec.Cluster)
}

// environmentRequestLabels returns the bounded API version and cluster labels of an EnvironmentRequest.
func environmentRequestLabels(environmentrequest *etosv1alpha1.EnvironmentRequest) (string, string) {
	return apiVersion(environmentrequest.Spec.SchemaVersion), clusterName(environmentrequest.Labels[clusterLabel])
}

// recordTestRunStarted records one persisted TestRun start transition.
// Outcome series are initialized so that dashboards see zero rather than missing data.
func recordTestRunStarted(testrun *etosv1alpha1.TestRun) {
	apiVersionLabel, cluster := testRunLabels(testrun)
	testRunsStarted.WithLabelValues(apiVersionLabel, cluster).Inc()
	for _, outcome := range testRunOutcomes {
		testRunsTerminal.WithLabelValues(apiVersionLabel, cluster, outcome)
	}
}

// recordTestRunTerminal records one persisted terminal TestRun transition.
func recordTestRunTerminal(testrun *etosv1alpha1.TestRun, reason string, now time.Time) {
	if !isTerminalReason(reason) {
		return
	}
	outcome := terminalOutcome(reason)
	apiVersionLabel, cluster := testRunLabels(testrun)
	testRunsTerminal.WithLabelValues(apiVersionLabel, cluster, outcome).Inc()
	if testrun.Status.StartTime != nil {
		testRunsDuration.WithLabelValues(apiVersionLabel, cluster, outcome).Observe(
			secondsSince(testrun.Status.StartTime.Time, now),
		)
	}
}

// recordTestRunEnvironmentReady records one persisted TestRun environment readiness transition.
func recordTestRunEnvironmentReady(testrun *etosv1alpha1.TestRun, now time.Time) {
	if testrun.Status.StartTime == nil {
		return
	}
	apiVersionLabel, cluster := testRunLabels(testrun)
	environmentRequestStageDuration.WithLabelValues(apiVersionLabel, cluster, stageReadiness, outcomeSuccess).Observe(
		secondsSince(testrun.Status.StartTime.Time, now),
	)
}

// recordEnvironmentRequestStarted records one persisted EnvironmentRequest start transition.
// Outcome series are initialized so that dashboards see zero rather than missing data.
func recordEnvironmentRequestStarted(environmentrequest *etosv1alpha1.EnvironmentRequest) {
	apiVersionLabel, cluster := environmentRequestLabels(environmentrequest)
	environmentRequestsStarted.WithLabelValues(apiVersionLabel, cluster).Inc()
	for _, outcome := range provisioningOutcomes {
		environmentRequestJobOutcomes.WithLabelValues(apiVersionLabel, cluster, stageProvisioning, outcome)
	}
	for _, outcome := range releaseOutcomes {
		environmentReleaseAttempts.WithLabelValues(apiVersionLabel, cluster, outcome)
	}
}

// recordEnvironmentRequestOutcome records one persisted terminal provisioning transition.
func recordEnvironmentRequestOutcome(environmentrequest *etosv1alpha1.EnvironmentRequest, outcome string, now time.Time) {
	apiVersionLabel, cluster := environmentRequestLabels(environmentrequest)
	environmentRequestJobOutcomes.WithLabelValues(apiVersionLabel, cluster, stageProvisioning, outcome).Inc()
	environmentRequestStageDuration.WithLabelValues(apiVersionLabel, cluster, stageProvisioning, outcome).Observe(
		secondsSince(environmentRequestStart(environmentrequest, now), now),
	)
}

// recordEnvironmentReleaseFailure records one failed environment release cleanup attempt.
func recordEnvironmentReleaseFailure(environmentrequest *etosv1alpha1.EnvironmentRequest) {
	apiVersionLabel, cluster := environmentRequestLabels(environmentrequest)
	environmentReleaseAttempts.WithLabelValues(apiVersionLabel, cluster, outcomeFailure).Inc()
}

// recordEnvironmentReleaseSuccess records one completed environment release after its finalizer was removed.
func recordEnvironmentReleaseSuccess(environmentrequest *etosv1alpha1.EnvironmentRequest, now time.Time) {
	apiVersionLabel, cluster := environmentRequestLabels(environmentrequest)
	environmentReleaseAttempts.WithLabelValues(apiVersionLabel, cluster, outcomeSuccess).Inc()
	start := environmentRequestStart(environmentrequest, now)
	if !environmentrequest.DeletionTimestamp.IsZero() {
		start = environmentrequest.DeletionTimestamp.Time
	}
	environmentReleaseDuration.WithLabelValues(apiVersionLabel, cluster).Observe(secondsSince(start, now))
}

// environmentRequestStart returns the best persisted start timestamp for a request.
func environmentRequestStart(environmentrequest *etosv1alpha1.EnvironmentRequest, now time.Time) time.Time {
	if environmentrequest.Status.StartTime != nil {
		return environmentrequest.Status.StartTime.Time
	}
	if !environmentrequest.CreationTimestamp.IsZero() {
		return environmentrequest.CreationTimestamp.Time
	}
	return now
}

// isTestRunActive reports whether a TestRun has started and not yet reached a terminal state.
func isTestRunActive(testrun *etosv1alpha1.TestRun) bool {
	return testrun.Status.StartTime != nil && testrun.Status.CompletionTime == nil
}
