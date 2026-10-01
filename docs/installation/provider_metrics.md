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
# Provider metrics

The default IUT, Log Area, and Execution Space providers export OpenTelemetry metrics over OTLP/gRPC to the collector configured in the ETOS `Cluster` resource (`spec.openTelemetry`).
When OpenTelemetry is disabled, no provider metrics are exported.

These metrics cover the v1alpha and v1beta provider Jobs started by the ETOS controller, and any custom Go provider built with the `pkg/provider` package (`RunIutProvider`, `RunLogAreaProvider`, or `RunExecutionSpaceProvider`).
The Environment Provider records its own metrics and is not covered here, and provider Job and EnvironmentRequest outcomes are covered by the [controller metrics](./metrics.md).

These metrics are alpha: names, attributes, and semantics may change between ETOS releases.

## Semantics

- **One attempt per process.** Each provider container runs a single provisioning or release attempt and records it exactly once, when the result is known.
  This includes startup failures, timeouts, and panics. The metrics are flushed when the provider exits.
- **No retries.** The providers do not retry, and their Jobs are not restarted, so there is no retry metric.
  A new attempt is a new provider process and is counted as a separate attempt.
- **Durations** are measured in seconds from provider start until the result is known, including writing the result to the termination-log.
- **Zero versus missing.** Attempt counters add `0` for the outcomes that did not occur, so a provider that reports only successes shows an explicit zero failure count.
  A missing series means that no telemetry arrived.
- **Attributes** are bounded. Resource names, provider values, and error messages are never used as attribute values.
- **Version.** The ETOS release is the `etos.version` OTLP resource attribute on every metric, set by the controller through the `ETOS_VERSION` environment variable.
  The provider's own build version is `service.version`.

### Attributes

| Attribute | Values | Description |
| --------- | ------ | ----------- |
| `etos.provider.type` | `iut`, `log_area`, `execution_space` | The provider type. |
| `etos.api_version` | `v1alpha1`, `v1beta1`, `unknown` | The schema version of the EnvironmentRequest. `unknown` when it is missing, unsupported, or the EnvironmentRequest could not be read. |
| `etos.outcome` | `success`, `failure`, `timeout` | The normalized result. `timeout` means that the EnvironmentRequest deadline or the wait deadline was exceeded. |
| `etos.stage` | `startup`, `operation`, `result` | Duration histograms of unsuccessful attempts only. Where the attempt failed: before the operation started (for example reading the EnvironmentRequest or connecting to the message bus), in the provisioning or release operation, or while writing the result. |
| `error.type` | `test_runner_error`, `shutdown`, `stream_closed`, `stream_error`, `timeout`, `canceled` | Test Runner readiness duration histogram of unsuccessful waits only. The Test Runner reported an error, a shutdown event arrived, the event stream ended or failed, the wait timed out, or the wait was canceled. |

## Reference

| Metric | Type | Unit | Attributes | Description |
| ------ | ---- | ---- | ---------- | ----------- |
| `etos.provider.provisioning.attempts` | counter | `{attempt}` | `etos.provider.type`, `etos.api_version`, `etos.outcome` | Provisioning attempts by outcome. |
| `etos.provider.provisioning.duration` | histogram | `s` | `etos.provider.type`, `etos.api_version`, `etos.outcome`, `etos.stage` | Duration of provisioning attempts. |
| `etos.provider.release.attempts` | counter | `{attempt}` | `etos.provider.type`, `etos.api_version`, `etos.outcome` | Release attempts by outcome. |
| `etos.provider.release.duration` | histogram | `s` | `etos.provider.type`, `etos.api_version`, `etos.outcome`, `etos.stage` | Duration of release attempts. |
| `etos.provider.test_runner.readiness.attempts` | counter | `{attempt}` | `etos.api_version`, `etos.outcome` | Test Runner readiness waits by the Execution Space Provider. |
| `etos.provider.test_runner.readiness.duration` | histogram | `s` | `etos.api_version`, `etos.outcome`, `error.type` | Time from the start of a readiness wait until the Test Runner reports its status or the wait fails. |

The Execution Space Provider starts each Test Runner Job and waits for the Test Runner to report its status through the ETOS SSE API, so it owns the Test Runner readiness observation.
It records one readiness wait per Test Runner. Test Runner execution is not covered by these metrics.

Duration histograms use the bucket boundaries 0.5, 1, 2.5, 5, 10, 30, 60, 120, 300, 600, 900, 1800, and 3600 seconds.
