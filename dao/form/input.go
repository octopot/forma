package form

// Input represents input element in form.
type Input struct {
	ID        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	Type      string `xml:"type,attr"`
	Title     string `xml:"title,attr"`
	MinLength int    `xml:"minlength,attr,omitempty"`
	MaxLength int    `xml:"maxlength,attr,omitempty"`
}
