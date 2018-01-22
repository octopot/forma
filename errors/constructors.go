package errors

import "github.com/pkg/errors"

// NotFound returns the application error related to an empty search result.
func NotFound(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError {
	return withCode{ResourceNotFoundCode, userMsg, errors.Wrapf(cause, ctxMsg, ctxArgs...)}
}

// Validation returns the application error related to invalid user input.
func Validation(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError {
	return withCode{InvalidInputCode, userMsg, errors.Wrapf(cause, ctxMsg, ctxArgs...)}
}

// Database returns the application error related to database problems.
func Database(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError {
	return withCode{DatabaseFailCode, userMsg, errors.Wrapf(cause, ctxMsg, ctxArgs...)}
}

// Serialization returns the application error related to serialization problems.
func Serialization(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError {
	return withCode{SerializationFailCode, userMsg, errors.Wrapf(cause, ctxMsg, ctxArgs...)}
}
