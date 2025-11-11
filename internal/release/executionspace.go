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
package release

import (
	"fmt"

	"github.com/eiffel-community/etos/api/v1alpha2"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

// ExecutionSpaceReleaser returns an ExecutionSpace releaser job specification.
func ExecutionSpaceReleaser(executionSpace *v1alpha2.ExecutionSpace, image string, noDelete bool) *batchv1.Job {
	return ReleaseJob(
		executionSpace.Name,
		"execution-space-releaser",
		executionSpace.Namespace,
		ExecutionSpaceReleaserContainer(executionSpace, image, noDelete),
	)
}

// ExecutionSpaceReleaserContainer returns an ExecutionSpace releaser container specification.
func ExecutionSpaceReleaserContainer(executionSpace *v1alpha2.ExecutionSpace, image string, noDelete bool) corev1.Container {
	return ReleaseContainer(
		executionSpace.Name,
		executionSpace.Namespace,
		image,
		fmt.Sprintf("-execution-space=%s", executionSpace.Name),
		noDelete,
	)
}
