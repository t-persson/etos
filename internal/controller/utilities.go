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
	etosv1alpha2 "github.com/eiffel-community/etos/api/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	providerFinalizer       = "etos.eiffel-community.github.io/managed-by-provider"
	LogAreaOwnerKey         = ".metadata.controller.log-area-provider"
	ExecutionSpaceOwnerKey  = ".metadata.controller.execution-space-provider"
	IutOwnerKey             = ".metadata.controller.iut-provider"
	APIv2GroupVersionString = etosv1alpha2.GroupVersion.String()
)

// hasOwner checks if a resource kind exists in ownerReferences.
func hasOwner(ownerReferences []metav1.OwnerReference, kind string) bool {
	for _, ownerReference := range ownerReferences {
		if ownerReference.Kind == kind {
			return true
		}
	}
	return false
}
