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
package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/eiffel-community/etos/api/v1alpha1"
	"github.com/eiffel-community/etos/api/v1alpha2"
	"github.com/eiffel-community/etos/internal/controller/jobs"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// RunExecutionSpaceProvider is the base runner for an ExecutionSpace provider. Checks input parameters and calls either Release or Provision on a Provider.
//
// This function panics on errors, propagating errors back to the controller that executed it.
func RunExecutionSpaceProvider(provider Provider) {
	params := loadParameters()

	if params.releaseEnvironment {
		if err := runExecutionSpaceReleaser(params, provider); err != nil {
			if err := WriteResult(params.logger,
				jobs.Result{
					Conclusion:  jobs.ConclusionFailed,
					Description: err.Error(),
					Verdict:     jobs.VerdictNone,
				}); err != nil {
				params.logger.Error(err, "failed to write error result to termination-log")
			}
			panic(err)
		}
		if err := WriteResult(params.logger,
			jobs.Result{
				Conclusion:  jobs.ConclusionSuccessful,
				Description: "Successfully released ExecutionSpace",
				Verdict:     jobs.VerdictNone,
			}); err != nil {
			params.logger.Error(err, "failed to write error result to termination-log")
			panic(err)
		}
	} else {
		if err := runExecutionSpaceProvider(params, provider); err != nil {
			if err := WriteResult(params.logger,
				jobs.Result{
					Conclusion:  jobs.ConclusionFailed,
					Description: err.Error(),
					Verdict:     jobs.VerdictNone,
				}); err != nil {
				params.logger.Error(err, "failed to write error result to termination-log")
			}
			panic(err)
		}
		if err := WriteResult(params.logger,
			jobs.Result{
				Conclusion:  jobs.ConclusionSuccessful,
				Description: "Successfully provisioned ExecutionSpaces",
				Verdict:     jobs.VerdictNone,
			}); err != nil {
			params.logger.Error(err, "failed to write error result to termination-log")
			panic(err)
		}
	}
}

// runExecutionSpaceReleaser is the base releaser for a ExecutionSpace provider. Checks input parameters and calls Release on a Provider.
func runExecutionSpaceReleaser(params parameters, provider Provider) error {
	ctx := context.Background()
	logger := params.logger
	if params.namespace == "" {
		return errors.New("Must set -namespace")
	}
	if params.name == "" {
		return errors.New("Must set -name")
	}
	return provider.Release(ctx, logger, ReleaseConfig{
		Name:      params.name,
		Namespace: params.namespace,
		NoDelete:  params.noDelete,
	})
}

// runExecutionSpaceProvider is the base runner for a ExecutionSpace provider. Checks input parameters and calls Provision on a Provider.
func runExecutionSpaceProvider(params parameters, provider Provider) error {
	ctx := context.Background()
	logger := params.logger
	if params.namespace == "" {
		return errors.New("Must set -namespace")
	}
	if params.environmentRequestName == "" {
		return errors.New("Must set -environment-request")
	}
	cli, err := kubernetesClient()
	if err != nil {
		return err
	}
	var request v1alpha1.EnvironmentRequest
	if err := cli.Get(ctx, types.NamespacedName{Name: params.environmentRequestName, Namespace: params.namespace}, &request); err != nil {
		return err
	}
	iuts, err := GetIUTs(ctx, request.Spec.ID, params.namespace)
	if err != nil {
		return err
	}
	return provider.Provision(ctx, logger, ProvisionConfig{
		EnvironmentRequest: &request,
		Namespace:          params.namespace,
		MaximumAmount:      request.Spec.MaximumAmount,
		MinimumAmount:      len(iuts.Items),
	})
}

// GetExecutionSpace gets an ExecutionSpace resource by name from Kubernetes.
func GetExecutionSpace(ctx context.Context, name, namespace string) (*v1alpha2.ExecutionSpace, error) {
	cli, err := kubernetesClient()
	if err != nil {
		return nil, err
	}
	var executionSpace v1alpha2.ExecutionSpace
	if err := cli.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, &executionSpace); err != nil {
		return nil, err
	}
	return &executionSpace, nil
}

// CreateExecutionSpace creates a new ExecutionSpace resource in Kubernetes.
//
// The spec.ID and spec.ProviderID fields are automatically populated by this function. They will be overwritten if set.
func CreateExecutionSpace(ctx context.Context, environmentrequest *v1alpha1.EnvironmentRequest, namespace string, spec v1alpha2.ExecutionSpaceSpec) error {
	logger, _ := logr.FromContext(ctx)

	logger.Info("Getting Kubernetes client")
	cli, err := kubernetesClient()
	if err != nil {
		return err
	}
	labels := map[string]string{
		"etos.eiffel-community.github.io/provider":               environmentrequest.Spec.Providers.ExecutionSpace.ID,
		"etos.eiffel-community.github.io/environment-request":    environmentrequest.Spec.Name,
		"etos.eiffel-community.github.io/environment-request-id": environmentrequest.Spec.ID,
		"app.kubernetes.io/name":                                 "execution-space-provider",
		"app.kubernetes.io/part-of":                              "etos",
	}
	if cluster := environmentrequest.Labels["etos.eiffel-community.github.io/cluster"]; cluster != "" {
		labels["etos.eiffel-community.github.io/cluster"] = cluster
	}
	if environmentrequest.Spec.Identifier != "" {
		labels["etos.eiffel-community.github.io/id"] = environmentrequest.Spec.Identifier
	}

	spec.ID = uuid.NewString()
	spec.ProviderID = environmentrequest.Spec.Providers.ExecutionSpace.ID

	isController := false
	blockOwnerDeletion := true
	return cli.Create(ctx, &v1alpha2.ExecutionSpace{
		ObjectMeta: metav1.ObjectMeta{
			Labels:       labels,
			GenerateName: fmt.Sprintf("%s-execution-space-", strings.ToLower(environmentrequest.Spec.Name)),
			Namespace:    namespace,
			OwnerReferences: []metav1.OwnerReference{{
				Kind:               "EnvironmentRequest",
				Name:               environmentrequest.GetName(),
				UID:                environmentrequest.GetUID(),
				APIVersion:         environmentrequest.GetResourceVersion(),
				Controller:         &isController,
				BlockOwnerDeletion: &blockOwnerDeletion,
			}},
		},
		Spec: spec,
	})
}

// DeleteExecutionSpace deletes an ExecutionSpace resource from Kubernetes.
func DeleteExecutionSpace(ctx context.Context, executionSpace *v1alpha2.ExecutionSpace) error {
	cli, err := kubernetesClient()
	if err != nil {
		return err
	}
	return cli.Delete(ctx, executionSpace)
}
