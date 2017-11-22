package form

import (
	"encoding/xml"
	"io"
	"net/url"
	"text/template"
)

var tpl = template.Must(template.New("schema.xml").Parse(`{{- $ := . -}}
<form id="{{ .ID }}" title="{{ .Title }}" action="{{ .URL }}" method="post" enctype="application/x-www-form-urlencoded">
    {{- range .Inputs -}}
    <input id="{{ $.ID }}_{{ .Name }}" name="{{ .Name }}" type="{{ .Type }}"
           {{- with .Title }} title="{{ . }}"{{ end -}}
           {{- with .MinLength }} minlength="{{ . }}"{{ end -}}
           {{- with .MaxLength }} maxlength="{{ . }}"{{ end -}}
           {{- with .Value }} value="{{ . }}"{{ end -}}
           {{- with .Required }} required="1"{{ end -}}
    />
    {{- end -}}
</form>
`))

// Schema represents form specification.
type Schema struct {
	ID     string   `xml:"-"`
	URL    *url.URL `xml:"-"`
	Title  string   `xml:"title,attr"`
	Inputs []Input  `xml:"input"`
}

// MarshalTo writes an encoded XML representation of self to the writer.
func (s Schema) MarshalTo(w io.Writer) error {
	return tpl.Execute(w, s)
}

// UnmarshalFrom parses the XML-encoded data and stores the result in self.
func (s *Schema) UnmarshalFrom(data []byte) error {
	return xml.Unmarshal(data, s)
}
