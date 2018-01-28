package domain

import (
	"strings"
	"unicode"
)

// Schema represents an HTML form.
type Schema struct {
	ID           string  `json:"id,omitempty"      yaml:"id,omitempty"      xml:"id,attr,omitempty"`
	Language     string  `json:"lang"              yaml:"lang"              xml:"lang,attr"`
	Title        string  `json:"title"             yaml:"title"             xml:"title,attr"`
	Action       string  `json:"action"            yaml:"action"            xml:"action,attr"`
	Method       string  `json:"method,omitempty"  yaml:"method,omitempty"  xml:"method,attr,omitempty"`
	EncodingType string  `json:"enctype,omitempty" yaml:"enctype,omitempty" xml:"enctype,attr,omitempty"`
	Inputs       []Input `json:"input"             yaml:"input"             xml:"input"`
}

// Apply uses filtration, normalization, and validation for input values.
func (s Schema) Apply(data map[string][]string) (map[string][]string, ValidationError) {
	return s.Validate(s.Normalize(s.Filter(data)))
}

// Filter applies the schema to input values to remove unspecified of them.
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
func (s Schema) Validate(data map[string][]string) (map[string][]string, ValidationError) {
	if len(s.Inputs) == 0 || len(data) == 0 {
		return data, nil
	}
	rules, index := makeRules(s.Inputs)
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

func makeRules(inputs []Input) (map[string][]Validator, map[string]int) {
	index := make(map[string]int, len(inputs))
	rules := make(map[string][]Validator, len(inputs))
	for i, input := range inputs {
		index[input.Name] = i
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
	return rules, index
}
