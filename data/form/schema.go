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
