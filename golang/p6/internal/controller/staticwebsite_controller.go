package controller

import (
	"context"

	webv1 "github.com/hassan/webforge-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// StaticWebsiteReconciler reconciles a StaticWebsite object.
type StaticWebsiteReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=web.easy.com,resources=staticwebsites,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=web.easy.com,resources=staticwebsites/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=web.easy.com,resources=staticwebsites/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;patch;delete

func (r *StaticWebsiteReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var site webv1.StaticWebsite

	if err := r.Get(ctx, req.NamespacedName, &site); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if err := r.reconcileDeployment(ctx, &site); err != nil {
		return ctrl.Result{}, err
	}

	if err := r.reconcileService(ctx, &site); err != nil {
		return ctrl.Result{}, err
	}

	if err := r.updateStatus(ctx, &site); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *StaticWebsiteReconciler) reconcileDeployment(ctx context.Context, site *webv1.StaticWebsite) error {
	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      site.Name,
			Namespace: site.Namespace,
		},
	}

	_, err := controllerutil.CreateOrUpdate(ctx, r.Client, deploy, func() error {
		if err := controllerutil.SetControllerReference(site, deploy, r.Scheme); err != nil {
			return err
		}

		replicas := int32(1)
		if site.Spec.Replicas != nil {
			replicas = *site.Spec.Replicas
		}

		labels := map[string]string{
			"app": site.Name,
		}

		deploy.Labels = labels

		deploy.Spec.Replicas = &replicas
		deploy.Spec.Selector = &metav1.LabelSelector{
			MatchLabels: labels,
		}

		deploy.Spec.Template = corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: labels,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "web",
						Image: site.Spec.Image,
						Ports: []corev1.ContainerPort{
							{
								ContainerPort: site.Spec.Port,
							},
						},
					},
				},
			},
		}

		return nil
	})

	return err
}

func (r *StaticWebsiteReconciler) reconcileService(ctx context.Context, site *webv1.StaticWebsite) error {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      site.Name,
			Namespace: site.Namespace,
		},
	}

	_, err := controllerutil.CreateOrUpdate(ctx, r.Client, svc, func() error {
		if err := controllerutil.SetControllerReference(site, svc, r.Scheme); err != nil {
			return err
		}

		labels := map[string]string{
			"app": site.Name,
		}

		svc.Labels = labels
		svc.Spec.Selector = labels
		svc.Spec.Ports = []corev1.ServicePort{
			{
				Name:       "http",
				Port:       site.Spec.Port,
				TargetPort: intstr.FromInt(int(site.Spec.Port)),
				Protocol:   corev1.ProtocolTCP,
			},
		}

		return nil
	})

	return err
}

func (r *StaticWebsiteReconciler) updateStatus(ctx context.Context, site *webv1.StaticWebsite) error {
	if site.Status.Phase == "Ready" &&
		site.Status.DeploymentName == site.Name &&
		site.Status.ServiceName == site.Name {
		return nil
	}

	patch := client.MergeFrom(site.DeepCopy())

	site.Status.Phase = "Ready"
	site.Status.DeploymentName = site.Name
	site.Status.ServiceName = site.Name

	return r.Status().Patch(ctx, site, patch)
}

// SetupWithManager sets up the controller with the Manager.
func (r *StaticWebsiteReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webv1.StaticWebsite{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Complete(r)
}
