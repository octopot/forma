package form

// Validator defines behavior of input validators.
type Validator interface {
	Validate(values []string) error
}

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
