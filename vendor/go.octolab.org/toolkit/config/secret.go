package config

// Secret tries to hide the self's content while printing.
type Secret string

// GoString implements fmt.GoStringer.
func (Secret) GoString() string {
	return "***"
}

// MarshalJSON implements json.Marshaler.
func (Secret) MarshalJSON() ([]byte, error) {
	return []byte(`"***"`), nil
}

// MarshalText implements encoding.TextMarshaler.
func (Secret) MarshalText() (text []byte, err error) {
	return []byte("***"), nil
}

// MarshalYAML implements gopkg.in/yaml.Marshaler.
func (Secret) MarshalYAML() (interface{}, error) {
	return "***", nil
}

// String implements fmt.Stringer.
func (Secret) String() string {
	return "***"
}
