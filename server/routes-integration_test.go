// Tests the routes.go file
package server

import (
	// Standard lib
	"fmt"
	"time"

	// Internal
	"github.com/deezone/forex-clock/config"

	// Third-party
	goutils "github.com/marksost/go-utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("routes.go", func() {
	var (
		// Mock server to test
		s *Server
		// The address of the server to use during integration tests
		serverAddress string
	)

	BeforeEach(func() {
		// Get empty port
		port, err := goutils.GetEmptyPort()
		if err != nil {
			panic("Error getting an empty port. Testing cannot continue. Error was: " + err.Error())
		}

		// Set server port and server address
		config.GetInstance().Server.Port = port
		serverAddress = "http://localhost:" + goutils.Int2String(port)

		// Create server instance
		s = NewServer()

		// Start server
		err = s.Start()
		if err != nil {
			panic("Error starting server. Testing cannot continue. Error was: " + err.Error())
		}

		// Sleep so the server can start
		time.Sleep(500 * time.Millisecond)
	})

	AfterEach(func() {
		// Stop Server
		s.Stop()
	})

	Describe("Route integration tests", func() {
		var (
			// Slice of route integration tests to run
			input []*RoutesTestData
		)

		BeforeEach(func() {
			// Set input
			input = []*RoutesTestData{
				// Non-valid route
				&RoutesTestData{Method: "GET", Route: "/invalid", ResponseCode: 404},

				/* Health/Ready Routes */

				// Health, Ready, and Version with invalid method
				&RoutesTestData{Method: "HEAD", Route: "/health", ResponseCode: 405},
				&RoutesTestData{Method: "HEAD", Route: "/ready", ResponseCode: 405},
				&RoutesTestData{Method: "HEAD", Route: "/version", ResponseCode: 405},
				// Health with valid method
				&RoutesTestData{Method: "GET", Route: "/health", ResponseCode: 200},
				// Version with valid method
				&RoutesTestData{Method: "GET", Route: "/version", ResponseCode: 200},
			}
		})

		It("Resolves requests properly", func() {
			// Loop through test data
			for _, conf := range input {
				// Make request config
				rc := goutils.NewRequestConfig()

				// Set values
				rc.Method = conf.Method
				rc.URL = serverAddress + conf.Route

				// Make request
				code, err := goutils.GetStatusCodeForRequest(rc)

				// Verify response
				Expect(err).To(Not(HaveOccurred()))
				Expect(code).To(Equal(conf.ResponseCode), fmt.Sprintf("%s method, %s route", conf.Method, conf.Route))
			}
		})
	})
})
