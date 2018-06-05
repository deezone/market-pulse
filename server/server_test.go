// Tests the server.go file
package server

import (

	"time"

	// Internal
	"github.com/deezone/forex-clock/config"

	// Third-party
	"github.com/marksost/go-utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("server.go", func() {
	var (
		// Mock server to test
		s *Server
	)

	BeforeEach(func() {
		// Set config values
		c := config.GetInstance()
		c.Server.Timeouts.Read = 1
		c.Server.Timeouts.Write = 1
	})

	Describe("`NewServer` method", func() {
		It("Returns a valid server", func() {
			// Call method
			s := NewServer()

			// Verify server was properly created and returned
			Expect(s.instance).To(Not(BeNil()))
		})
	})

	Describe("Server struct methods", func() {
		BeforeEach(func() {
			// Get empty port
			port, err := goutils.GetEmptyPort()
			if err != nil {
				panic("Error getting an empty port. Testing cannot continue. Error was: " + err.Error())
			}

			// Set server port
			config.GetInstance().Server.Port = port

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

		Describe("`Start` method", func() {
			Context("When a server is already running", func() {
				It("Returns an error", func() {
					// Call method
					err := s.Start()

					// Verify return value
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When a server is not already running", func() {
				It("Configures and starts a server", func() {
					// Verify configuration values were set on the server
					Expect(s.instance.ReadTimeout).To(Equal(time.Duration(1) * time.Second))
					Expect(s.instance.WriteTimeout).To(Equal(time.Duration(1) * time.Second))
				})
			})
		})

		Describe("`Stop` method", func() {
			Context("When a server is already running", func() {
				It("Returns the response from the server stop call", func() {
					// Call method
					err := s.Stop()

					// Verify return value
					Expect(err).To(Not(HaveOccurred()))
				})
			})

			Context("When a server is not already running", func() {
				BeforeEach(func() {
					// Create server instance
					s = NewServer()
				})

				It("Returns an error", func() {
					// Call method
					err := s.Stop()

					// Verify return value
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Describe("`GetInstance` method", func() {
			It("Returns an golang http server", func() {
				// Verify return value
				Expect(s.GetInstance()).To(Not(BeNil()))
				Expect(s.GetInstance().IdleTimeout).To(Not(BeNil()))
			})
		})

		Describe("`IsRunning` method", func() {
			Context("When a server is already running", func() {
				It("Returns true", func() {
					// Verify return value
					Expect(s.IsRunning()).To(BeTrue())
				})
			})

			Context("When a server is not already running", func() {
				BeforeEach(func() {
					// Create server instance
					s = NewServer()
				})

				It("Returns false", func() {
					// Verify return value
					Expect(s.IsRunning()).To(BeFalse())
				})
			})
		})
	})
})
