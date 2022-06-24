package namespace

import (
	"github.com/douglasmakey/admissioncontroller"
	log "k8s.io/klog/v2"

	"k8s.io/api/admission/v1beta1"
)

func validateDelete() admissioncontroller.AdmitFunc {
	return func(r *v1beta1.AdmissionRequest) (*admissioncontroller.Result, error) {
		if r.Kind.Kind == "Namespace" {
			log.Info("You cannot delete namespace: ", r.Name)
			return &admissioncontroller.Result{Allowed: false}, nil
		} else {
			return &admissioncontroller.Result{Allowed: true}, nil
		}
	}
}
