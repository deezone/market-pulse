// middleware of the HTTP server
// The recovery middleware attempts to recover from panics that occur within the request execution chain
package middleware

import (
	// Standard Lib
	"net/http"

	// Internal
	"github.com/deezone/forex-clock/helpers"

	// Third Party
	log "github.com/sirupsen/logrus"
)

type (
	// Struct representing recovery middleware
	Recovery struct{}
)

// NewRecovery creates and returns a new instance of recovery middleware
func NewRecovery() Recovery { return Recovery{} }

// Handler handles the processing of the request
// The recovery middleware handler attempts to recover from any panics
// that occur within the request execution chain
func (m Recovery) Handler(next http.Handler) http.Handler {
	// Middleware handler function
	fn := func(w http.ResponseWriter, req *http.Request) {
		// Defer a panic check until the end of the request,
		// and handle it as needed
		defer func() {
			if err := recover(); err != nil {
				log.WithField("error", err).Warn("Recovering from panic")
				helpers.InternalError(w, req)
			}
		}()

		// Pass the request through
		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}
