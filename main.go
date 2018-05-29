// An API to provide information about the world FOREX markets
// Started 27 May 2018
// Governed by the license that can be found in the repository LICENSE file

package main

import (
	// Standard lib
	"fmt"
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
	fmt.Println(m)
	log.Info(m)

	// Initialize configuration
	config.Init()

	// Create new server
	s := server.NewServer()

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
		log.Info("Server is shutting down")
	}
}
