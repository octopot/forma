package form

// Schema represents form specification.
type Schema struct {
	ID      string  `xml:"id,attr,omitempty"`
	Title   string  `xml:"title,attr"`
	Action  string  `xml:"action,attr"`
	Method  string  `xml:"method,attr,omitempty"`
	EncType string  `xml:"enctype,attr,omitempty"`
	Inputs  []Input `xml:"input"`
}
