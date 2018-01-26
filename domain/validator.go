package domain

import (
	"fmt"
	"strings"
)

const (
	// EmailType specifies `<input type="email">`.
	EmailType = "email"
	// HiddenType specifies `<input type="hidden">`
	HiddenType = "hidden"
	// TextType specifies `<input type="text">`.
	TextType = "text"
)

// ValidationError represents an error related to invalid input values.
type ValidationError interface {
	error
	// InputWithErrors returns map of form inputs and their errors.
	InputWithErrors() map[Input][]error
}

// Validator defines basic behavior of input validators.
type Validator interface {
	// Validate validates the input values.
	Validate(values []string) error
}

// The ValidatorFunc type is an adapter to allow the use of ordinary functions as a validator.
type ValidatorFunc func(values []string) error

// Validate calls fn(values).
func (fn ValidatorFunc) Validate(values []string) error {
	return fn(values)
}

func LengthValidator(min, max int) ValidatorFunc {
	return func(values []string) error {
		for i, value := range values {
			if min != 0 && len(value) < min {
				return validationError{true, i, value, fmt.Sprintf("value length is less than %d", min)}
			}
			if max != 0 && len(value) > max {
				return validationError{true, i, value, fmt.Sprintf("value length is greater than %d", max)}
			}
		}
		return nil
	}
}

func RequireValidator() ValidatorFunc {
	return func(values []string) error {
		if len(values) == 0 {
			return validationError{message: "values are empty"}
		}
		for i, value := range values {
			if value == "" {
				return validationError{single: true, position: i, message: "value is empty"}
			}
		}
		return nil
	}
}

func TypeValidator(inputType string, strict bool) ValidatorFunc {
	return func(values []string) error {
		switch inputType {
		case EmailType:
			for i, value := range values {
				// https://davidcel.is/posts/stop-validating-email-addresses-with-regex/
				if !strings.Contains(value, `@`) {
					return validationError{true, i, value, "value is not a valid email"}
				}
				if strict {
					// TODO v2: support `strict`
					// - net.LookupMX
					// - smtp.Dial
					// - smtp.Client.Hello("checkmail.me")
					// - smtp.Client.Mail(...)
					// - smtp.Client.Rcpt(value)
					// see https://github.com/badoux/checkmail as example
				}
			}
		case HiddenType, TextType:
			// nothing special
		default:
			panic(fmt.Sprintf("not supported input type %q", inputType))
		}
		return nil
	}
}

type validationError struct {
	single   bool
	position int
	value    string
	message  string
}

func (err validationError) Error() string {
	if err.single {
		if err.value != "" {
			return fmt.Sprintf("value %q at position %d is invalid: %s", err.value, err.position, err.message)
		}
		return fmt.Sprintf("value at position %d is invalid: %s", err.position, err.message)
	}
	return err.message
}

type dataValidationError struct {
	dataValidationResult
}

func (dataValidationError) Error() string {
	return "input data has error"
}

func (err dataValidationError) InputWithErrors() map[Input][]error {
	m := make(map[Input][]error, len(err.results))
	for _, r := range err.results {
		m[r.input] = r.errors
	}
	return m
}

type dataValidationResult struct {
	data    map[string][]string
	results []inputValidationResult
}

// AsError converts the result into error if it contains at least one input validation error.
func (r dataValidationResult) AsError() ValidationError {
	for _, sub := range r.results {
		if sub.HasError() {
			return dataValidationError{r}
		}
	}
	return nil
}

type inputValidationResult struct {
	input  Input
	errors []error
}

// HasError returns true if the result contains at least one, not nil error.
func (r inputValidationResult) HasError() bool {
	for _, err := range r.errors {
		if err != nil {
			return true
		}
	}
	return false
}
