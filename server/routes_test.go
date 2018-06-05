// Tests the routes.go file
package server

import (
	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("routes.go", func() {
	var (
		// Mock server to test
		s *Server
	)

	BeforeEach(func() {
		// Create server instance
		s = NewServer()
	})

	Describe("Server struct methods", func() {
		Describe("`SetRoutes` method", func() {
			It("Sets up a mux handler and adds routes to it", func() {
				// Call method
				s.SetRoutes()

				// Verify routes were set
				Expect(s.GetInstance().Handler).To(Not(BeNil()))
			})
		})
	})
})
