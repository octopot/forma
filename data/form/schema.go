package form

import "encoding/xml"

// Schema represents form specification.
type Schema struct {
	XMLName xml.Name `json:"-"                 yaml:"-"                 xml:"form"`
	ID      string   `json:"id,omitempty"      yaml:"id,omitempty"      xml:"id,attr,omitempty"`
	Title   string   `json:"title"             yaml:"title"             xml:"title,attr"`
	Action  string   `json:"action"            yaml:"action"            xml:"action,attr"`
	Method  string   `json:"method,omitempty"  yaml:"method,omitempty"  xml:"method,attr,omitempty"`
	EncType string   `json:"enctype,omitempty" yaml:"enctype,omitempty" xml:"enctype,attr,omitempty"`
	Inputs  []Input  `json:"input"             yaml:"input"             xml:"input"`
}

// Apply applies the schema to input values to filter them.
// It removes all values not fitted by the schema.
func (s Schema) Apply(data map[string][]string) map[string][]string {
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

// Validate validates input values and returns all occurred errors.
func (s Schema) Validate(data map[string][]string) ValidationError {
	if len(s.Inputs) == 0 || len(data) == 0 {
		return nil
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
	return validation.AsError()
}
