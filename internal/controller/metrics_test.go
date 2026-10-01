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
	"strings"
	"testing"
	"time"

	etosv1alpha1 "github.com/eiffel-community/etos/api/v1alpha1"
	etosv1alpha2 "github.com/eiffel-community/etos/api/v1alpha2"
	"github.com/eiffel-community/etos/internal/controller/status"
	"github.com/prometheus/client_golang/prometheus/testutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

// TestTerminalOutcomeUsesBoundedValues verifies terminal metrics use the documented outcome vocabulary.
func TestTerminalOutcomeUsesBoundedValues(t *testing.T) {
	tests := map[string]string{
		status.ReasonCompleted: outcomeSuccess,
		status.ReasonFailed:    outcomeFailure,
		status.ReasonTimedOut:  outcomeTimeout,
		"unexpected":           outcomeFailure,
	}
	for reason, want := range tests {
		if got := terminalOutcome(reason); got != want {
			t.Errorf("terminalOutcome(%q) = %q, want %q", reason, got, want)
		}
	}
}

// TestLabelsUseBoundedValues verifies API version and cluster labels are normalized.
func TestLabelsUseBoundedValues(t *testing.T) {
	for _, version := range []string{"v1alpha1", "v1beta1"} {
		if got := apiVersion(version); got != version {
			t.Errorf("apiVersion(%q) = %q, want %q", version, got, version)
		}
	}
	if got := apiVersion("unexpected"); got != unknownLabel {
		t.Errorf("apiVersion(unexpected) = %q, want %q", got, unknownLabel)
	}
	if got := clusterName(""); got != unknownLabel {
		t.Errorf("clusterName(\"\") = %q, want %q", got, unknownLabel)
	}
}

// TestBuildInfoExposesVersion verifies the release version is exposed once through build info.
func TestBuildInfoExposesVersion(t *testing.T) {
	if got := testutil.ToFloat64(buildInfo); got != 1 {
		t.Errorf("etos_controller_build_info = %v, want 1", got)
	}
}

// TestTestRunStartInitializesOutcomes verifies start records one sample and creates zero-valued outcome series.
func TestTestRunStartInitializesOutcomes(t *testing.T) {
	testrun := newTestRun("init-cluster", "v1alpha1", nil, nil)
	recordTestRunStarted(testrun)
	if got := testutil.ToFloat64(testRunsStarted.WithLabelValues("v1alpha1", "init-cluster")); got != 1 {
		t.Errorf("started = %v, want 1", got)
	}
	for _, outcome := range testRunOutcomes {
		if got := testutil.ToFloat64(testRunsTerminal.WithLabelValues("v1alpha1", "init-cluster", outcome)); got != 0 {
			t.Errorf("terminal %s = %v, want 0", outcome, got)
		}
	}
}

// TestTestRunTerminalCountsOnlyTerminalReasons verifies terminal metrics ignore non-terminal reasons.
func TestTestRunTerminalCountsOnlyTerminalReasons(t *testing.T) {
	start := metav1.NewTime(time.Now().Add(-time.Minute))
	testrun := newTestRun("terminal-cluster", "v1beta1", &start, nil)
	recordTestRunTerminal(testrun, status.ReasonActive, time.Now())
	recordTestRunTerminal(testrun, status.ReasonTimedOut, time.Now())
	if got := testutil.ToFloat64(testRunsTerminal.WithLabelValues("v1beta1", "terminal-cluster", outcomeTimeout)); got != 1 {
		t.Errorf("terminal timeout = %v, want 1", got)
	}
	if got := testutil.CollectAndCount(testRunsDuration, "etos_controller_testrun_duration_seconds"); got == 0 {
		t.Error("expected a duration observation")
	}
}

