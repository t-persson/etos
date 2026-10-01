<!---
   Copyright Axis Communications AB
   For a full list of individual contributors, please see the commit history.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
--->
# Controller metrics

The ETOS controller exposes Prometheus metrics on the controller-runtime metrics endpoint, next to the standard controller-runtime metrics.
These metrics cover the TestRuns and EnvironmentRequests reconciled by the controller.

These metrics are alpha: names, labels, and semantics may change between ETOS releases.

## Semantics

- **Counters and histograms** are recorded only after a lifecycle transition has been persisted to Kubernetes, so retries and update conflicts do not count a transition twice.
  They are process-local and reset when the controller restarts. Query them with `rate()` or `increase()`, which handle counter resets.
- **State gauges** are computed at scrape time from the controller cache, so they are correct after a controller restart.
  Only the elected leader reports them, so replicas do not report duplicate values.
- **Labels** are bounded. `api_version` is `v1alpha1`, `v1beta1`, or `unknown`. `cluster` is the ETOS `Cluster` name, or `unknown` when it is missing.
  Resource identifiers and error messages are never used as label values.
- **Version** information is exposed once through `etos_controller_build_info` instead of on every series. Join on it when you need the version.

## Reference

| Metric | Type | Labels | Description |
| ------ | ---- | ------ | ----------- |
| `etos_controller_build_info` | gauge | `etos_version` | Always `1`; carries the ETOS release version of the controller. |
| `etos_controller_testruns_started_total` | counter | `api_version`, `cluster` | TestRuns whose controller lifecycle started. |
| `etos_controller_testruns_terminal_total` | counter | `api_version`, `cluster`, `outcome` | TestRuns that reached `success`, `failure`, or `timeout`. |
| `etos_controller_testrun_duration_seconds` | histogram | `api_version`, `cluster`, `outcome` | Time from TestRun start to terminal transition. |
| `etos_controller_testruns_active` | gauge | `api_version`, `cluster` | TestRuns that have started and not yet completed. |
| `etos_controller_testruns_oldest_active_start_timestamp_seconds` | gauge | `api_version`, `cluster` | Start time of the oldest active TestRun. |
| `etos_controller_environment_requests_started_total` | counter | `api_version`, `cluster` | EnvironmentRequests whose controller lifecycle started. |
| `etos_controller_environment_request_jobs_total` | counter | `api_version`, `cluster`, `stage`, `outcome` | Persisted provisioning outcomes (`stage="provisioning"`). |
| `etos_controller_environment_request_stage_duration_seconds` | histogram | `api_version`, `cluster`, `stage`, `outcome` | Provisioning duration, and TestRun time until the environment is ready (`stage="readiness"`). |
| `etos_controller_environment_release_attempts_total` | counter | `api_version`, `cluster`, `outcome` | Environment release cleanup attempts. A failure is counted for each failed attempt. |
| `etos_controller_environment_release_duration_seconds` | histogram | `api_version`, `cluster` | Time from EnvironmentRequest deletion to completed release. |
| `etos_controller_environment_requests_releasing` | gauge | `api_version`, `cluster` | Deleted EnvironmentRequests waiting for release to complete. |
| `etos_controller_environment_release_residual_resources` | gauge | `api_version`, `cluster` | Environment, IUT, LogArea, and ExecutionSpace resources still owned by releasing EnvironmentRequests. |
| `etos_controller_state_metrics_collection_errors_total` | counter | | Failed attempts to compute the state gauges. |

State gauges report an explicit `0` for every `Cluster` resource and supported API version, so an idle cluster is distinguishable from missing data.

## Example queries

TestRun failure ratio over the last hour:

```promql
sum by (cluster) (increase(etos_controller_testruns_terminal_total{outcome!="success"}[1h]))
/
sum by (cluster) (increase(etos_controller_testruns_terminal_total[1h]))
```

95th percentile TestRun duration:

```promql
histogram_quantile(0.95, sum by (le, cluster) (rate(etos_controller_testrun_duration_seconds_bucket[1h])))
```

Age of the oldest active TestRun:

```promql
time() - etos_controller_testruns_oldest_active_start_timestamp_seconds
```

Controller version per instance:

```promql
etos_controller_build_info
```
