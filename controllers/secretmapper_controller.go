/*
Copyright 2023 Eirik Opheim.

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

package controllers

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/EirikOpheim/operator-testing.git/api/v1alpha1"
)

// SecretmapperReconciler reconciles a Secretmapper object
type SecretmapperReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=theplatformteam.com,resources=secretmappers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=theplatformteam.com,resources=secretmappers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=theplatformteam.com,resources=secretmappers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Secretmapper object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *SecretmapperReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.WithValues("secretmapper", req.NamespacedName)

	// Fetch the SecretMapper custom resource
	var secretMapper v1alpha1.SecretMapper
	if err := r.Get(ctx, req.NamespacedName, &secretMapper); err != nil {
		log.Error(err, "Unable to fetch SecretMapper")
		return ctrl.Result{}, err
	}

	// Extract namespace and secret name from the custom resource spec
	namespace := secretMapper.Spec.Source.Namespace
	secretName := secretMapper.Spec.Source.Name

	// Get the Kubernetes configuration from the manager
	config := r.Mgr.GetConfig()
	if config == nil {
		log.Error(nil, "Kubernetes configuration not found")
		return ctrl.Result{}, nil
	}

	// Create a Kubernetes clientset using the obtained configuration
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error(err, "Unable to create Kubernetes client")
		return ctrl.Result{}, err
	}

	// Get the Secret
	secret, err := clientset.CoreV1().Secrets(namespace).Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		log.Error(err, "Unable to fetch Secret")
		return ctrl.Result{}, err
	}

	// Use the secret data as needed
	log.Info("Secret data:", "data", secret.Data)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SecretmapperReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Secretmapper{}).
		Complete(r)
}
