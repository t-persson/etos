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
	"context"
	"errors"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	etosv1alpha2 "github.com/eiffel-community/etos/api/v1alpha2"
	"github.com/eiffel-community/etos/internal/controller/jobs"
	"github.com/eiffel-community/etos/internal/controller/status"
)

// ExecutionSpaceReconciler reconciles a ExecutionSpace object
type ExecutionSpaceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=etos.eiffel-community.github.io,resources=executionspaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=etos.eiffel-community.github.io,resources=executionspaces/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=etos.eiffel-community.github.io,resources=executionspaces/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.22.1/pkg/reconcile
func (r *ExecutionSpaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := logf.FromContext(ctx)

	executionSpace := &etosv1alpha2.ExecutionSpace{}
	err := r.Get(ctx, req.NamespacedName, executionSpace)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if hasOwner(executionSpace.OwnerReferences, "Environment") {
		if controllerutil.ContainsFinalizer(executionSpace, providerFinalizer) {
			meta.SetStatusCondition(&executionSpace.Status.Conditions,
				metav1.Condition{
					Status:  metav1.ConditionTrue,
					Type:    status.StatusActive,
					Reason:  status.ReasonActive,
					Message: "In use",
				})
			controllerutil.RemoveFinalizer(executionSpace, providerFinalizer)
			if err := r.Update(ctx, executionSpace); err != nil {
				if apierrors.IsConflict(err) {
					return ctrl.Result{Requeue: true}, nil
				}
				return ctrl.Result{}, err
			}
		}
		logger.Info("ExecutionSpace is being managed by Environment", "executionSpace", executionSpace.Name)
		// We no longer own this ExecutionSpace. Let the Environment controller manage it.
		return ctrl.Result{}, nil
	}
	// If the ExecutionSpace is considered 'Completed', it has been released. Check that the object is
	// being deleted and contains the finalizer and remove the finalizer.
	if executionSpace.Status.CompletionTime != nil {
		if !executionSpace.ObjectMeta.DeletionTimestamp.IsZero() {
			if controllerutil.ContainsFinalizer(executionSpace, providerFinalizer) {
				controllerutil.RemoveFinalizer(executionSpace, providerFinalizer)
				if err := r.Update(ctx, executionSpace); err != nil {
					if apierrors.IsConflict(err) {
						return ctrl.Result{Requeue: true}, nil
					}
					return ctrl.Result{}, err
				}
			}
		}
		return ctrl.Result{}, nil
	}
	if err := r.reconcile(ctx, executionSpace); err != nil {
		if apierrors.IsConflict(err) {
			return ctrl.Result{Requeue: true}, nil
		}
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// reconcile an ExecutionSpace resource to its desired state.
func (r *ExecutionSpaceReconciler) reconcile(ctx context.Context, executionSpace *etosv1alpha2.ExecutionSpace) error {
	logger := logf.FromContext(ctx)

	// Set initial statuses if not set.
	if meta.FindStatusCondition(executionSpace.Status.Conditions, status.StatusActive) == nil {
		meta.SetStatusCondition(&executionSpace.Status.Conditions,
			metav1.Condition{
				Status:  metav1.ConditionFalse,
				Type:    status.StatusActive,
				Reason:  status.ReasonPending,
				Message: "Waiting for environment",
			})
		return r.Status().Update(ctx, executionSpace)
	}
	if executionSpace.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(executionSpace, providerFinalizer) {
			controllerutil.AddFinalizer(executionSpace, providerFinalizer)
			logger.Info("ExecutionSpace is being managed by ExecutionSpace controller", "executionSpace", executionSpace.Name)
			return r.Update(ctx, executionSpace)
		}
	}

	if !executionSpace.ObjectMeta.DeletionTimestamp.IsZero() {
		return r.reconcileExecutionSpaceReleaser(ctx, executionSpace)
	}
	return nil
}

// reconcileExecutionSpaceReleaser gets the status of a release job, creating a new release job if necessary.
func (r *ExecutionSpaceReconciler) reconcileExecutionSpaceReleaser(ctx context.Context, executionSpace *etosv1alpha2.ExecutionSpace) error {
	conditions := &executionSpace.Status.Conditions
	jobManager := jobs.NewJob(r.Client, ExecutionSpaceOwnerKey, executionSpace.GetName(), executionSpace.GetNamespace())
	jobStatus, err := jobManager.Status(ctx)
	if err != nil {
		return err
	}
	switch jobStatus {
	case jobs.StatusFailed:
		result := jobManager.Result(ctx)
		if meta.SetStatusCondition(conditions,
			metav1.Condition{
				Type:    status.StatusActive,
				Status:  metav1.ConditionFalse,
				Reason:  status.ReasonFailed,
				Message: result.Description,
			}) {
			return r.Status().Update(ctx, executionSpace)
		}
	case jobs.StatusSuccessful:
		result := jobManager.Result(ctx)
		var condition metav1.Condition
		if result.Conclusion == jobs.ConclusionFailed {
			condition = metav1.Condition{
				Type:    status.StatusActive,
				Status:  metav1.ConditionFalse,
				Reason:  status.ReasonFailed,
				Message: result.Description,
			}
		} else {
			condition = metav1.Condition{
				Type:    status.StatusActive,
				Status:  metav1.ConditionFalse,
				Reason:  status.ReasonCompleted,
				Message: result.Description,
			}
		}
		executionSpaceCondition := meta.FindStatusCondition(*conditions, status.StatusActive)
		executionSpace.Status.CompletionTime = &executionSpaceCondition.LastTransitionTime
		if meta.SetStatusCondition(conditions, condition) {
			return errors.Join(r.Status().Update(ctx, executionSpace), jobManager.Delete(ctx))
		}
	case jobs.StatusActive:
		if meta.SetStatusCondition(conditions,
			metav1.Condition{
				Type:    status.StatusActive,
				Status:  metav1.ConditionFalse,
				Reason:  status.ReasonPending,
				Message: "Releasing ExecutionSpace",
			}) {
			return r.Status().Update(ctx, executionSpace)
		}
	default:
		// Since this is a release job, we don't want to release if we are not deleting.
		if executionSpace.GetDeletionTimestamp().IsZero() {
			return nil
		}
		if err := jobManager.Create(ctx, executionSpace, r.releaseJob); err != nil {
			if meta.SetStatusCondition(conditions,
				metav1.Condition{
					Type:    status.StatusActive,
					Status:  metav1.ConditionFalse,
					Reason:  status.ReasonFailed,
					Message: err.Error(),
				}) {
				return r.Status().Update(ctx, executionSpace)
			}
			return err
		}
		if meta.SetStatusCondition(conditions, metav1.Condition{
			Status:  metav1.ConditionFalse,
			Type:    status.StatusActive,
			Reason:  status.ReasonPending,
			Message: "Releasing ExecutionSpace",
		}) {
			return r.Status().Update(ctx, executionSpace)
		}
	}
	return nil
}

// releaseJob is the job definition for an execution space releaser.
func (r ExecutionSpaceReconciler) releaseJob(ctx context.Context, obj client.Object) (*batchv1.Job, error) {
	ttl := int32(300)
	grace := int64(30)
	backoff := int32(0)

	executionSpace, ok := obj.(*etosv1alpha2.ExecutionSpace)
	if !ok {
		return nil, errors.New("object received from job manager is not an ExecutionSpace")
	}

	provider, err := getProvider(ctx, r, executionSpace.Spec.ProviderID, executionSpace.GetNamespace())
	if err != nil {
		return nil, err
	}

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app.kubernetes.io/name":    "execution-space-releaser",
				"app.kubernetes.io/part-of": "etos",
			},
			Annotations: make(map[string]string),
			Name:        executionSpace.Name,
			Namespace:   executionSpace.Namespace,
		},
		Spec: batchv1.JobSpec{
			TTLSecondsAfterFinished: &ttl,
			BackoffLimit:            &backoff,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: executionSpace.Name,
					Labels: map[string]string{
						"app.kubernetes.io/name":    "execution-space-releaser",
						"app.kubernetes.io/part-of": "etos",
					},
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: &grace,
					ServiceAccountName:            "provider", // TODO: Wrong service account
					RestartPolicy:                 "Never",
					Containers: []corev1.Container{
						{
							Name:            executionSpace.Name,
							Image:           image(provider),
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
							Args: []string{
								"-release",
								"-nodelete", // The resource is already being deleted, no need to delete again.
								fmt.Sprintf("-namespace=%s", executionSpace.GetNamespace()),
								fmt.Sprintf("-execution-space=%s", executionSpace.GetName()),
							},
						},
					},
				},
			},
		},
	}
	return jobSpec, ctrl.SetControllerReference(executionSpace, jobSpec, r.Scheme)
}

// registerOwnerIndexForJob will set an index of the jobs that an execution space controller owns.
func (r *ExecutionSpaceReconciler) registerOwnerIndexForJob(mgr ctrl.Manager) error {
	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &batchv1.Job{}, ExecutionSpaceOwnerKey, func(rawObj client.Object) []string {
		job := rawObj.(*batchv1.Job)
		owner := metav1.GetControllerOf(job)
		if owner == nil {
			return nil
		}
		if owner.APIVersion != APIv2GroupVersionString || owner.Kind != "ExecutionSpace" {
			return nil
		}

		return []string{owner.Name}
	}); err != nil {
		return err
	}
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ExecutionSpaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// Register indexes for faster lookups
	if err := r.registerOwnerIndexForJob(mgr); err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&etosv1alpha2.ExecutionSpace{}).
		Named("executionspace").
		Owns(&batchv1.Job{}). // Release job
		Complete(r)
}
