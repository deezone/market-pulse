// middleware of the HTTP server
// The logger middleware adds logging support to every request
package middleware

import (
	// Standard Lib
	"log"
	"net/http"
	"os"
	"time"
)

const (
	// String defining the format of the logger's date output
	LogDateFormat = "01/02 - 15:04:05"
)

type (
	// Struct representing logger middleware
	Logger struct {
		Log *log.Logger
	}
)

var (
	// Stream to write to for stdout
	// NOTE: Pattern used to allow for test-based overrides of stdout
	OutputStream *os.File = os.Stdout

	// Handler variables
	date, ip, method, path string
	latency                time.Duration
	start, end             time.Time
)

// NewLogger creates and returns a new instance of logger middleware
func NewLogger() Logger {
	return Logger{
		Log: log.New(OutputStream, "", 0),
	}
}

// Handler handles the processing of the request
// The pre-flight middleware handler adds logging support to every request
func (m Logger) Handler(next http.Handler) http.Handler {
	// Middleware handler function
	fn := func(w http.ResponseWriter, req *http.Request) {
		// Set variable values
		ip = req.RemoteAddr
		method = req.Method
		path = req.URL.Path
		start = time.Now()

		// Pass the request through
		// NOTE: Avoids latency introduced by logging
		next.ServeHTTP(w, req)

		end = time.Now() // NOTE: Uses `now` instance of `since` to allow later formatting
		date = end.Format(LogDateFormat)
		latency = end.Sub(start)

		// Print log
		m.Log.Printf("%s %4v %s %s %s\n", date, latency, ip, method, path)
	}

	return http.HandlerFunc(fn)
}
