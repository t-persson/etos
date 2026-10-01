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
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
)

// TestUpsertEnv verifies that overrides win by name without duplicates or modifying the input.
func TestUpsertEnv(t *testing.T) {
	secret := &corev1.EnvVarSource{SecretKeyRef: &corev1.SecretKeySelector{Key: "key"}}
	tests := []struct {
		name      string
		base      []corev1.EnvVar
		overrides []corev1.EnvVar
		want      []corev1.EnvVar
	}{
		{
			name:      "empty base",
			overrides: []corev1.EnvVar{{Name: "ETOS_VERSION", Value: "1.0.0"}},
			want:      []corev1.EnvVar{{Name: "ETOS_VERSION", Value: "1.0.0"}},
		},
		{
			name:      "no overlap appends overrides",
			base:      []corev1.EnvVar{{Name: "A", Value: "a"}},
			overrides: []corev1.EnvVar{{Name: "ETOS_VERSION", Value: "1.0.0"}},
			want:      []corev1.EnvVar{{Name: "A", Value: "a"}, {Name: "ETOS_VERSION", Value: "1.0.0"}},
		},
		{
			name: "override replaces in place",
			base: []corev1.EnvVar{
				{Name: "A", Value: "a"},
				{Name: "ETOS_VERSION", Value: "provider"},
				{Name: "B", Value: "$(ETOS_VERSION)"},
			},
			overrides: []corev1.EnvVar{{Name: "ETOS_VERSION", Value: "1.0.0"}, {Name: "C", Value: "c"}},
			want: []corev1.EnvVar{
				{Name: "A", Value: "a"},
				{Name: "ETOS_VERSION", Value: "1.0.0"},
				{Name: "B", Value: "$(ETOS_VERSION)"},
				{Name: "C", Value: "c"},
			},
		},
		{
			name:      "override replaces valueFrom",
			base:      []corev1.EnvVar{{Name: "ETOS_VERSION", ValueFrom: secret}},
			overrides: []corev1.EnvVar{{Name: "ETOS_VERSION", Value: "1.0.0"}},
			want:      []corev1.EnvVar{{Name: "ETOS_VERSION", Value: "1.0.0"}},
		},
		{
			name: "duplicate base names collapse into the override",
			base: []corev1.EnvVar{
				{Name: "ETOS_VERSION", Value: "first"},
				{Name: "A", Value: "a"},
				{Name: "ETOS_VERSION", Value: "second"},
			},
			overrides: []corev1.EnvVar{{Name: "ETOS_VERSION", Value: "1.0.0"}},
			want:      []corev1.EnvVar{{Name: "ETOS_VERSION", Value: "1.0.0"}, {Name: "A", Value: "a"}},
		},
		{
			name: "base duplicates without override are kept",
			base: []corev1.EnvVar{{Name: "A", Value: "1"}, {Name: "A", Value: "2"}},
			want: []corev1.EnvVar{{Name: "A", Value: "1"}, {Name: "A", Value: "2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base := make([]corev1.EnvVar, len(tt.base), len(tt.base)+len(tt.overrides))
			copy(base, tt.base)
			got := upsertEnv(base, tt.overrides)
			if !sameEnv(got, tt.want) {
				t.Fatalf("upsertEnv() = %v, want %v", got, tt.want)
			}
			if !sameEnv(base, tt.base) {
				t.Fatalf("upsertEnv() modified base: %v, want %v", base, tt.base)
			}
		})
	}
}

// sameEnv compares environment variable lists, treating nil and empty lists as equal.
func sameEnv(a, b []corev1.EnvVar) bool {
	return len(a) == len(b) && (len(a) == 0 || equality.Semantic.DeepEqual(a, b))
}
