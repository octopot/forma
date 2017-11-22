package server

import (
	"encoding/json"
	"net/http"
)

// Error represents HTTP error.
type Error struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Details string `json:"details"`

	origin error
}

// Error implements built-in error interface.
func (e Error) Error() string {
	return e.Message
}

// Cause implements `github.com/pkg/errors` causer interface.
func (e Error) Cause() error {
	return e.origin
}

// NotProvidedUUID returns prepared client error.
func (*Error) NotProvidedUUID() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Form UUID is not provided",
		Details: "Please provide UUID compatible with RFC 4122",
	}
}

// InvalidUUID returns prepared client error.
func (*Error) InvalidUUID() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid form UUID is provided",
		Details: "Please provide UUID compatible with RFC 4122",
	}
}

// InvalidFormData returns prepared client error.
func (*Error) InvalidFormData(err error) Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Request PostForm is invalid",
		Details: err.Error(),
		origin:  err,
	}
}

// NoReferer returns prepared client error.
func (*Error) NoReferer() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Request does not contain HTTP referer",
		Details: "Please provide required header",
	}
}

// NoReferer returns prepared client error.
func (*Error) InvalidReferer(err error) Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Request contains invalid HTTP referer",
		Details: err.Error(),
		origin:  err,
	}
}

// IsClient returns true if the error is a client error.
func (e Error) IsClient() bool {
	return e.Code%400 < 100
}

// IsServer returns true if the error is a server error.
func (e Error) IsServer() bool {
	return e.Code%500 < 100
}

// MarshalTo writes an encoded JSON representation of self to the response writer.
func (e Error) MarshalTo(rw http.ResponseWriter) error {
	rw.WriteHeader(e.Code)
	rw.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(e)
}
