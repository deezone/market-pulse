// The server package contains all code relating to the HTTP server used within the application
package server

import (
	// Standard lib
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	// Internal
	"github.com/deezone/forex-clock/config"
	"github.com/deezone/forex-clock/db"
	"github.com/deezone/forex-clock/server/middleware"

	// Third-party
	"github.com/labstack/gommon/log"
)

type (
	// Struct representing the various internal resources request handlers may need to access
	Resources struct {
		DB  db.DB   // The database instance to use
	}
	// Struct representing the actual http.Server and helper data
	Server struct {
		instance  *http.Server // HTTP server that will be serving requests
		resources *Resources   // Internal resources request handlers may need to access
		running   bool         // Boolean indicating if the server is running or not
	}
)

// NewServer creates and returns a new instance of a server
func NewServer() *Server {
	c := config.GetInstance()

	return &Server {
		instance: &http.Server{
			Addr:         fmt.Sprintf(":%d", c.Server.Port),
			ReadTimeout:  time.Duration(c.Server.Timeouts.Read) * time.Second,
			WriteTimeout: time.Duration(c.Server.Timeouts.Write) * time.Second,
		},
		resources: &Resources{
			DB:  db.NewFCDB(),
		},
		running: false,
	}
}

// Start is used to start a non-running server
func (s *Server) Start() error {
	// Check if the server is running
	if s.IsRunning() {
		return errors.New("Attempted to start a server that is already running at address: " + s.instance.Addr)
	}

	// Set routes and middleware
	s.SetRoutes()
	s.GetInstance().Handler = middleware.NewMiddleware().Then(s.GetInstance().Handler)

	m := "Listening for requests..."
	log.Info(m)

	go s.GetInstance().ListenAndServe()

	s.running = true

	return nil
}

// Stop is used to stop a running server
func (s *Server) Stop() error {
	// Check if the server is running
	if !s.IsRunning() {
		return errors.New("Attempted to stop a non-running server")
	}

	// Shut down server gracefully, but wait no longer than a configured amount of seconds before halting
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.GetInstance().Server.Timeouts.ShutDown)*time.Second)
	if err := s.instance.Shutdown(ctx); err != nil {
		return err
	}

	s.running = false

	return nil
}

// GetInstance returns the internal http.Server instance of the server
func (s *Server) GetInstance() *http.Server { return s.instance }

// IsRunning returns a boolean indicating if the server is running or not
func (s *Server) IsRunning() bool { return s.running }