// TestStateCollectorRebuildsStateAfterRestart verifies a fresh collector derives gauges from persisted resources.
func TestStateCollectorRebuildsStateAfterRestart(t *testing.T) {
	older := metav1.NewTime(time.Unix(1000, 0))
	newer := metav1.NewTime(time.Unix(2000, 0))
	completed := metav1.NewTime(time.Unix(3000, 0))
	deleted := metav1.NewTime(time.Unix(4000, 0))
	objects := []client.Object{
		&etosv1alpha1.Cluster{ObjectMeta: metav1.ObjectMeta{Name: "idle", Namespace: "ns"}},
		withName(newTestRun("busy", "v1beta1", &older, nil), "a"),
		withName(newTestRun("busy", "v1beta1", &newer, nil), "b"),
		withName(newTestRun("busy", "v1beta1", &older, &completed), "c"),
		newReleasingRequest("r1", "id-1", "busy", &deleted),
		newReleasingRequest("r2", "id-2", "busy", &deleted),
		newOwned(&etosv1alpha1.Environment{}, "env-1", "id-1"),
		newOwned(&etosv1alpha2.Iut{}, "iut-1", "id-1"),
		newOwned(&etosv1alpha2.LogArea{}, "log-2", "id-2"),
		newOwned(&etosv1alpha2.ExecutionSpace{}, "es-unrelated", "id-3"),
	}
	collector := &stateCollector{reader: newFakeReader(t, objects...)}
	expected := `
# HELP etos_controller_environment_release_residual_resources Number of Environment, IUT, LogArea, and ExecutionSpace resources still owned by releasing EnvironmentRequests.
# TYPE etos_controller_environment_release_residual_resources gauge
etos_controller_environment_release_residual_resources{api_version="v1alpha1",cluster="idle"} 0
etos_controller_environment_release_residual_resources{api_version="v1beta1",cluster="busy"} 3
etos_controller_environment_release_residual_resources{api_version="v1beta1",cluster="idle"} 0
# HELP etos_controller_environment_requests_releasing Number of EnvironmentRequests that are deleted and waiting for environment release to complete.
# TYPE etos_controller_environment_requests_releasing gauge
etos_controller_environment_requests_releasing{api_version="v1alpha1",cluster="idle"} 0
etos_controller_environment_requests_releasing{api_version="v1beta1",cluster="busy"} 2
etos_controller_environment_requests_releasing{api_version="v1beta1",cluster="idle"} 0
# HELP etos_controller_testruns_active Number of TestRuns that have started and not reached a terminal state.
# TYPE etos_controller_testruns_active gauge
etos_controller_testruns_active{api_version="v1alpha1",cluster="idle"} 0
etos_controller_testruns_active{api_version="v1beta1",cluster="busy"} 2
etos_controller_testruns_active{api_version="v1beta1",cluster="idle"} 0
# HELP etos_controller_testruns_oldest_active_start_timestamp_seconds Unix start time of the oldest active TestRun. Use time() minus this value for its age.
# TYPE etos_controller_testruns_oldest_active_start_timestamp_seconds gauge
etos_controller_testruns_oldest_active_start_timestamp_seconds{api_version="v1beta1",cluster="busy"} 1000
`
	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		t.Error(err)
	}
}

// TestStateCollectorReportsOnlyOnLeader verifies non-leader replicas do not report duplicate state.
func TestStateCollectorReportsOnlyOnLeader(t *testing.T) {
	start := metav1.NewTime(time.Unix(1000, 0))
	elected := make(chan struct{})
	collector := &stateCollector{
		reader:  newFakeReader(t, withName(newTestRun("busy", "v1beta1", &start, nil), "a")),
		elected: elected,
	}
	if got := testutil.CollectAndCount(collector); got != 0 {
		t.Errorf("non-leader reported %d metrics, want 0", got)
	}
	close(elected)
	if got := testutil.CollectAndCount(collector); got == 0 {
		t.Error("leader reported no metrics")
	}
}

// newFakeReader returns a fake cache reader containing objects.
func newFakeReader(t *testing.T, objects ...client.Object) client.Reader {
	t.Helper()
	scheme := runtime.NewScheme()
	if err := etosv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}
	if err := etosv1alpha2.AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}
	return fake.NewClientBuilder().WithScheme(scheme).WithObjects(objects...).Build()
}

// newTestRun returns a TestRun with the provided lifecycle timestamps.
func newTestRun(cluster, schemaVersion string, start, completion *metav1.Time) *etosv1alpha1.TestRun {
	testrun := &etosv1alpha1.TestRun{ObjectMeta: metav1.ObjectMeta{Name: "testrun", Namespace: "ns"}}
	testrun.Spec.Cluster = cluster
	testrun.Spec.SchemaVersion = schemaVersion
	testrun.Status.StartTime = start
	testrun.Status.CompletionTime = completion
	return testrun
}

// withName sets the name of a TestRun.
func withName(testrun *etosv1alpha1.TestRun, name string) *etosv1alpha1.TestRun {
	testrun.Name = name
	return testrun
}

// newReleasingRequest returns a deleted EnvironmentRequest still holding the release finalizer.
func newReleasingRequest(name, identifier, cluster string, deleted *metav1.Time) *etosv1alpha1.EnvironmentRequest {
	environmentrequest := &etosv1alpha1.EnvironmentRequest{ObjectMeta: metav1.ObjectMeta{
		Name:              name,
		Namespace:         "ns",
		Labels:            map[string]string{clusterLabel: cluster},
		Finalizers:        []string{releaseFinalizer},
		DeletionTimestamp: deleted,
	}}
	environmentrequest.Spec.Identifier = identifier
	environmentrequest.Spec.SchemaVersion = "v1beta1"
	return environmentrequest
}

// newOwned sets the metadata of a release-tracked resource.
func newOwned(obj client.Object, name, identifier string) client.Object {
	obj.SetName(name)
	obj.SetNamespace("ns")
	obj.SetLabels(map[string]string{identifierLabel: identifier})
	return obj
}
