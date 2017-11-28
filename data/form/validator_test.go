package form_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kamilsk/form-api/data/form"
	"github.com/stretchr/testify/assert"
)

func TestLengthValidator(t *testing.T) {
	var validator form.Validator
	for _, tc := range []struct {
		name     string
		min, max int
		input    []string
		error    error
	}{
		{
			name:  "length less than",
			min:   5,
			input: []string{"test"},
			error: fmt.Errorf("value %q at position %d has length less than %d", "test", 0, 5),
		},
		{
			name:  "length greater than",
			max:   3,
			input: []string{"test"},
			error: fmt.Errorf("value %q at position %d has length greater than %d", "test", 0, 3),
		},
		{
			name:  "valid",
			min:   3,
			max:   5,
			input: []string{"test"},
			error: nil,
		},
	} {
		validator = form.LengthValidator(tc.min, tc.max)
		assert.Equal(t, tc.error, validator.Validate(tc.input), fmt.Sprintf("test case %q failed", tc.name))
	}
}

func TestRequireValidator(t *testing.T) {
	var validator form.Validator
	for _, tc := range []struct {
		name  string
		input []string
		error error
	}{
		{
			name:  "nil values",
			input: nil,
			error: errors.New("values are empty"),
		},
		{
			name:  "empty values",
			input: []string{},
			error: errors.New("values are empty"),
		},
		{
			name:  "empty value at position 2",
			input: []string{"go", "test", ""},
			error: fmt.Errorf("value at position %d is empty", 2),
		},
		{
			name:  "valid",
			input: []string{"go", "test"},
			error: nil,
		},
	} {
		validator = form.RequireValidator
		assert.Equal(t, tc.error, validator.Validate(tc.input), fmt.Sprintf("test case %q failed", tc.name))
	}
}

func TestTypeValidator(t *testing.T) {
	var validator form.Validator
	for _, tc := range []struct {
		name      string
		inputType string
		input     []string
		error     error
	}{
		{
			name:      "email validation: invalid email at position 2",
			inputType: form.EmailType,
			input:     []string{"no-reply@github.com", "no-reply@golang.org", "@golang.org"},
			error:     fmt.Errorf("value %q at position %d is not a valid email", "@golang.org", 2),
		},
		{
			name:      "email validation: valid",
			inputType: form.EmailType,
			input:     []string{"no-reply@github.com", "no-reply@golang.org"},
			error:     nil,
		},
		{
			name:      "custom validation: not supported",
			inputType: "custom",
			input:     []string{"test"},
			error:     fmt.Errorf("not supported input type %q", "custom"),
		},
	} {
		validator = form.TypeValidator(tc.inputType)
		assert.Equal(t, tc.error, validator.Validate(tc.input), fmt.Sprintf("test case %q failed", tc.name))
	}
}
