package http

import (
	"fmt"
	"github.com/joker-bai/validate-namespace/namespace"
	"net/http"
)

// NewServer creates and return a http.Server
func NewServer(port string) *http.Server {
	// Instances hooks
	nsValidation := namespace.NewValidationHook()

	// Routers
	ah := newAdmissionHandler()
	mux := http.NewServeMux()
	mux.Handle("/healthz", healthz())
	mux.Handle("/validate/delete-namespace", ah.Serve(nsValidation))

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
}
