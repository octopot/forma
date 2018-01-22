package errors

import (
	"encoding/json"
	"net/http"

	"github.com/kamilsk/form-api/transfer/encoding"
)

// Error represents HTTP error.
type Error struct {
	Code    int         `json:"-"`
	Message string      `json:"message"`
	Details string      `json:"details"`
	Headers http.Header `json:"-"`

	cause error
}

// Cause returns the underlying cause of the error.
// It is friendly to `github.com/pkg/errors.Cause` method.
func (e Error) Cause() error {
	return e.cause
}

// Error implements built-in `error` interface.
func (e Error) Error() string {
	return e.Message
}

// IsClient returns true if the error is a client error.
func (e Error) IsClient() bool {
	return e.Code%400 < 100
}

// IsServerError returns true if the error is a server error.
func (e Error) IsServer() bool {
	return e.Code%500 < 100
}

// MarshalTo writes an encoded JSON representation of self to the response writer.
func (e Error) MarshalTo(rw http.ResponseWriter) error {
	rw.Header().Set("Content-Type", encoding.JSON)
	for key, values := range e.Headers {
		for _, value := range values {
			rw.Header().Add(key, value)
		}
	}
	rw.WriteHeader(e.Code)
	return json.NewEncoder(rw).Encode(e)
}

// FromGetV1 checks passed error and convert it into HTTP error.
func FromGetV1(err error) Error {
	return Error{
		Code:    http.StatusNotFound,
		Message: "Error occurred",
		Details: err.Error(),
		cause:   err,
	}
}
