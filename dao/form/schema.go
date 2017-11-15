package form

// Schema represents form specification.
type Schema struct {
	ID     string  `xml:"id,attr"`
	Title  string  `xml:"title,attr"`
	Inputs []Input `xml:"input"`
}
