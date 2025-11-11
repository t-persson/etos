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
package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/eiffel-community/etos/api/v1alpha2"
	"github.com/eiffel-community/etos/pkg/provider"
	"github.com/fernet/fernet-go"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
)

type genericExecutionSpaceProvider struct{}

// main creates a new ExecutionSpace resource based on data in an EnvironmentRequest.
func main() {
	provider.RunExecutionSpaceProvider(&genericExecutionSpaceProvider{})
}

// Provision provisions a new ExecutionSpace.
// TODO: This does not work properly at the moment. Need to launch an ETR job as well.
func (p *genericExecutionSpaceProvider) Provision(ctx context.Context, logger logr.Logger, cfg provider.ProvisionConfig) error {
	environmentRequest := cfg.EnvironmentRequest
	logger.Info("Provisioning a new ExecutionSpace for EnvironmentRequest", "EnvironmentRequest", environmentRequest.Name, "Namespace", environmentRequest.Namespace)
	if cfg.MinimumAmount <= 0 {
		return errors.New("minimum amount of ExecutionSpaces requested is less than or equal to 0")
	}
	encryptionKey, err := fernet.DecodeKey(environmentRequest.Spec.Config.EncryptionKey.Value)
	if err != nil {
		return err
	}
	id := uuid.NewString()
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	for range cfg.MinimumAmount {
		if err := provider.CreateExecutionSpace(ctx, environmentRequest, cfg.Namespace, v1alpha2.ExecutionSpaceSpec{
			Instructions: v1alpha2.Instructions{
				Identifier: environmentRequest.Spec.Identifier,
				Image:      "ghcr.io/eiffel-community/etos-base-test-runner:bullseye", // TODO:
				Parameters: map[string]string{},
				Environment: map[string]string{
					"ENVIRONMENT_ID":         id,
					"ENVIRONMENT_URL":        fmt.Sprintf("%s/v1alpha/testrun/%s", environmentRequest.Spec.Config.EtosApi, id),
					"SOURCE_HOST":            hostname,
					"ETOS_API":               environmentRequest.Spec.Config.EtosApi,
					"ETR_VERSION":            environmentRequest.Spec.Providers.ExecutionSpace.TestRunner,
					"ETOS_GRAPHQL_SERVER":    environmentRequest.Spec.Config.GraphQlServer,
					"ETOS_RABBITMQ_EXCHANGE": environmentRequest.Spec.Config.EtosMessageBus.Exchange,
					"ETOS_RABBITMQ_HOST":     environmentRequest.Spec.Config.EtosMessageBus.Host,
					"ETOS_RABBITMQ_PASSWORD": string(encrypt(environmentRequest.Spec.Config.EtosMessageBus.Password.Value, encryptionKey)),
					"ETOS_RABBITMQ_PORT":     environmentRequest.Spec.Config.EtosMessageBus.Port,
					"ETOS_RABBITMQ_USERNAME": environmentRequest.Spec.Config.EtosMessageBus.Username,
					"ETOS_RABBITMQ_VHOST":    environmentRequest.Spec.Config.EtosMessageBus.Vhost,
					"ETOS_RABBITMQ_SSL":      environmentRequest.Spec.Config.EtosMessageBus.SSL,
					"RABBITMQ_EXCHANGE":      environmentRequest.Spec.Config.EiffelMessageBus.Exchange,
					"RABBITMQ_HOST":          environmentRequest.Spec.Config.EiffelMessageBus.Host,
					"RABBITMQ_PASSWORD":      string(encrypt(environmentRequest.Spec.Config.EtosMessageBus.Password.Value, encryptionKey)),
					"RABBITMQ_PORT":          environmentRequest.Spec.Config.EiffelMessageBus.Port,
					"RABBITMQ_USERNAME":      environmentRequest.Spec.Config.EiffelMessageBus.Username,
					"RABBITMQ_VHOST":         environmentRequest.Spec.Config.EiffelMessageBus.Vhost,
					"RABBITMQ_SSL":           environmentRequest.Spec.Config.EiffelMessageBus.SSL,
					// TODO: This is extra for testing.
					"DEV":            "true",
					"ETR_BRANCH":     "waiting-etr",
					"ETR_REPOSITORY": "https://github.com/t-persson/etos-test-runner",
				},
			},
		}); err != nil {
			return err
		}
	}
	return nil
}

// Release releases an ExecutionSpace.
func (p *genericExecutionSpaceProvider) Release(ctx context.Context, logger logr.Logger, cfg provider.ReleaseConfig) error {
	logger.Info("Releasing ExecutionSpace", "Name", cfg.Name, "Namespace", cfg.Namespace)
	executionSpace, err := provider.GetExecutionSpace(ctx, cfg.Name, cfg.Namespace)
	if err != nil {
		return err
	}
	logger.Info("ExecutionSpace", "name", executionSpace.Name)
	if cfg.NoDelete {
		return nil
	} else {
		return provider.DeleteExecutionSpace(ctx, executionSpace)
	}
}

func encrypt(s string, key *fernet.Key) []byte {
	e, err := fernet.EncryptAndSign([]byte(s), key)
	if err != nil {
		panic(err)
	}
	return e
}
