package errors_test

import (
	"testing"

	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	tests := []struct {
		name    string
		panic   func()
		checker func(assert.TestingT, interface{}, ...interface{}) bool
	}{
		{"error with stack trace", func() { panic(errors.Errorf("panic")) }, assert.NotEmpty},
		{"error without stack trace", func() { panic(errors.Simple("panic")) }, assert.Empty},
		{"not error panic", func() { panic("panic") }, assert.Empty},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			var err error
			assert.NotPanics(t, func() {
				defer errors.Recover(&err)
				tc.panic()
			})
			assert.Error(t, err)
			tc.checker(t, errors.StackTrace(err))
		})
	}
}

func TestWrapper(t *testing.T) {
	tests := []struct {
		name string
		wrap func(error) error
	}{
		{"wrap by WithMessage", func(err error) error { return errors.WithMessage(err, "wrapped") }},
		{"wrap by Wrapf", func(err error) error { return errors.Wrapf(err, "wrapped") }},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			var err error
			err = tc.wrap(errors.Simple("test"))
			assert.Error(t, err)
			err = tc.wrap(nil)
			assert.NoError(t, err)
		})
	}
}
