// middleware of the HTTP server
package middleware

import (

	// Third-party
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

// NewMiddleware creates and returns a new instance of a middleware chain
func NewMiddleware() alice.Chain {
	return alice.New(
		NewVersion().Handler,
		NewPreflight().Handler,
		cors.New(cors.Options{}).Handler,
		NewLogger().Handler,
		NewRecovery().Handler,
	)
}