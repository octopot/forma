package form

import "encoding/xml"

type schema Schema

func (s Schema) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "form"}
	return e.EncodeElement(schema(s), start)
}
