package namespace

import (
	"github.com/douglasmakey/admissioncontroller"
	"log"

	"k8s.io/api/admission/v1beta1"
)

func validateDelete() admissioncontroller.AdmitFunc {
	return func(r *v1beta1.AdmissionRequest) (*admissioncontroller.Result, error) {
		if r.Kind.Kind == "Namespace" {
			log.Println("begin......", r.Object.Raw)
			namespace, err := parseNamespace(r.Object.Raw)
			log.Println("end.....", namespace)
			if err != nil {
				return &admissioncontroller.Result{Msg: err.Error()}, nil
			} else {
				return &admissioncontroller.Result{Msg: "You cannot delete namespace: " + namespace.Name, Allowed: false}, nil
			}
		} else {
			return &admissioncontroller.Result{Allowed: true}, nil
		}
	}
}
