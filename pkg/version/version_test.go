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
package version

import "testing"

func TestCurrentPrefersEnvironment(t *testing.T) {
	original := Version
	t.Cleanup(func() { Version = original })
	Version = "embedded"

	t.Setenv(EnvironmentVariable, "")
	if got := Current(); got != "embedded" {
		t.Fatalf("Current() = %q without %s, want %q", got, EnvironmentVariable, "embedded")
	}
	t.Setenv(EnvironmentVariable, "9.1.0")
	if got := Current(); got != "9.1.0" {
		t.Fatalf("Current() = %q with %s set, want %q", got, EnvironmentVariable, "9.1.0")
	}
}
