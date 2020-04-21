package domain

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ValidationError defines the behavior of error related to invalid input values.
type ValidationError interface {
	error

	// HasError returns true if the input value has at least one error.
	HasError(input Input) bool
	// InputWithErrors returns a map of form inputs with at least one error and their errors.
	InputWithErrors() map[Input][]error
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
	return "validation error"
}

func (err dataValidationError) HasError(input Input) bool {
	for _, result := range err.results {
		if result.input.Name == input.Name {
			return result.HasError()
		}
	}
	return false
}

func (err dataValidationError) InputWithErrors() map[Input][]error {
	m := make(map[Input][]error, len(err.results))
	for _, r := range err.results {
		if r.HasError() {
			m[r.input] = r.errors
		}
	}
	return m
}

type dataValidationResult struct {
	results []inputValidationResult
}

// AsError converts the result into error if it contains at least one input validation error.
func (r dataValidationResult) AsError() ValidationError {
	for _, result := range r.results {
		if result.HasError() {
			return dataValidationError{r}
		}
	}
	return nil
}

type inputValidationResult struct {
	input  Input
	errors []error
}

// HasError returns true if the result contains at least one error.
func (r inputValidationResult) HasError() bool {
	for _, err := range r.errors {
		if err != nil {
			return true
		}
	}
	return false
}

// Validate gets input metadata, validation rules and checks the input data.
func Validate(inputs []Input, rules map[string][]Validator, data map[string][]string) ValidationError {
	index := make(map[string]int, len(inputs))
	for i, input := range inputs {
		index[input.Name] = i
	}
	validation := dataValidationResult{}
	for name, values := range data {
		i, found := index[name]
		if !found {
			continue
		}
		inputValidation := inputValidationResult{input: inputs[i]}
		validators := rules[name]
		for _, validator := range validators {
			if err := validator.Validate(values); err != nil {
				inputValidation.errors = append(inputValidation.errors, err)
			}
		}
		validation.results = append(validation.results, inputValidation)
	}
	return validation.AsError()
}

// Validator defines the basic behavior of input validators.
type Validator interface {
	// Validate validates the input values.
	Validate(values []string) error
}

// ValidatorFunc type is an adapter to allow the use of ordinary functions as a validator.
type ValidatorFunc func(values []string) error

// Validate calls fn(values).
func (fn ValidatorFunc) Validate(values []string) error {
	return fn(values)
}

// LengthValidator returns the validator to check an input value length.
func LengthValidator(min, max int) ValidatorFunc {
	return func(values []string) error {
		for i, value := range values {
			if min != 0 && utf8.RuneCountInString(value) < min {
				return validationError{true, i, value, fmt.Sprintf("value length is less than %d", min)}
			}
			if max != 0 && utf8.RuneCountInString(value) > max {
				return validationError{true, i, value, fmt.Sprintf("value length is greater than %d", max)}
			}
		}
		return nil
	}
}

// RequireValidator returns the validator to check an input value for a not-empty.
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

// TypeValidator returns the validator to check an input value for compliance the type.
// It can raise the panic if the input type is unsupported.
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
				}
			}
		case HiddenType, TextType:
			// nothing special
		default:
			panic(fmt.Sprintf("input type %q is not supported", inputType))
		}
		return nil
	}
}
