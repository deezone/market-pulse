// The server package contains all code relating to the HTTP server used within the application
package server

import (
	// Standard lib
	"fmt"
	"net/http"
	"time"
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

	return &Server{
		instance: &http.Server{
			Addr:         fmt.Sprintf(":%d", c.Server.Port),
			ReadTimeout:  time.Duration(c.Server.Timeouts.Read) * time.Second,
			WriteTimeout: time.Duration(c.Server.Timeouts.Write) * time.Second,
		},
		resources: &Resources{
			DB:  db.NewSCDB(),
		},
		running: false,
	}
}
