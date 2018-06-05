// Test suite setup for the server package
package server

import (
	// Standard lib
	"io/ioutil"
	"testing"

	// Third-party
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
)

type (
	// Struct representing route integration test data
	RoutesTestData struct {
		Method       string
		Route        string
		ResponseCode int
	}
)

// Tests the server package
func TestServer(t *testing.T) {
	// Register gomega fail handler
	RegisterFailHandler(Fail)

	// Have go's testing package run package specs
	RunSpecs(t, "Server Suite")
}

func init() {
	// Set logger output so as not to log during tests
	log.SetOutput(ioutil.Discard)
}
