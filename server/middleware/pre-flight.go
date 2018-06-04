// middleware of the HTTP server
// The pre-flight middleware adds common response headers to requests
package middleware

import (
	// Standard Lib
	"net/http"

	// Internal
	"github.com/deezone/forex-clock/config"
)

type (
	// Struct representing pre-flight middleware
	Preflight struct{}
)

// NewPreflight creates and returns a new instance of pre-flight middleware
func NewPreflight() Preflight { return Preflight{} }

// Handler handles the processing of the request
// The pre-flight middleware handler adds common response headers to requests
func (m Preflight) Handler(next http.Handler) http.Handler {
	// Middleware handler function
	fn := func(w http.ResponseWriter, req *http.Request) {
		// Add application name header
		w.Header().Add("X-Powered-By", config.GetInstance().Name)

		// TO-DO: Options check, exit early

		// Pass the request through
		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}
