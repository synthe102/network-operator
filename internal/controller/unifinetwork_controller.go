/*
Copyright 2024.

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
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/paultyng/go-unifi/unifi"
	networkv1alpha1 "github.com/synthe102/network-operator/api/v1alpha1"
)

// UnifiNetworkReconciler reconciles a UnifiNetwork object
type UnifiNetworkReconciler struct {
	client.Client
	Scheme      *runtime.Scheme
	UnifiClient *unifi.Client
}

//+kubebuilder:rbac:groups=network.suslian.engineer,resources=unifinetworks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=network.suslian.engineer,resources=unifinetworks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=network.suslian.engineer,resources=unifinetworks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the UnifiNetwork object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *UnifiNetworkReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info(fmt.Sprintf("Started reconciliation for network %s", req.Name))

	var un networkv1alpha1.UnifiNetwork
	if err := r.Client.Get(ctx, req.NamespacedName, &un); err != nil {
		if client.IgnoreNotFound(err) != nil {
			log.Error(err, "failed to get unifi network CR instance")
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	if un.DeletionTimestamp != nil {

	}

	networks, err := r.UnifiClient.ListNetwork(ctx, un.Spec.Site)
	if err != nil {
		log.Error(err, "failed to list networks")
		return ctrl.Result{}, err
	}

	shouldCreate := true
	for _, n := range networks {
		if n.Name == un.Spec.Name {
			if n.VLAN == un.Spec.VlanID {
				log.Info("network alreay exists")
				// No network is found with the same name and VLAN ID.
				shouldCreate = false
			}
		}
	}
	if shouldCreate {
		network := &unifi.Network{
			Name:        un.Spec.Name,
			Purpose:     "vlan-only",
			VLAN:        un.Spec.VlanID,
			VLANEnabled: true,
			Enabled:     true,
		}
		network, err = r.UnifiClient.CreateNetwork(ctx, un.Spec.Site, network)
		if err != nil {
			log.Error(err, "failed to create unifi network")
			return ctrl.Result{}, err
		}
		un.Status.CIDR = network.IPSubnet
		un.Status.ID = network.ID
		if err := r.Client.Status().Update(ctx, &un); err != nil {
			log.Error(err, "failed to update unifi network CR status")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{
		RequeueAfter: 15 * time.Second,
	}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *UnifiNetworkReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&networkv1alpha1.UnifiNetwork{}).
		Complete(r)
}
