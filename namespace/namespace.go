package namespace

import (
	"github.com/douglasmakey/admissioncontroller"
)

// NewValidationHook delete namespace validation hook
func NewValidationHook() admissioncontroller.Hook {
	return admissioncontroller.Hook{
		Delete: validateDelete(),
	}
}
