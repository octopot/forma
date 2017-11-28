package form

// Input represents input element in a form.
type Input struct {
	ID        string `json:"id,omitempty"        xml:"id,attr,omitempty"`
	Name      string `json:"name"                xml:"name,attr"`
	Type      string `json:"type"                xml:"type,attr"`
	Title     string `json:"title,omitempty"     xml:"title,attr,omitempty"`
	MinLength int    `json:"minlength,omitempty" xml:"minlength,attr,omitempty"`
	MaxLength int    `json:"maxlength,omitempty" xml:"maxlength,attr,omitempty"`
	Required  bool   `json:"required,omitempty"  xml:"required,attr,omitempty"`
	Value     string `json:"value,omitempty"     xml:"value,attr,omitempty"`
}
