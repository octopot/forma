package form

// Input represents input element in a form.
type Input struct {
	ID        string `xml:"id,attr,omitempty"`
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr"`
	Title     string `xml:"title,attr"`
	MinLength int    `xml:"minlength,attr,omitempty"`
	MaxLength int    `xml:"maxlength,attr,omitempty"`
	Required  bool   `xml:"required,attr,omitempty"`
	Value     string `xml:"value,attr,omitempty"`
}
