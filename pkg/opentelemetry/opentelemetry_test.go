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
package opentelemetry

import (
	"context"
	"testing"

	"github.com/eiffel-community/etos/pkg/version"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"
)

func TestResourceCarriesETOSVersion(t *testing.T) {
	t.Setenv(version.EnvironmentVariable, "9.1.0")
	res, err := New("provider", "iut").newOtelResource(context.Background())
	if err != nil {
		t.Fatalf("newOtelResource() error = %v", err)
	}
	got, ok := res.Set().Value(attribute.Key("etos.version"))
	if !ok || got.AsString() != "9.1.0" {
		t.Fatalf("etos.version = %q (present %t), want %q", got.AsString(), ok, "9.1.0")
	}
}

func TestShutdownAfterPartialStart(t *testing.T) {
	tracer := &ETOSTracer{enabled: true}
	if err := tracer.Shutdown(context.Background()); err != nil {
		t.Fatalf("Shutdown() with no initialized providers error = %v", err)
	}
	tracer.tracerProvider = trace.NewTracerProvider()
	if err := tracer.Shutdown(context.Background()); err != nil {
		t.Fatalf("Shutdown() with only a tracer provider error = %v", err)
	}
}
