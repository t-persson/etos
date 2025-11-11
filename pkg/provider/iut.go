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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// RunIutProvider is the base runner for an IUT provider. Checks input parameters and calls either Release or Provision on a Provider.
//
// This function panics on errors, propagating errors back to the controller that executed it.
func RunIutProvider(provider Provider) {
	params := loadParameters()

	if params.releaseEnvironment {
		if err := runIutReleaser(params, provider); err != nil {
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
				Description: "Successfully released IUT",
				Verdict:     jobs.VerdictNone,
			}); err != nil {
			params.logger.Error(err, "failed to write error result to termination-log")
			panic(err)
		}
	} else {
		if err := runIutProvider(params, provider); err != nil {
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
				Description: "Successfully provisioned IUTs",
				Verdict:     jobs.VerdictNone,
			}); err != nil {
			params.logger.Error(err, "failed to write error result to termination-log")
			panic(err)
		}
	}
}

// runIutReleaser is the base releaser for an IUT provider. Checks input parameters and calls Release on a Provider.
func runIutReleaser(params parameters, provider Provider) error {
	ctx := context.Background()
	logger := params.logger
	if params.namespace == "" {
		return errors.New("Must set -namespace")
	}
	if params.name == "" {
		return errors.New("Must set -name")
	}
	logger.Info("Release flag set, releasing IUT")
	return provider.Release(ctx, logger, params.name, params.namespace, params.noDelete)
}

// runIutProvider is the base provider for an IUT provider. Checks input parameters and calls Provision on a Provider.
func runIutProvider(params parameters, provider Provider) error {
	ctx := context.Background()
	logger := params.logger
	if params.namespace == "" {
		return errors.New("Must set -namespace")
	}
	if params.environmentRequestName == "" {
		return errors.New("Must set -environment-request")
	}
	var request v1alpha1.EnvironmentRequest
	if err := cli.Get(ctx, types.NamespacedName{Name: params.environmentRequestName, Namespace: params.namespace}, &request); err != nil {
		return err
	}
	return provider.Provision(ctx, logger, request, params.namespace)
}

// GetIUT gets an IUT resource by name from Kubernetes.
func GetIUT(ctx context.Context, name, namespace string) (*v1alpha2.Iut, error) {
	cli, err := kubernetesClient()
	if err != nil {
		return nil, err
	}
	var iut v1alpha2.Iut
	if err := cli.Get(ctx, types.NamespacedName{Name: name, Namespace: namespace}, &iut); err != nil {
		return nil, err
	}
	return &iut, nil
}

// CreateIUT creates a new IUT resource in Kubernetes.
// TODO: Allow inputs for the IutSpec
func CreateIUT(ctx context.Context, environmentrequest *v1alpha1.EnvironmentRequest, namespace string) error {
	logger, _ := logr.FromContext(ctx)

	logger.Info("Getting Kubernetes client")
	cli, err := kubernetesClient()
	if err != nil {
		return err
	}
	labels := map[string]string{
		"etos.eiffel-community.github.io/provider":               environmentrequest.Spec.Providers.IUT.ID,
		"etos.eiffel-community.github.io/environment-request":    environmentrequest.Spec.Name,
		"etos.eiffel-community.github.io/environment-request-id": environmentrequest.Spec.ID,
		"app.kubernetes.io/name":                                 "iut-provider",
		"app.kubernetes.io/part-of":                              "etos",
	}
	if cluster := environmentrequest.Labels["etos.eiffel-community.github.io/cluster"]; cluster != "" {
		labels["etos.eiffel-community.github.io/cluster"] = cluster
	}
	if environmentrequest.Spec.Identifier != "" {
		labels["etos.eiffel-community.github.io/id"] = environmentrequest.Spec.Identifier
	}
	isController := false
	blockOwnerDeletion := true
	logger.Info("Creating IUT", "environmentrequest", environmentrequest.Spec.Name)
	return cli.Create(ctx, &v1alpha2.Iut{
		ObjectMeta: metav1.ObjectMeta{
			Labels:       labels,
			GenerateName: fmt.Sprintf("%s-iut-", strings.ToLower(environmentrequest.Spec.Name)),
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
		Spec: v1alpha2.IutSpec{
			// ID:         uuid.NewString(),
			ProviderID: environmentrequest.Spec.Providers.IUT.ID,
			// Identity:   environmentrequest.Spec.Identity,
		},
	})
}

// DeleteIUT deletes an IUT resource from Kubernetes.
func DeleteIUT(ctx context.Context, iut *v1alpha2.Iut) error {
	cli, err := kubernetesClient()
	if err != nil {
		return err
	}
	return cli.Delete(ctx, iut)
}
