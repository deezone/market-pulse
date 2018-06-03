// server package - HTTP server functionality within application
// routes.go - all routes within server
package server

import (
	// Internal
	"github.com/deezone/forex-clock/handlers"

	// Third Party
	"github.com/gorilla/mux"
)

// SetRoute is used to set available routes used by the server
func (s *Server) SetRoutes() {
	// Create a Mux that will be used for routing
	mux := mux.NewRouter()

	// Create handlers
	hh := handlers.NewHealthHandler(s.resources.DB)

	// Set up health/readiness/version routes
	mux.HandleFunc(handlers.HealthRoute, hh.Health)
	mux.HandleFunc(handlers.ReadyRoute, hh.Ready)
	mux.HandleFunc(handlers.VersionRoute, hh.Version)

	// Set the server's routing handler to be the mux
	s.GetInstance().Handler = mux
}
