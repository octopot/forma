// +build draft

package draft

import "github.com/kamilsk/form-api/pkg/domain"

// TODO issue#refactoring
// This is a preparation to separate the domain structures
// from their representation.

type JSONInput struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Title       string `json:"title,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Value       string `json:"value,omitempty"`
	MinLength   int    `json:"minlength,omitempty"`
	MaxLength   int    `json:"maxlength,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Strict      bool   `json:"strict,omitempty"`
}

type JSONSchema struct {
	ID           string      `json:"id,omitempty"`
	Language     string      `json:"lang"`
	Title        string      `json:"title"`
	Action       string      `json:"action"`
	Method       string      `json:"method,omitempty"`
	EncodingType string      `json:"enctype,omitempty"`
	Inputs       []JSONInput `json:"input"`
}

func ConvertSchemaToJSON(in struct {
	ID           string
	Language     string
	Title        string
	Action       string
	Method       string
	EncodingType string
	Inputs       []domain.Input
}) JSONSchema {
	for _, input := range in.Inputs {
		ConvertInputToJSON(input)
	}
	return JSONSchema{}
}

func ConvertInputToJSON(in struct {
	ID          string
	Name        string
	Type        string
	Title       string
	Placeholder string
	Value       string
	MinLength   int
	MaxLength   int
	Required    bool
	Strict      bool
}) JSONInput {
	return JSONInput{}
}

func init() {
	ConvertSchemaToJSON(domain.Schema{})
}
