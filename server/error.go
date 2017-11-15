package server

import (
	"encoding/json"
	"net/http"
)

// ErrorMessage represents HTTP error.
type ErrorMessage struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// NotProvidedUUID returns prepared client error.
func (*ErrorMessage) NotProvidedUUID() ErrorMessage {
	return ErrorMessage{
		Code:    http.StatusBadRequest,
		Message: "Form UUID is not provided",
		Details: "Please pass UUID compatible with RFC 4122",
	}
}

// InvalidUUID returns prepared client error.
func (*ErrorMessage) InvalidUUID() ErrorMessage {
	return ErrorMessage{
		Code:    http.StatusBadRequest,
		Message: "Invalid form UUID is provided",
		Details: "Please pass UUID compatible with RFC 4122",
	}
}

// MarshalTo writes an encoded JSON representation of self to response.
func (e ErrorMessage) MarshalTo(rw http.ResponseWriter) error {
	rw.WriteHeader(e.Code)
	return json.NewEncoder(rw).Encode(e)
}
