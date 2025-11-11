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
	"flag"

	"github.com/eiffel-community/etos/internal/controller/jobs"
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// RunLogAreaProvider is the base runner for an LogArea provider. Checks input parameters and calls either Release or Provision on a Provider.
//
// This function panics on errors, propagating errors back to the controller that executed it.
func RunLogAreaProvider(provider Provider) {
	opts := zap.Options{
		Development: true,
	}
	params := parameters{}
	flag.BoolVar(&params.releaseEnvironment, "release", false, "Release LogArea instead of creating")
	flag.BoolVar(&params.noDelete, "nodelete", false, "Don't delete the LogArea resource")
	flag.StringVar(&params.environmentRequestName, "environment-request", "", "The environment request to request LogAreas for.")
	flag.StringVar(&params.name, "log-area", "", "The LogArea to release.")
	flag.StringVar(&params.providerName, "provider", "", "The provider used to release.")
	flag.StringVar(&params.namespace, "namespace", "", "The namespace of the environment request.")
	opts.BindFlags(flag.CommandLine)
	flag.Parse()
	logger := zap.New(zap.UseFlagOptions(&opts))

	if err := runLogAreaProvider(params, logger, provider); err != nil {
		if err := WriteResult(logger,
			jobs.Result{
				Conclusion:  jobs.ConclusionFailed,
				Description: err.Error(),
				Verdict:     jobs.VerdictNone,
			}); err != nil {
			logger.Error(err, "failed to write error result to termination-log")
		}
		panic(err)
	}

	if params.releaseEnvironment {
		if err := WriteResult(logger,
			jobs.Result{
				Conclusion:  jobs.ConclusionSuccessful,
				Description: "Successfully released LogArea",
				Verdict:     jobs.VerdictNone,
			}); err != nil {
			logger.Error(err, "failed to write error result to termination-log")
			panic(err)
		}
	} else {
		if err := WriteResult(logger,
			jobs.Result{
				Conclusion:  jobs.ConclusionSuccessful,
				Description: "Successfully provisioned LogAreas",
				Verdict:     jobs.VerdictNone,
			}); err != nil {
			logger.Error(err, "failed to write error result to termination-log")
			panic(err)
		}
	}
}

// runLogAreaProvider is the base runner for an LogArea provider. Checks input parameters and calls either Release or Provision on a Provider.
func runLogAreaProvider(params parameters, logger logr.Logger, provider Provider) error {
	return nil
}
