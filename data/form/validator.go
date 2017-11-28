package form

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	EmailType = "email"
)

// Validator defines behavior of input validators.
type Validator interface {
	// Validate ...
	Validate(values []string) error
}

// ValidatorFunc ...
type ValidatorFunc func(values []string) error

// Validate ...
func (fn ValidatorFunc) Validate(values []string) error {
	return fn(values)
}

var (
	// LengthValidator ...
	LengthValidator = func(min, max int) ValidatorFunc {
		return func(values []string) error {
			for i, value := range values {
				if min != 0 && len(value) < min {
					return fmt.Errorf("value %q at position %d has length less than %d", value, i, min)
				}
				if max != 0 && len(value) > max {
					return fmt.Errorf("value %q at position %d has length greater than %d", value, i, max)
				}
			}
			return nil
		}
	}
	// RequireValidator ...
	RequireValidator ValidatorFunc = func(values []string) error {
		if len(values) == 0 {
			return errors.New("values are empty")
		}
		for i, value := range values {
			if value == "" {
				return fmt.Errorf("value at position %d is empty", i)
			}
		}
		return nil
	}
	// TypeValidator ...
	TypeValidator = func(inputType string) ValidatorFunc {
		var (
			email = regexp.MustCompile(`(?i:^[^@]+@[^@]+$)`) // TODO replace by correct method
		)
		return func(values []string) error {
			switch inputType {
			case EmailType:
				for i, value := range values {
					if !email.MatchString(value) {
						return fmt.Errorf("value %q at position %d is not a valid email", value, i)
					}
				}
			default:
				return fmt.Errorf("not supported input type %q", inputType)
			}
			return nil
		}
	}
)

/* TODO not implemented yet
type validationError struct {
	input  Input
	data   []string
	errors []error
}

func (err validationError) Error() string {
	return fmt.Sprintf("input %q is not valid", err.input.Name)
}
*/
