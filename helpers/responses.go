// helpers package contains helper functions and structs to use throughout the application
// responses contains standard response functions to use within HTTP handlers
package helpers

import (
	// Standard lib
	"encoding/json"
	"net/http"
)

const (
	// Response content type to use
	ResponseContentType = "application/json"
)

type (
	// Meta informatation about the collection response
	CollectionMeta struct {
		Count int `json:"count"`
	}
	// Meta informatation about the error response
	ErrorMeta struct{}
	// Meta informatation about the resource response
	ResourceMeta struct{}
	// Error is a struct that represents a single error to be used in an "error" response
	Error struct {
		Message string `json:"message"`
	}
	// CollectionResponse is a struct defining properties all "collection" responses should contain
	CollectionResponse struct {
		Code int             `json:"-"`    // HTTP status code
		Meta *CollectionMeta `json:"meta"` // Meta informatation about the response
		Data []interface{}   `json:"data"` // A slice of zero or more resources
	}
	// ErrorResponse is a struct defining properties all "error" responses should contain
	ErrorResponse struct {
		Code   int        `json:"-"`      // HTTP status code
		Meta   *ErrorMeta `json:"meta"`   // Meta informatation about the response
		Errors []*Error   `json:"errors"` // A slice of zero or more errors
	}
	// ResourceResponse is a struct defining properties all "resource" responses should contain
	ResourceResponse struct {
		Code int           `json:"-"`    // HTTP status code
		Meta *ResourceMeta `json:"meta"` // Meta informatation about the response
		Data interface{}   `json:"data"` // A single resource
	}
)

var (
	// Common status code responses
	BadRequestResponse = &ErrorResponse{ // 400
		Code:   http.StatusBadRequest,
		Meta:   &ErrorMeta{},
		Errors: make([]*Error, 0),
	}
	CreatedResponseResource = &ResourceResponse{ // 201
		Code: http.StatusCreated,
		Meta: &ResourceMeta{},
	}
	MethodNotAllowedResponse = &ResourceResponse{ // 405
		Code: http.StatusMethodNotAllowed,
		Meta: &ResourceMeta{},
	}
	NotFoundResponse = &ResourceResponse{ // 404
		Code: http.StatusNotFound,
		Meta: &ResourceMeta{},
	}
	NotImplementedResponse = &ResourceResponse{ // 501
		Code: http.StatusNotImplemented,
		Meta: &ResourceMeta{},
	}
	OKResponse = &CollectionResponse{ // 200
		Code: http.StatusOK,
		Meta: &CollectionMeta{},
		Data: make([]interface{}, 0),
	}
	OKResponseResource = &ResourceResponse{ // 200
		Code: http.StatusOK,
		Meta: &ResourceMeta{},
	}
	ServerErrorResponse = &ResourceResponse{ // 500
		Code: http.StatusInternalServerError,
		Meta: &ResourceMeta{},
	}
	UnauthorizedResponse = &ResourceResponse{ // 401
		Code: http.StatusUnauthorized,
		Meta: &ResourceMeta{},
	}
)

// BadRequest sends a Bad Request response with JSON-encoded body
func BadRequest(w http.ResponseWriter, req *http.Request, errors []*Error) {
	// Set content type and stauts code
	w.Header().Set("Content-Type", ResponseContentType)
	w.WriteHeader(http.StatusBadRequest)

	// Set errors on response
	BadRequestResponse.Errors = errors

	// Form output
	json, _ := json.Marshal(*BadRequestResponse)

	w.Write([]byte(json))
}

// Created sends an Created response with JSON-encoded body
func Created(w http.ResponseWriter, req *http.Request, data interface{}) {
	// Set content type and status code
	w.Header().Set("Content-Type", ResponseContentType)
	w.WriteHeader(http.StatusCreated)

	// Set body of response
	CreatedResponseResource.Data = data

	// Form output
	json, _ := json.Marshal(*CreatedResponseResource)

	w.Write([]byte(json))
}

// InternalError sends an Internal Server Error response with JSON-encoded body
func InternalError(w http.ResponseWriter, req *http.Request) {
	// Set content type and stauts code
	w.Header().Set("Content-Type", ResponseContentType)
	w.WriteHeader(http.StatusInternalServerError)

	// Form output
	json, _ := json.Marshal(*ServerErrorResponse)

	w.Write([]byte(json))
}

// MethodNotAllowed sends a Method Not Allowed response with JSON-encoded body
func MethodNotAllowed(w http.ResponseWriter, req *http.Request) {
	// Set content type and stauts code
	w.Header().Set("Content-Type", ResponseContentType)
	w.WriteHeader(http.StatusMethodNotAllowed)

	// Form output
	json, _ := json.Marshal(*MethodNotAllowedResponse)

	w.Write([]byte(json))
}

// MethodNotImplemented sends a Method Not Implemented response with JSON-encoded body
func MethodNotImplemented(w http.ResponseWriter, req *http.Request) {
	// Set content type and stauts code
	w.Header().Set("Content-Type", ResponseContentType)
	w.WriteHeader(http.StatusNotImplemented)

	// Form output
	json, _ := json.Marshal(*NotImplementedResponse)

	w.Write([]byte(json))
}

// NotFound sends a Not Found response with JSON-encoded body
func NotFound(w http.ResponseWriter, req *http.Request) {
	// Set content type and stauts code
	w.Header().Set("Content-Type", ResponseContentType)
	w.WriteHeader(http.StatusNotFound)

	// Form output
	json, _ := json.Marshal(*NotFoundResponse)

	w.Write([]byte(json))
}

// OK sends an OK response with JSON-encoded body
func OK(w http.ResponseWriter, req *http.Request, data interface{}) {
	// Set content type and status code
	w.Header().Set("Content-Type", ResponseContentType)
	w.WriteHeader(http.StatusOK)

	// Set body of response
	OKResponseResource.Data = data

	// Form output
	json, _ := json.Marshal(*OKResponseResource)

	w.Write([]byte(json))
}
