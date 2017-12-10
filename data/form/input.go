package form

// Input represents an element of a HTML form.
type Input struct {
	ID          string `json:"id,omitempty"          yaml:"id,omitempty"          xml:"id,attr,omitempty"`
	Name        string `json:"name"                  yaml:"name"                  xml:"name,attr"`
	Type        string `json:"type"                  yaml:"type"                  xml:"type,attr"`
	Title       string `json:"title,omitempty"       yaml:"title,omitempty"       xml:"title,attr,omitempty"`
	Placeholder string `json:"placeholder,omitempty" yaml:"placeholder,omitempty" xml:"placeholder,attr,omitempty"`
	Value       string `json:"value,omitempty"       yaml:"value,omitempty"       xml:"value,attr,omitempty"`
	MinLength   int    `json:"minlength,omitempty"   yaml:"minlength,omitempty"   xml:"minlength,attr,omitempty"`
	MaxLength   int    `json:"maxlength,omitempty"   yaml:"maxlength,omitempty"   xml:"maxlength,attr,omitempty"`
	Required    bool   `json:"required,omitempty"    yaml:"required,omitempty"    xml:"required,attr,omitempty"`
	Strict      bool   `json:"strict,omitempty"      yaml:"strict,omitempty"      xml:"strict,attr,omitempty"`
}
