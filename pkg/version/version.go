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

// Package version contains the ETOS release version embedded in each binary.
package version

import "os"

// Version is set by the release build and is intentionally not derived from a
// commit SHA.
var Version = "unknown"

// EnvironmentVariable is the name used to propagate the ETOS release version.
const EnvironmentVariable = "ETOS_VERSION"

// Current returns the ETOS release version for this process. The ETOS_VERSION
// environment variable, set by the ETOS controller on the workloads it creates,
// takes precedence over the version embedded at build time, because workload
// images such as providers are not built with an embedded version.
func Current() string {
	if v := os.Getenv(EnvironmentVariable); v != "" {
		return v
	}
	return Version
}
