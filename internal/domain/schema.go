package domain

import (
	"strings"
	"unicode"
)

// Schema represents an HTML form.
//go:generate easyjson -all
type Schema struct {
	ID           string  `json:"id,omitempty"      yaml:"id,omitempty"      xml:"id,attr,omitempty"`
	Language     string  `json:"lang,omitempty"    yaml:"lang,omitempty"    xml:"lang,attr,omitempty"`
	Title        string  `json:"title,omitempty"   yaml:"title,omitempty"   xml:"title,attr,omitempty"`
	Action       string  `json:"action,omitempty"  yaml:"action,omitempty"  xml:"action,attr,omitempty"`
	Method       string  `json:"method,omitempty"  yaml:"method,omitempty"  xml:"method,attr,omitempty"`
	EncodingType string  `json:"enctype,omitempty" yaml:"enctype,omitempty" xml:"enctype,attr,omitempty"`
	Inputs       []Input `json:"input"             yaml:"input"             xml:"input"`
}

// IsEmpty TODO issue#173
func (s *Schema) IsEmpty() bool {
	return s.ID == "" && s.Language == "" && s.Title == "" && len(s.Inputs) == 0
}

// Apply uses filtration, normalization, and validation for input values.
// It can raise the panic if the input type is unsupported.
func (s *Schema) Apply(data map[string][]string) (map[string][]string, ValidationError) {
	data, err := s.Validate(s.Normalize(s.Filter(data)))
	for i, input := range s.Inputs {
		if values, found := data[input.Name]; found && len(values) > 0 {
			s.Inputs[i].Value = values[0]
		}
	}
	return data, err
}

// Filter applies the schema to input values to remove unspecified from them.
func (s *Schema) Filter(data map[string][]string) map[string][]string {
	if len(s.Inputs) == 0 || len(data) == 0 {
		return nil
	}
	index := make(map[string]struct{}, len(s.Inputs))
	for _, input := range s.Inputs {
		index[input.Name] = struct{}{}
	}
	filtered := make(map[string][]string)
	for name, values := range data {
		if _, found := index[name]; found {
			filtered[name] = values
		}
	}
	return filtered
}

// Normalize removes unnecessary characters from input values.
// TODO move to InputData
func (s *Schema) Normalize(data map[string][]string) map[string][]string {
	for _, values := range data {
		for i, value := range values {
			values[i] = strings.TrimFunc(value, func(r rune) bool {
				if unicode.IsSpace(r) {
					return true
				}
				// https://www.compart.com/en/unicode/U+200B, Zero Width Space
				if r == '\u200B' {
					return true
				}
				// https://www.compart.com/en/unicode/U+200C, Zero Width Non-joiner
				if r == '\u200C' {
					return true
				}
				// https://www.compart.com/en/unicode/U+200D, Zero Width Joiner
				if r == '\u200D' {
					return true
				}
				// https://www.compart.com/en/unicode/U+2060, Word Joiner
				if r == '\u2060' {
					return true
				}
				return false
			})
		}
	}
	return data
}

// Validate checks input values for errors.
// It can raise the panic if the input type is unsupported.
// TODO use InputData
func (s *Schema) Validate(data map[string][]string) (map[string][]string, ValidationError) {
	if len(s.Inputs) == 0 || len(data) == 0 {
		return data, nil
	}
	rules := make(map[string][]Validator, len(s.Inputs))
	for _, input := range s.Inputs {
		validators := make([]Validator, 0, 3)
		validators = append(validators, TypeValidator(input.Type, input.Strict))
		if input.MinLength != 0 || input.MaxLength != 0 {
			validators = append(validators, LengthValidator(input.MinLength, input.MaxLength))
		}
		if input.Required {
			validators = append(validators, RequireValidator())
		}
		rules[input.Name] = validators
	}
	return data, Validate(s.Inputs, rules, data)
}

// Input searches an Input by its name in a case-insensitive manner and returns it
// or nil if can't find it. Useful in templates:
//
//     {{ with .Schema.Input "email" }}
//         {{ template "input" . }}
//     {{ end }}
//
func (s *Schema) Input(name string) *Input {
	for i, input := range s.Inputs {
		if strings.EqualFold(input.Name, name) {
			return &s.Inputs[i]
		}
	}
	return nil
}
