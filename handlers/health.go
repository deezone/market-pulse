package handlers

import (
	// Standard lib
	"encoding/json"
	"net/http"
	"time"

	// Internal
	"github.com/deezone/forex-clock/config"
	"github.com/deezone/forex-clock/db"
	"github.com/deezone/forex-clock/helpers"

	// Third-party
	"github.com/marksost/go-utils"
	log "github.com/sirupsen/logrus"
)

const (
	// Routes
	HealthRoute  = "/health"
	ReadyRoute   = "/ready"
	VersionRoute = "/version"

	// Ready statuses
	ReadyStatusOK    = "ok"
	ReadyStatusError = "error"
)

type (
	// Struct representing a route handler for health / ready / version routes
	HealthHandler struct {
		db  db.DB   // DB instance that the handler can use
	}
	// HealthResponse is a struct defining properties all "health" responses should contain
	HealthResponse struct {
		Uptime int64 `json:"uptime"`
	}
	// ReadyResponse is a struct defining properties all "ready" responses should contain
	ReadyResponse struct {
		// Embeeded field
		*HealthResponse
		Service string `json:"service"`
		DB      string `json:"db"`
		DBType  string `json:"db-type"`
	}
	// VersionResponse is a struct defining properties all "version" responses should contain
	VersionResponse struct {
		// Embeeded field
		*HealthResponse
		Version string `json:"version"`
	}
)

// NewHealthHandler creates and returns a new instance of a health handler
func NewHealthHandler(db db.DB) *HealthHandler {
	return &HealthHandler{
		db:  db,
	}
}

// NewHealthResponse creates and returns a new instance of a health response
func NewHealthResponse() *HealthResponse {
	return &HealthResponse{
		Uptime: time.Now().Unix() - config.GetInstance().StartTime.Unix(),
	}
}

// NewReadyResponse creates and returns a new instance of a ready response
func NewReadyResponse(db db.DB) *ReadyResponse {
	return &ReadyResponse{
		HealthResponse: NewHealthResponse(),
		Service:        ReadyStatusOK,
		DB:             checkDB(db),
		DBType:         db.String(),
	}
}

// NewVersionResponse creates and returns a new instance of a version response
func NewVersionResponse() *VersionResponse {
	return &VersionResponse{
		HealthResponse: NewHealthResponse(),
		Version:        config.GetInstance().ReleaseVersion,
	}
}

// Health is an http handler used to fulfill "health" requests
func (h HealthHandler) Health(w http.ResponseWriter, req *http.Request) {
	// Check for valid method
	if req.Method != http.MethodGet {
		helpers.MethodNotAllowed(w, req)
		return
	}

	// Use helper response method
	helpers.OK(w, req, NewHealthResponse())
}

// Ready is an http handler used to fulfill "ready" requests
func (h HealthHandler) Ready(w http.ResponseWriter, req *http.Request) {
	// Check for valid method
	if req.Method != http.MethodGet {
		helpers.MethodNotAllowed(w, req)
		return
	}

	// Form response
	resp := NewReadyResponse(h.db)

	// Ready checks to ensure aren't errors
	resps := []string{resp.DB}

	// Check for errors, output 500 response
	if goutils.SliceContains(ReadyStatusError, resps) {
		// Set content type and stauts code
		w.Header().Set("Content-Type", helpers.ResponseContentType)
		w.WriteHeader(http.StatusInternalServerError)

		// Form output
		json, _ := json.Marshal(*resp)

		w.Write([]byte(json))
		return
	}

	// Use helper response method
	helpers.OK(w, req, resp)
}

// Version is an http handler used to fulfill "version" requests
func (h HealthHandler) Version(w http.ResponseWriter, req *http.Request) {
	// Check for valid method
	if req.Method != http.MethodGet {
		helpers.MethodNotAllowed(w, req)
		return
	}

	// Use helper response method
	helpers.OK(w, req, NewVersionResponse())
}

// checkDB performs a PING on a database to ensure it's reachable by the application
func checkDB(db db.DB) string {
	// Check if database is ready
	if err := db.Ready(); err != nil {
		log.WithError(err).Error("Error checking database")
		return ReadyStatusError
	}

	return ReadyStatusOK
}
