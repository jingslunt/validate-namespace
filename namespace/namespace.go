package namespace

import (
	"encoding/json"

	"github.com/douglasmakey/admissioncontroller"

	v1 "k8s.io/api/core/v1"
)

// NewValidationHook delete namespace validation hook
func NewValidationHook() admissioncontroller.Hook {
	return admissioncontroller.Hook{
		Delete: validateDelete(),
	}
}

func parseNamespace(object []byte) (*v1.Namespace, error) {
	var ns v1.Namespace
	if err := json.Unmarshal(object, &ns); err != nil {
		return nil, err
	}

	return &ns, nil
}
