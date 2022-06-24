package namespace

import (
	"fmt"
	"github.com/douglasmakey/admissioncontroller"

	"k8s.io/api/admission/v1beta1"
)

func validateDelete() admissioncontroller.AdmitFunc {
	return func(r *v1beta1.AdmissionRequest) (*admissioncontroller.Result, error) {
		ns, err := parseNamespace(r.Object.Raw)
		fmt.Println(ns)
		if err != nil {
			return &admissioncontroller.Result{Msg: err.Error()}, nil
		} else {
			return &admissioncontroller.Result{Msg: "You cannot delete namespace" + ns.Name}, nil
		}
	}
}
