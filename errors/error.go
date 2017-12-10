package errors

import "github.com/pkg/errors"

const (
	// ClientError is a code of client error.
	ClientError = iota
	// ResourceNotFound is a code of client error when the requested resource is not found.
	ResourceNotFound
	// InvalidInputData is a code of client error when data provided by a user is invalid.
	InvalidInputData
)

const (
	// ServerError is a code of server error.
	ServerError = 100 + iota
	// DatabaseFail is a code of server error related to database problems.
	DatabaseFail
	// SerializationFail is a code of server error related to serialization problems.
	SerializationFail
)

const (
	// ClientErrorMessage is a default message for client error.
	ClientErrorMessage = "Error"
	// ServerErrorMessage is a default message for server error.
	ServerErrorMessage = "Server Error"
	// NeutralMessage is a default message.
	NeutralMessage = "Something went wrong"
	// FormInvalidMessage is a default message in case when input values are invalid.
	FormInvalidMessage = "Form data contains error"
	// SchemaNotFoundMessage is a default message in case when schema is not found.
	SchemaNotFoundMessage = "Schema not found"
)

// ApplicationError defines behavior of application errors.
type ApplicationError interface {
	error
	// Cause returns the underlying cause of the error.
	Cause() error
	// IsUser returns true if the error on the client side.
	IsUser() bool
	// IsNotFound returns true if the error related to an empty search result.
	IsNotFound() bool
	// IsInvalid returns true if the error related to invalid data provided by a user.
	IsInvalid() bool
	// IsServer returns true if the error on the server side.
	IsServer() bool
	// Message returns an error message intended for a user.
	Message() string
}

// NotFound returns an application error related to an empty search result.
func NotFound(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError {
	return &withCode{ResourceNotFound, userMsg, errors.Wrapf(cause, ctxMsg, ctxArgs...)}
}

// Validation returns an application error related to invalid input values.
func Validation(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError {
	return &withCode{InvalidInputData, userMsg, errors.Wrapf(cause, ctxMsg, ctxArgs...)}
}

// Database returns an application error related to database problems.
func Database(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError {
	return &withCode{DatabaseFail, userMsg, errors.Wrapf(cause, ctxMsg, ctxArgs...)}
}

// Serialization returns an application error related to serialization problems.
func Serialization(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError {
	return &withCode{SerializationFail, userMsg, errors.Wrapf(cause, ctxMsg, ctxArgs...)}
}

type withCode struct {
	code  int
	msg   string
	cause error
}

func (err *withCode) Error() string {
	msg := err.msg
	if msg == "" {
		if err.IsServer() {
			msg = "Server Error"
		} else {
			msg = "Error"
		}
	}
	return msg
}

// Message returns an error message intended for a user.
func (err *withCode) Message() string { return err.msg }

// Cause returns the underlying cause of the error.
// It is friendly to `github.com/pkg/errors.Cause` method.
func (err *withCode) Cause() error { return err.cause }

// IsServer returns true if the error on the server side.
func (err *withCode) IsServer() bool { return err.code > ServerError }

// IsUser returns true if the error on the client side.
func (err *withCode) IsUser() bool { return err.code < ServerError }

// IsNotFound returns true if the error related to an empty search result.
func (err *withCode) IsNotFound() bool { return err.code == ResourceNotFound }

// IsInvalid returns true if the error related to invalid data provided by a user.
func (err *withCode) IsInvalid() bool { return err.code == InvalidInputData }
