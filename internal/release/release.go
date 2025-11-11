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

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReleaseJob returns a batchv1.Job schema populated with containers provided.
func ReleaseJob(jobName, name, namespace string, containers ...corev1.Container) *batchv1.Job {
	ttl := int32(300)
	grace := int64(30)
	backoff := int32(0)
	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/name":    name,
				"app.kubernetes.io/part-of": "etos",
			},
			Annotations: make(map[string]string),
			Name:        jobName,
			Namespace:   namespace,
		},
		Spec: batchv1.JobSpec{
			TTLSecondsAfterFinished: &ttl,
			BackoffLimit:            &backoff,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: jobName,
					Labels: map[string]string{
						"app.kubernetes.io/name":    name,
						"app.kubernetes.io/part-of": "etos",
					},
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: &grace,
					ServiceAccountName:            "provider", // TODO: Wrong service account
					RestartPolicy:                 "Never",
					Containers:                    containers,
				},
			},
		},
	}
}

// ReleaseContainer returns a container specification.
func ReleaseContainer(name, namespace, image, parameter string, noDelete bool) corev1.Container {
	args := []string{
		"-release",
		fmt.Sprintf("-namespace=%s", namespace),
		parameter,
	}
	if noDelete {
		args = append(args, "-nodelete")
	}
	return corev1.Container{
		Name:            name,
		Image:           image,
		ImagePullPolicy: corev1.PullIfNotPresent,
		// TODO: Verify these resourceclaims
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceMemory: resource.MustParse("256Mi"),
				corev1.ResourceCPU:    resource.MustParse("250m"),
			},
			Requests: corev1.ResourceList{
				corev1.ResourceMemory: resource.MustParse("128Mi"),
				corev1.ResourceCPU:    resource.MustParse("100m"),
			},
		},
		Args: args,
	}
}
