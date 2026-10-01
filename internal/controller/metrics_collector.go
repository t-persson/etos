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
	"context"
	"time"

	etosv1alpha1 "github.com/eiffel-community/etos/api/v1alpha1"
	etosv1alpha2 "github.com/eiffel-community/etos/api/v1alpha2"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

const (
	stateCollectionTimeout = 5 * time.Second
	identifierLabel        = "etos.eiffel-community.github.io/id"
)

var (
	testRunsActiveDesc = prometheus.NewDesc(
		"etos_controller_testruns_active",
		"Number of TestRuns that have started and not reached a terminal state.",
		[]string{"api_version", "cluster"}, nil,
	)
	testRunsOldestActiveStartDesc = prometheus.NewDesc(
		"etos_controller_testruns_oldest_active_start_timestamp_seconds",
		"Unix start time of the oldest active TestRun. Use time() minus this value for its age.",
		[]string{"api_version", "cluster"}, nil,
	)
	environmentRequestsReleasingDesc = prometheus.NewDesc(
		"etos_controller_environment_requests_releasing",
		"Number of EnvironmentRequests that are deleted and waiting for environment release to complete.",
		[]string{"api_version", "cluster"}, nil,
	)
	environmentResidualResourcesDesc = prometheus.NewDesc(
		"etos_controller_environment_release_residual_resources",
		"Number of Environment, IUT, LogArea, and ExecutionSpace resources still owned by releasing EnvironmentRequests.",
		[]string{"api_version", "cluster"}, nil,
	)
	stateCollectionErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "etos_controller_state_metrics_collection_errors_total",
			Help: "Number of failed attempts to compute controller state metrics from the informer cache.",
		},
	)
)

// stateKey is the bounded label set of a state gauge.
type stateKey struct {
	apiVersion string
	cluster    string
}

// stateCollector computes current-state gauges from the informer cache at scrape time.
// Deriving gauges from Kubernetes resources, rather than incrementing them on transitions,
// keeps them correct across controller restarts, leader changes, and missed events.
type stateCollector struct {
	reader  client.Reader
	elected <-chan struct{}
}

// RegisterStateMetrics registers the controller state metrics collector with the controller-runtime registry.
// State is reported only by the elected leader so that replicas do not report duplicate values.
func RegisterStateMetrics(mgr ctrl.Manager) error {
	if err := metrics.Registry.Register(stateCollectionErrors); err != nil {
		return err
	}
	return metrics.Registry.Register(&stateCollector{reader: mgr.GetClient(), elected: mgr.Elected()})
}

// Describe implements prometheus.Collector.
func (c *stateCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- testRunsActiveDesc
	ch <- testRunsOldestActiveStartDesc
	ch <- environmentRequestsReleasingDesc
	ch <- environmentResidualResourcesDesc
}

// Collect implements prometheus.Collector.
func (c *stateCollector) Collect(ch chan<- prometheus.Metric) {
	if !c.isLeader() {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), stateCollectionTimeout)
	defer cancel()
	state, err := c.collect(ctx)
	if err != nil {
		stateCollectionErrors.Inc()
		return
	}
	state.emit(ch)
}

// isLeader reports whether this controller instance is the elected leader.
func (c *stateCollector) isLeader() bool {
	if c.elected == nil {
		return true
	}
	select {
	case <-c.elected:
		return true
	default:
		return false
	}
}

// controllerState holds aggregated state gauge values.
type controllerState struct {
	keys              map[stateKey]struct{}
	activeTestRuns    map[stateKey]int
	oldestActiveStart map[stateKey]time.Time
	releasing         map[stateKey]int
	residual          map[stateKey]int
}

// newControllerState creates an empty controllerState.
func newControllerState() *controllerState {
	return &controllerState{
		keys:              map[stateKey]struct{}{},
		activeTestRuns:    map[stateKey]int{},
		oldestActiveStart: map[stateKey]time.Time{},
		releasing:         map[stateKey]int{},
		residual:          map[stateKey]int{},
	}
}

