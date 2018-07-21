package config

// Secret tries to hide the self's content while printing.
// Based on my presentation
// - https://speakerdeck.com/kamilsk/octolab-cookbook-go-lang-tips-and-tricks-protection-of-sensitive-config-data
// - https://www.slideshare.net/ssuserb6ada9/octolab-cookbook-go-lang-tips-and-tricks-protection-of-sensitive-config-data
type Secret string

// GoString implements `fmt.GoStringer`.
func (Secret) GoString() string {
	return "***"
}

// MarshalJSON implements `json.Marshaler`.
func (Secret) MarshalJSON() ([]byte, error) {
	return []byte(`"***"`), nil
}

// MarshalText implements `encoding.TextMarshaler`
func (Secret) MarshalText() (text []byte, err error) {
	return []byte("***"), nil
}

// MarshalYAML implements `gopkg.in/yaml.Marshaler`.
func (Secret) MarshalYAML() (interface{}, error) {
	return "***", nil
}

// String implements `fmt.Stringer`.
func (Secret) String() string {
	return "***"
}
