package errors_test

import (
	"fmt"
	"testing"

	deep "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/ecosystem/forma/internal/errors"
)

func TestApplicationError(t *testing.T) {
	const userMsg, ctxMsg = "error", "context"

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
			constructor func(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError
			expected    Result
		}
	)

	tests := []TestCase{
		{"not found", NotFound, Result{isUserError: true, isResourceNotFound: true}},
		{"validation", Validation, Result{isUserError: true, isInvalidInput: true}},
		{"database", Database, Result{isServerError: true, isDatabaseFail: true}},
		{"serialization", Serialization, Result{isServerError: true, isSerializationFail: true}},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			cause := fmt.Errorf(tc.name)
			err := tc.constructor(userMsg, cause, ctxMsg)
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
	tests := []struct {
		name        string
		constructor func(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) ApplicationError
		args        func(name string) (userMsg string, cause error, ctxMsg string)
		expected    func() (err, msg string)
	}{
		{"not found", NotFound,
			func(name string) (string, error, string) {
				return "", fmt.Errorf(name), "uuid is not presented"
			},
			func() (string, string) { return "error: uuid is not presented: not found", ClientErrorMessage }},
		{"validation", Validation,
			func(name string) (string, error, string) { return "validation", nil, "invalid email" },
			func() (string, string) { return "validation", "validation" }},
		{"database", Database,
			func(name string) (string, error, string) { return "", fmt.Errorf(name), "connection is lost" },
			func() (string, string) {
				return "server error: connection is lost: database", ServerErrorMessage
			}},
		{"serialization", Serialization,
			func(name string) (string, error, string) { return "serialization", nil, "corrupted data" },
			func() (string, string) { return "serialization", "serialization" }},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			userMsg, cause, ctxMsg := tc.args(tc.name)
			err := tc.constructor(userMsg, cause, ctxMsg)
			errMsg, userMsg := tc.expected()
			assert.Equal(t, errMsg, err.Error())
			assert.Equal(t, userMsg, err.Message())
		})
	}
}
