package errors

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kamilsk/form-api/data/encoder"
	"github.com/pkg/errors"
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

// IsServer returns true if the error is a server error.
func (e Error) IsServer() bool {
	return e.Code%500 < 100
}

// MarshalTo writes an encoded JSON representation of self to the response writer.
func (e Error) MarshalTo(rw http.ResponseWriter) error {
	rw.Header().Set("Content-Type", encoder.JSON)
	for key, values := range e.Headers {
		for _, value := range values {
			rw.Header().Add(key, value)
		}
	}
	rw.WriteHeader(e.Code)
	return json.NewEncoder(rw).Encode(e)
}

// NotSupportedContentType returns prepared client error related to "Accept" header.
func NotSupportedContentType(supported []string) Error {
	return Error{
		Code:    http.StatusNotAcceptable,
		Message: "Request's header `Accept` does not contain supported MIME type",
		Details: "Please review response's header `Accept` with supported MIME types",
		Headers: http.Header{"Accept": supported},
	}
}

// NotProvidedUUID returns prepared client error related to empty form identifier.
func NotProvidedUUID() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Form UUID is not provided",
		Details: "Please provide UUID compatible with RFC 4122",
	}
}

// InvalidUUID returns prepared client error related to invalid form identifier.
func InvalidUUID() Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid form UUID is provided",
		Details: "Please provide UUID compatible with RFC 4122",
	}
}

// ~~~

// InvalidFormData returns prepared client error.
func (*Error) InvalidFormData(err error) Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Request PostForm is invalid",
		Details: err.Error(),
		cause:   err,
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

// InvalidReferer returns prepared client error.
func (*Error) InvalidReferer(err error) Error {
	return Error{
		Code:    http.StatusBadRequest,
		Message: "Request contains invalid HTTP referer",
		Details: err.Error(),
		cause:   err,
	}
}

// FromGetV1 returns prepared error.
func (*Error) FromGetV1(err error) Error {
	return Error{
		Code:    classify(err),
		Message: "Error occurred",
		Details: err.Error(),
		cause:   err,
	}
}

func classify(err error) int {
	cause := errors.Cause(err)
	switch {
	case strings.Contains(cause.Error(), "sql:"):
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
