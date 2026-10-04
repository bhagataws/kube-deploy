/*
Copyright 2026 bhagataws.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	corev1 "github.com/bhagataws/kube-deploy.git/api/v1"
)

// WebappReconciler reconciles a Webapp object
type WebappReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=core.cafeco.in,resources=webapps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core.cafeco.in,resources=webapps/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core.cafeco.in,resources=webapps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Webapp object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.23.3/pkg/reconcile
func (r *WebappReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logs := logf.FromContext(ctx)
	webapp := &corev1.Webapp{}
	if err := r.Get(ctx, req.NamespacedName, webapp); err != nil {
		if apierrors.IsNotFound(err) {
			logs.Info("Webapp does not exist", "name", req.Name, "namespace", req.Namespace)
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	logs.Info("Webapp exists", "name", webapp.Name, "namespace", webapp.Namespace, "replicas", webapp.Spec.Replicas, "image", webapp.Spec.Image, "context", webapp.Spec.Context)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *WebappReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Webapp{}).
		Named("webapp").
		Complete(r)
}
