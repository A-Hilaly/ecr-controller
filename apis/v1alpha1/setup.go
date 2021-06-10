package v1alpha1

import (
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var logg = logf.Log.WithName("repository-resource")

func (r *Repository) SetupWebhookWithManager(mgr ctrl.Manager) error {
	fmt.Println("setting up webhook")
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}
