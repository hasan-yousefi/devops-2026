package controller

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	corev1 "k8s.io/api/core/v1"
	janitorv1alpha1 "mydomain.io/namespace-janitor/api/v1alpha1"
)

type TTLNamespaceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=janitor.mydomain.io,resources=ttlnamespaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=janitor.mydomain.io,resources=ttlnamespaces/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch;delete

func (r *TTLNamespaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// 1. Fetch the TTLNamespace resource
	var ttlNS janitorv1alpha1.TTLNamespace
	if err := r.Get(ctx, req.NamespacedName, &ttlNS); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// 2. Calculate expiration if not already set
	if ttlNS.Status.ExpiresAt == nil {
		duration, err := time.ParseDuration(ttlNS.Spec.TTL)
		if err != nil {
			logger.Error(err, "Failed to parse TTL duration")
			return ctrl.Result{}, nil // Don't retry invalid formats
		}

		expiration := metav1.NewTime(ttlNS.CreationTimestamp.Add(duration))
		ttlNS.Status.ExpiresAt = &expiration
		ttlNS.Status.Phase = "Active"

		if err := r.Status().Update(ctx, &ttlNS); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true}, nil
	}

	// 3. Check if TTL has expired
	now := time.Now()
	if now.After(ttlNS.Status.ExpiresAt.Time) {
		logger.Info("TTL expired. Deleting target namespace", "namespace", ttlNS.Spec.TargetNamespace)

		// Fetch the target namespace to delete it
		var targetNS corev1.Namespace
		err := r.Get(ctx, client.ObjectKey{Name: ttlNS.Spec.TargetNamespace}, &targetNS)
		if err != nil {
			if errors.IsNotFound(err) {
				// Already deleted
				return ctrl.Result{}, nil
			}
			return ctrl.Result{}, err
		}

		// Delete the target namespace
		if err := r.Delete(ctx, &targetNS); err != nil {
			return ctrl.Result{}, err
		}

		// Clean up the CRD itself
		return ctrl.Result{}, r.Delete(ctx, &ttlNS)
	}

	// 4. If not expired, requeue the reconciliation exactly when it *should* expire
	remainingTime := ttlNS.Status.ExpiresAt.Sub(now)
	return ctrl.Result{RequeueAfter: remainingTime}, nil
}

func (r *TTLNamespaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&janitorv1alpha1.TTLNamespace{}).
		Complete(r)
}
