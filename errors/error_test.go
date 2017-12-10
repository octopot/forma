package errors_test

import (
	"fmt"
	"testing"

	deep "github.com/pkg/errors"

	"github.com/kamilsk/form-api/errors"
	"github.com/stretchr/testify/assert"
)

func TestApplicationError(t *testing.T) {
	const emptyMessage = ""
	for _, tc := range []struct {
		name     string
		err      func(userMsg string, cause error, ctxMsg string, ctxArgs ...interface{}) errors.ApplicationError
		expected struct {
			isServer   bool
			isUser     bool
			isNotFound bool
			isInvalid  bool
		}
	}{
		{"not found", errors.NotFound, struct {
			isServer   bool
			isUser     bool
			isNotFound bool
			isInvalid  bool
		}{false, true, true, false}},
		{"validation", errors.Validation, struct {
			isServer   bool
			isUser     bool
			isNotFound bool
			isInvalid  bool
		}{false, true, false, true}},
		{"database", errors.Database, struct {
			isServer   bool
			isUser     bool
			isNotFound bool
			isInvalid  bool
		}{true, false, false, false}},
		{"serialization", errors.Serialization, struct {
			isServer   bool
			isUser     bool
			isNotFound bool
			isInvalid  bool
		}{true, false, false, false}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cause := fmt.Errorf(tc.name)
			err := tc.err(emptyMessage, cause, "~")
			assert.NotEqual(t, cause, err.Cause())
			assert.Equal(t, cause, deep.Cause(err))
			assert.Equal(t, err.IsServer(), tc.expected.isServer)
			assert.Equal(t, err.IsUser(), tc.expected.isUser)
			assert.Equal(t, err.IsNotFound(), tc.expected.isNotFound)
			assert.Equal(t, err.IsInvalid(), tc.expected.isInvalid)
			assert.Equal(t, err.Message(), emptyMessage)
			if err.IsServer() {
				assert.Equal(t, errors.ServerErrorMessage, err.Error())
			} else {
				assert.Equal(t, errors.ClientErrorMessage, err.Error())
			}
		})
	}
}
