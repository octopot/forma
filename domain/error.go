package domain

import "fmt"

// AccumulatedError represents an error related to invalid input values.
type AccumulatedError interface {
	error

	// InputWithErrors returns map of form inputs and their errors.
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

func (err dataValidationError) InputWithErrors() map[Input][]error {
	m := make(map[Input][]error, len(err.results))
	for _, r := range err.results {
		m[r.input] = r.errors
	}
	return m
}

type dataValidationResult struct {
	results []inputValidationResult
}

// AsError converts the result into error if it contains at least one input validation error.
func (r dataValidationResult) AsError() AccumulatedError {
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
