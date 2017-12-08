package form

import (
	"strings"
	"unicode"
)

// Schema represents form specification.
type Schema struct {
	ID      string  `json:"id,omitempty"      yaml:"id,omitempty"      xml:"id,attr,omitempty"`
	Title   string  `json:"title"             yaml:"title"             xml:"title,attr"`
	Action  string  `json:"action"            yaml:"action"            xml:"action,attr"`
	Method  string  `json:"method,omitempty"  yaml:"method,omitempty"  xml:"method,attr,omitempty"`
	EncType string  `json:"enctype,omitempty" yaml:"enctype,omitempty" xml:"enctype,attr,omitempty"`
	Inputs  []Input `json:"input"             yaml:"input"             xml:"input"`
}

// Apply uses filtration, normalization and validation for input values.
func (s Schema) Apply(data map[string][]string) (map[string][]string, ValidationError) {
	return s.Validate(s.Normalize(s.Filter(data)))
}

// Filter applies the schema to input values to filter them.
// It omits all values not fitted by the schema.
func (s Schema) Filter(data map[string][]string) map[string][]string {
	if len(s.Inputs) == 0 || len(data) == 0 {
		return nil
	}
	index := make(map[string]struct{}, len(s.Inputs))
	for _, input := range s.Inputs {
		index[input.Name] = struct{}{}
	}
	filtered := make(map[string][]string)
	for name, values := range data {
		if _, ok := index[name]; ok {
			filtered[name] = values
		}
	}
	return filtered
}

// Normalize removes unnecessary characters from input values.
func (s Schema) Normalize(data map[string][]string) map[string][]string {
	for _, values := range data {
		for i, value := range values {
			values[i] = strings.TrimFunc(value, func(r rune) bool {
				if unicode.IsSpace(r) {
					return true
				}
				// U+200B ZeroWidth
				if r == '\u200B' {
					return true
				}
				// U+200C ZeroWidthNoJoiner
				if r == '\u200C' {
					return true
				}
				// U+200D ZeroWidthJoiner
				if r == '\u200D' {
					return true
				}
				// U+2060 WordJoiner
				if r == '\u2060' {
					return true
				}
				return false
			})
		}
	}
	return data
}

// Validate checks input values.
func (s Schema) Validate(data map[string][]string) (map[string][]string, ValidationError) {
	if len(s.Inputs) == 0 || len(data) == 0 {
		return data, nil
	}
	index := make(map[string]int, len(s.Inputs))
	rules := make(map[string][]Validator, len(s.Inputs))
	for i, input := range s.Inputs {
		index[input.Name] = i
		validators := []Validator{
			typeValidator(input.Type, input.Strict),
		}
		if input.MinLength != 0 || input.MaxLength != 0 {
			validators = append(validators, lengthValidator(input.MinLength, input.MaxLength))
		}
		if input.Required {
			validators = append(validators, requireValidator)
		}
		rules[input.Name] = validators
	}
	validation := dataValidationResult{data: data}
	for name, values := range data {
		i, found := index[name]
		if !found {
			continue
		}
		inputValidation := inputValidationResult{input: s.Inputs[i]}
		validators := rules[name]
		for _, validator := range validators {
			if err := validator.Validate(values); err != nil {
				inputValidation.errors = append(inputValidation.errors, err)
			}
		}
		validation.results = append(validation.results, inputValidation)
	}
	return data, validation.AsError()
}
