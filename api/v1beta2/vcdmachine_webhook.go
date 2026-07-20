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

package v1beta2

import (
	"context"

	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var vcdmachinelog = logf.Log.WithName("vcdmachine-resource")

func (r *VCDMachine) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr, r).
		WithValidator(r).
		WithDefaulter(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-infrastructure-cluster-x-k8s-io-v1beta2-vcdmachine,mutating=true,failurePolicy=fail,sideEffects=None,groups=infrastructure.cluster.x-k8s.io,resources=vcdmachines,verbs=create;update,versions=v1beta2,name=mutation.vcdmachine.infrastructure.cluster.x-k8s.io,admissionReviewVersions=v1

var _ admission.Defaulter[*VCDMachine] = &VCDMachine{}

// Default implements admission.Defaulter so a webhook will be registered for the type
func (r *VCDMachine) Default(ctx context.Context, obj *VCDMachine) error {
	vcdmachinelog.Info("default", "name", obj.Name)

	// TODO(user): fill in your defaulting logic.
	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-infrastructure-cluster-x-k8s-io-v1beta2-vcdmachine,mutating=false,failurePolicy=fail,sideEffects=None,groups=infrastructure.cluster.x-k8s.io,resources=vcdmachines,verbs=create;update,versions=v1beta2,name=validation.vcdmachine.infrastructure.cluster.x-k8s.io,admissionReviewVersions=v1

var _ admission.Validator[*VCDMachine] = &VCDMachine{}

// ValidateCreate implements admission.Validator so a webhook will be registered for the type
func (r *VCDMachine) ValidateCreate(ctx context.Context, obj *VCDMachine) (admission.Warnings, error) {
	vcdmachinelog.Info("validate create", "name", obj.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements admission.Validator so a webhook will be registered for the type
func (r *VCDMachine) ValidateUpdate(ctx context.Context, oldObj, newObj *VCDMachine) (admission.Warnings, error) {
	vcdmachinelog.Info("validate update", "name", newObj.Name)

	// the relation between CR->VM_NAME have to be consistent therefore VmNamingTemplate is immutable
	if newObj.Spec.VmNamingTemplate != oldObj.Spec.VmNamingTemplate {
		return nil, field.Forbidden(field.NewPath("spec.vmNamingTemplate"), "cannot be modified")
	}

	return nil, nil
}

// ValidateDelete implements admission.Validator so a webhook will be registered for the type
func (r *VCDMachine) ValidateDelete(ctx context.Context, obj *VCDMachine) (admission.Warnings, error) {
	vcdmachinelog.Info("validate delete", "name", obj.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
