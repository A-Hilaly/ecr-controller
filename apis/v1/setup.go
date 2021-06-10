package v1

import (
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var logg = logf.Log.WithName("repository-resource")

func (r *Repository) SetupWebhookWithManager(mgr ctrl.Manager) error {
	fmt.Println("setting up webhook v1")
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}
