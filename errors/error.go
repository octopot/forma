package errors

import "github.com/pkg/errors"

const (
	ResourceNotFound = 1 + iota
	InvalidInputData
)

const (
	ServerError = 100 + iota
	DatabaseError
	SerializationError
)

// NotFound returns an application error related to an empty search result.
func NotFound(cause error, format string, args ...interface{}) *withCode {
	return &withCode{cause: errors.Wrapf(cause, format, args...), code: ResourceNotFound}
}

// Validation returns an application error related to invalid input values.
func Validation(cause error, format string, args ...interface{}) *withCode {
	return &withCode{cause: errors.Wrapf(cause, format, args...), code: InvalidInputData}
}

// Database returns an application error related to database problems.
func Database(cause error, format string, args ...interface{}) *withCode {
	return &withCode{cause: errors.Wrapf(cause, format, args...), code: DatabaseError}
}

// Serialization returns an application error related to serialization problems.
func Serialization(cause error, format string, args ...interface{}) *withCode {
	return &withCode{cause: errors.Wrapf(cause, format, args...), code: SerializationError}
}

type withCode struct {
	cause error
	code  int
}

// Error returns a message of the underlying cause of the error
// or default message for server or client side error.
func (err *withCode) Error() string {
	if err.cause == nil {
		if err.IsServer() {
			return "server error"
		}
		return "user error"
	}
	return err.cause.Error()
}

// Cause returns the underlying cause of the error.
// It is friendly to `github.com/pkg/errors.Cause` method.
func (err *withCode) Cause() error { return err.cause }

// IsServer returns true if the error on the server side.
func (err *withCode) IsServer() bool { return err.code > ServerError }

// IsUser returns true if the error on the client side.
func (err *withCode) IsUser() bool { return err.code < ServerError }

// IsNotFound returns true if the error related to an empty search result.
func (err *withCode) IsNotFound() bool { return err.code == ResourceNotFound }

// IsInvalid returns true if the error related to invalid input values.
func (err *withCode) IsInvalid() bool { return err.code == InvalidInputData }
