// middleware of the HTTP server
// Checks for and normalizes an API version for each request
package middleware

import (
	// Standard Lib
	"fmt"
	"net/http"
	"regexp"
	"strings"

	// Internal
	"github.com/deezone/forex-clock/config"
	"github.com/deezone/forex-clock/helpers"

	// Third-party
	goutils "github.com/marksost/go-utils"
)

const (
	AcceptHeaderPattern = "application/vnd.soulcycle.fed.v(\\d+(?:\\.\\d+)?)\\+json"
	VersionPattern      = "\\d+\\.\\d+"
)

type (
	// Struct representing version middleware
	Version struct{}
)

var (
	AcceptHeaderRegex *regexp.Regexp // Regex to check for valid Accept headers
	VersionRegex      *regexp.Regexp // Regex to get version from Accept header

	// Handler variables
	header  string
	matches []string
	version = config.Version10
)

// NewVersion creates and returns a new instance of version middleware
func NewVersion() Version { return Version{} }

// Handler handles the processing of the request
// The version middleware handler checks for an API version
// within the `Accept` header and parses it to set the version
// of the current request
func (m Version) Handler(next http.Handler) http.Handler {
	// Middleware handler function
	fn := func(w http.ResponseWriter, req *http.Request) {
		// Get accept header as a string and match via regex
		header = string(req.Header.Get("Accept"))
		matches = AcceptHeaderRegex.FindStringSubmatch(header)

		// Check if there's a match, reset version
		if len(matches) > 1 {
			version = matches[1]
		}

		// Check if version matches version regex
		// NOTE: If it doesn't, that means the "minor" part of the version is missing
		// and needs it appended
		if !VersionRegex.MatchString(version) {
			version += ".0"
		}

		// Check that version is supported
		if !goutils.SliceContains(version, config.SupportedVersions) {
			// Output bad request
			msg := fmt.Sprintf("Unsupported version: %s. Supported versions: %s", version, strings.Join(config.SupportedVersions, ", "))
			helpers.BadRequest(w, req, []*helpers.Error{
				&helpers.Error{
					Message: msg,
				},
			})
			return
		}

		// Set version header
		w.Header().Set("X-Detected-Version", version)

		// Pass the request through
		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}

func init() {
	// Compile regexes once
	AcceptHeaderRegex = regexp.MustCompile(AcceptHeaderPattern)
	VersionRegex = regexp.MustCompile(VersionPattern)
}
