package form

import "encoding/xml"

// Schema represents form specification.
type Schema struct {
	XMLName xml.Name `json:"-" xml:"form"`
	ID      string   `json:"id,omitempty"      xml:"id,attr,omitempty"`
	Title   string   `json:"title"             xml:"title,attr"`
	Action  string   `json:"action"            xml:"action,attr"`
	Method  string   `json:"method,omitempty"  xml:"method,attr,omitempty"`
	EncType string   `json:"enctype,omitempty" xml:"enctype,attr,omitempty"`
	Inputs  []Input  `json:"input"             xml:"input"`
}

// Apply applies the schema to input values to filter them.
// It removes all values not fitted by the schema.
func (s Schema) Apply(in map[string][]string) map[string][]string {
	if len(s.Inputs) == 0 || len(in) == 0 {
		return nil
	}
	index := make(map[string]struct{}, len(s.Inputs))
	for _, input := range s.Inputs {
		index[input.Name] = struct{}{}
	}
	out := make(map[string][]string)
	for name, values := range in {
		if _, ok := index[name]; ok {
			out[name] = values
		}
	}
	return out
}

// Validate validates input values and returns all occurred errors.
func (s Schema) Validate(in map[string][]string) []error {
	if len(s.Inputs) == 0 || len(in) == 0 {
		return nil
	}
	rules := make(map[string][]Validator, len(s.Inputs))
	for _, input := range s.Inputs {
		validators := []Validator{
			TypeValidator(input.Type),
		}
		if input.MinLength != 0 || input.MaxLength != 0 {
			validators = append(validators, LengthValidator(input.MinLength, input.MaxLength))
		}
		if input.Required {
			validators = append(validators, RequireValidator)
		}
		rules[input.Name] = validators
	}

	// TODO not implemented yet
	errors := make([]error, 0)
	for name, values := range in {
		validators := rules[name]
		for _, validator := range validators {
			if err := validator.Validate(values); err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}
