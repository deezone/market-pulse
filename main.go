// An API to provide information about the world FOREX markets
// Started 27 May 2018
// Governed by the license that can be found in the repository LICENSE file

package main

import (
	"os"
	"os/signal"

	// Internal
	"github.com/deezone/forex-clock/config"
	"github.com/deezone/forex-clock/server"

	// Third-party
	log "github.com/sirupsen/logrus"
)

// Main function
// Starting point for application - `go run`
func main() {
	m := "Starting forex-clock application..."
	log.Info(m)

	// Initialize configuration
	config.Init()

	m = "Configuration loaded..."
	log.Info(m)

	// Create new server
	s := server.NewServer()

	m = "Creating new server..."
	log.Info(m)

	// Start server
	if err := s.Start(); err != nil {
		panic("Error starting application. Error was: " + err.Error())
	}

	// Listen for and exit the application on SIGKILL or SIGINT
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt, os.Kill)

	select {
	case <-stop:
		// Attempt to stop the server
		s.Stop()

		// Log shut down
		m = "Server is shutting down..."
		log.Info(m)
	}
}
