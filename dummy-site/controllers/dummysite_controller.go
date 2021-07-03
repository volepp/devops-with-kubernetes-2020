/*
Copyright 2021.

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
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	stabledwkv1 "stable.dwk/api/v1"
)

// DummySiteReconciler reconciles a DummySite object
type DummySiteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func randomBase16String(l int) string {
	buff := make([]byte, int(math.Round(float64(l)/2)))
	rand.Read(buff)
	str := hex.EncodeToString(buff)
	return str[:l] // strip 1 extra character we get from odd length results
}

//+kubebuilder:rbac:groups=stable.dwk.stable.dwk,resources=dummysites,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=stable.dwk.stable.dwk,resources=dummysites/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=stable.dwk.stable.dwk,resources=dummysites/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DummySite object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *DummySiteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var dummySite stabledwkv1.DummySite
	if err := r.Get(ctx, req.NamespacedName, &dummySite); err != nil {
		fmt.Println(err, "unable to fetch DummySite")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	dep, err := getDep(&dummySite, r)
	if err != nil {
		fmt.Printf("Error constructing deployment: %s\n", err)
		return ctrl.Result{}, err
	}

	if err := r.Create(ctx, dep); err != nil {
		fmt.Println("Unable to create Deployment for DummySite:", err)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func getDep(dummySite *stabledwkv1.DummySite, r *DummySiteReconciler) (*appsv1.Deployment, error) {
	replicas := new(int32)
	*replicas = 1
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dummySite.Name + "-dep-" + randomBase16String(6),
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "dummy-site",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "dummy-site",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "dummy-site",
							Image: "volepp/dummy-site",
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 3000,
								},
							},
							Env: []apiv1.EnvVar{
								{
									Name:  "WEBSITE_URL",
									Value: dummySite.Spec.WebsiteUrl,
								},
							},
						},
					},
				},
			},
		},
	}

	if err := ctrl.SetControllerReference(dummySite, dep, r.Scheme); err != nil {
		return nil, err
	}

	return dep, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DummySiteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&stabledwkv1.DummySite{}).
		Complete(r)
}
