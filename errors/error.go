package errors

import "github.com/pkg/errors"

// Server returns an application error related to server side.
func Server() *Error {
	return &Error{server: true}
}

// User returns an application error related to client side.
func User() *Error {
	return &Error{server: false}
}

// Error represents an application error.
type Error struct {
	cause  error
	server bool
}

// Error returns a message of the underlying cause of the error,
// or default message for server or client side error.
func (err *Error) Error() string {
	if err.cause == nil {
		if err.server {
			return "server error"
		}
		return "user error"
	}
	return err.cause.Error()
}

// Cause returns the underlying cause of the error.
// It is friendly to `github.com/pkg/errors.Cause` method.
func (err *Error) Cause() error { return err.cause }

// IsServer returns true if the error on server side.
func (err *Error) IsServer() bool { return err.server }

// IsUser returns true if the error on client side.
func (err *Error) IsUser() bool { return !err.server }

// Wrapf returns an error annotating cause with a stack trace
// at the point Wrapf is call, and the format specifier.
// If cause is nil, Wrapf returns nil.
func (err *Error) Wrapf(cause error, format string, args ...interface{}) error {
	if cause == nil {
		return nil
	}
	err.cause = cause
	return errors.Wrapf(err, format, args...)
}
