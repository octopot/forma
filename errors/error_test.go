package errors_test

import (
	"fmt"
	"testing"

	deep "github.com/pkg/errors"

	"github.com/kamilsk/form-api/errors"
	"github.com/stretchr/testify/assert"
)

func TestApplicationError(t *testing.T) {
	const errorMessage, contextMessage = "error", "context"

	type (
		Result struct {
			isServerError       bool
			isDatabaseFail      bool
			isSerializationFail bool
			isUserError         bool
			isInvalidInput      bool
			isResourceNotFound  bool
		}
		TestCase struct {
			name        string
			constructor func(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) errors.ApplicationError
			expected    Result
		}
	)

	tests := []TestCase{
		{"not found", errors.NotFound, Result{isUserError: true, isResourceNotFound: true}},
		{"validation", errors.Validation, Result{isUserError: true, isInvalidInput: true}},
		{"database", errors.Database, Result{isServerError: true, isDatabaseFail: true}},
		{"serialization", errors.Serialization, Result{isServerError: true, isSerializationFail: true}},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			cause := fmt.Errorf(tc.name)
			err := tc.constructor(errorMessage, cause, contextMessage)
			{
				assert.NotEqual(t, cause, err.Cause())
				assert.Equal(t, cause, deep.Cause(err))
			}
			{
				obtained := Result{}
				if serverErr, is := err.IsServerError(); is {
					obtained.isServerError = is
					obtained.isDatabaseFail = serverErr.IsDatabaseFail()
					obtained.isSerializationFail = serverErr.IsSerializationFail()
				}
				if userErr, is := err.IsClientError(); is {
					obtained.isUserError = is
					obtained.isInvalidInput = userErr.IsInvalidInput()
					obtained.isResourceNotFound = userErr.IsResourceNotFound()
				}
				assert.Equal(t, tc.expected, obtained)
			}
		})
	}
}

func TestApplicationErrorMessage(t *testing.T) {
	const emptyMessage, validationMessage, serializationMessage = "", "validation", "serialization"

	tests := []struct {
		name        string
		constructor func(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) errors.ApplicationError
		userMsg     string
		ctxMsg      string
		cause       func(name string) error
		expected    func() (err, msg string)
	}{
		{"not found", errors.NotFound,
			emptyMessage, "uuid is not presented",
			func(name string) error { return fmt.Errorf(name) },
			func() (string, string) { return "error: uuid is not presented: not found", errors.ClientErrorMessage }},
		{"validation", errors.Validation,
			validationMessage, "invalid email",
			func(name string) error { return nil },
			func() (string, string) { return "validation: <nil>", validationMessage }},
		{"database", errors.Database,
			emptyMessage, "connection is lost",
			func(name string) error { return fmt.Errorf(name) },
			func() (string, string) {
				return "server error: connection is lost: database", errors.ServerErrorMessage
			}},
		{"serialization", errors.Serialization,
			serializationMessage, "corrupted data",
			func(name string) error { return nil },
			func() (string, string) { return "serialization: <nil>", serializationMessage }},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			cause := tc.cause(tc.name)
			err := tc.constructor(tc.userMsg, cause, tc.ctxMsg)
			errMsg, userMsg := tc.expected()
			assert.Equal(t, errMsg, err.Error())
			assert.Equal(t, userMsg, err.Message())
		})
	}
}
