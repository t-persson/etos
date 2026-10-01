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
# Environment Provider metrics

The Environment Provider is the container in the controller's Environment Provider Job that creates the `Environment` resources after the IUT, Log Area, and Execution Space providers have produced their data.
It exports OpenTelemetry metrics over OTLP to the collector configured for the ETOS `Cluster` (`spec.openTelemetry`).
When OpenTelemetry is disabled, no metrics are exported.

These metrics are alpha: names, attributes, and semantics may change between ETOS releases.
They complement the [controller metrics](metrics.md), which own the EnvironmentRequest lifecycle and Job outcome; the Environment Provider metrics describe what happens inside the Environment Provider.

## Semantics

- **Export**: the provider is short-lived, so its final values are flushed when it exits, whether provisioning succeeded, failed, or timed out.
  Counters are cumulative per Environment Provider process; aggregate them across processes with a sum.
- **One record per attempt**: each provisioning attempt records exactly one `provisioning.duration` observation. A panic is recorded as a `failure` in the stage that panicked before the provider exits.
- **Outcomes**: `etos.outcome` is `success`, `failure`, or `timeout`. A provisioning attempt that exceeds the EnvironmentRequest deadline is a `timeout`.
- **Stages**: on `failure` or `timeout`, `etos.stage` is the stage that failed:
    - `resource_lookup`: reading the IUT, Log Area, and Execution Space resources.
    - `capacity_check`: fewer resources were provided than the minimum number of environments required.
    - `environment_creation`: creating an `Environment` resource.
- **Attributes** are bounded. `etos.api_version` is `v1alpha1`, `v1beta1`, or `unknown`, taken from the EnvironmentRequest schema version.
  Resource names, identifiers, provider data, and error messages are never used as attributes.
- **Version**: the ETOS release is the `etos.version` resource attribute, and the provider's own version is `service.version`. They are not metric attributes.
- **Zero versus missing**: each provisioning attempt reports `environment.creations` for every outcome and `environments.ready`, even when the value is `0`. A missing series means that no telemetry was received, not that nothing was created.

## Reference

All metrics use the instrumentation scope `github.com/eiffel-community/etos/cmd/environmentprovider`.

| Metric | Type | Unit | Attributes | Description |
| ------ | ---- | ---- | ---------- | ----------- |
| `etos.environment_provider.provisioning.duration` | histogram | `s` | `etos.api_version`, `etos.outcome`, `etos.stage` (not on success) | Duration of one provisioning attempt, from collecting the provider resources until all `Environment` resources are created or provisioning fails. The histogram count is the number of attempts per outcome. |
| `etos.environment_provider.readiness.duration` | histogram | `s` | `etos.api_version` | Time from EnvironmentRequest creation until the Environment Provider has created all of its `Environment` resources (time to ready). Recorded for successful provisioning only. |
| `etos.environment_provider.environment.creations` | counter | `{environment}` | `etos.api_version`, `etos.outcome` | `Environment` resource creation attempts. Provisioning stops at the first failed creation. |
| `etos.environment_provider.environments.requested` | counter | `{environment}` | `etos.api_version` | Minimum number of environments required by each provisioning attempt. |
| `etos.environment_provider.environments.ready` | counter | `{environment}` | `etos.api_version` | Environments handed off by successful provisioning attempts. Failed attempts add `0`. It can exceed `requested` when more resources than the minimum were provided, up to the maximum amount. |

The readiness duration ends when the `Environment` resources exist. It does not include the time for the Environment Provider Job status to reach the controller, which is part of the controller's `provisioning` stage duration.

## Interpreting the metrics

- Provisioning failure ratio: attempts with `etos.outcome` other than `success` divided by all attempts, from the `provisioning.duration` histogram count.
- Failure hotspots: group failed attempts by `etos.stage`.
- Partial provisioning: compare the `environments.ready` sum with the `environments.requested` sum, and watch `environment.creations` with `etos.outcome="failure"`.
- Time to ready: percentiles of `readiness.duration`, filtered by `etos.api_version` and `etos.version`.