// collect lists resources from the cache and aggregates state gauges.
func (c *stateCollector) collect(ctx context.Context) (*controllerState, error) {
	var clusters etosv1alpha1.ClusterList
	if err := c.reader.List(ctx, &clusters); err != nil {
		return nil, err
	}
	var testruns etosv1alpha1.TestRunList
	if err := c.reader.List(ctx, &testruns); err != nil {
		return nil, err
	}
	var environmentrequests etosv1alpha1.EnvironmentRequestList
	if err := c.reader.List(ctx, &environmentrequests); err != nil {
		return nil, err
	}
	owned, err := c.ownedResources(ctx)
	if err != nil {
		return nil, err
	}
	return aggregateState(clusters.Items, testruns.Items, environmentrequests.Items, owned), nil
}

// ownedResources counts release-tracked resources by namespace and EnvironmentRequest identifier.
func (c *stateCollector) ownedResources(ctx context.Context) (map[string]int, error) {
	owned := map[string]int{}
	lists := []client.ObjectList{
		&etosv1alpha1.EnvironmentList{},
		&etosv1alpha2.IutList{},
		&etosv1alpha2.LogAreaList{},
		&etosv1alpha2.ExecutionSpaceList{},
	}
	for _, list := range lists {
		if err := c.reader.List(ctx, list, client.HasLabels{identifierLabel}); err != nil {
			return nil, err
		}
		if err := meta.EachListItem(list, func(obj runtime.Object) error {
			accessor, err := meta.Accessor(obj)
			if err != nil {
				return err
			}
			owned[accessor.GetNamespace()+"/"+accessor.GetLabels()[identifierLabel]]++
			return nil
		}); err != nil {
			return nil, err
		}
	}
	return owned, nil
}

// aggregateState computes state gauge values from listed resources.
func aggregateState(
	clusters []etosv1alpha1.Cluster,
	testruns []etosv1alpha1.TestRun,
	environmentrequests []etosv1alpha1.EnvironmentRequest,
	owned map[string]int,
) *controllerState {
	state := newControllerState()
	for _, cluster := range clusters {
		for _, version := range supportedAPIVersions {
			state.keys[stateKey{apiVersion: version, cluster: clusterName(cluster.Name)}] = struct{}{}
		}
	}
	for i := range testruns {
		testrun := &testruns[i]
		if !isTestRunActive(testrun) {
			continue
		}
		apiVersionLabel, cluster := testRunLabels(testrun)
		key := stateKey{apiVersion: apiVersionLabel, cluster: cluster}
		state.keys[key] = struct{}{}
		state.activeTestRuns[key]++
		start := testrun.Status.StartTime.Time
		if oldest, ok := state.oldestActiveStart[key]; !ok || start.Before(oldest) {
			state.oldestActiveStart[key] = start
		}
	}
	for i := range environmentrequests {
		environmentrequest := &environmentrequests[i]
		if environmentrequest.DeletionTimestamp.IsZero() ||
			!controllerutil.ContainsFinalizer(environmentrequest, releaseFinalizer) {
			continue
		}
		apiVersionLabel, cluster := environmentRequestLabels(environmentrequest)
		key := stateKey{apiVersion: apiVersionLabel, cluster: cluster}
		state.keys[key] = struct{}{}
		state.releasing[key]++
		state.residual[key] += owned[environmentrequest.Namespace+"/"+environmentrequest.Spec.Identifier]
	}
	return state
}

// emit sends state gauge values for every known label set, including explicit zeros.
func (s *controllerState) emit(ch chan<- prometheus.Metric) {
	for key := range s.keys {
		ch <- prometheus.MustNewConstMetric(testRunsActiveDesc, prometheus.GaugeValue,
			float64(s.activeTestRuns[key]), key.apiVersion, key.cluster)
		ch <- prometheus.MustNewConstMetric(environmentRequestsReleasingDesc, prometheus.GaugeValue,
			float64(s.releasing[key]), key.apiVersion, key.cluster)
		ch <- prometheus.MustNewConstMetric(environmentResidualResourcesDesc, prometheus.GaugeValue,
			float64(s.residual[key]), key.apiVersion, key.cluster)
		if oldest, ok := s.oldestActiveStart[key]; ok {
			ch <- prometheus.MustNewConstMetric(testRunsOldestActiveStartDesc, prometheus.GaugeValue,
				float64(oldest.Unix()), key.apiVersion, key.cluster)
		}
	}
}
