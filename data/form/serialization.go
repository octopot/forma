package form

import (
	"bytes"
	"encoding/xml"
	"html/template"
)

var html = template.Must(template.New("form").Parse(`
<form id="{{ .ID }}" title="{{ .Title }}" action="{{ .Action }}" method="{{ .Method }}" enctype="{{ .EncType }}">
{{- $ := . -}}
{{- range .Inputs -}}
    {{- if .Title -}}
        <label {{ if .ID }}for="{{ .ID }}"{{ else }}for="{{ $.ID }}_{{ .Name }}"{{ end }}>{{ .Title }}</label>
    {{- end -}}
    <input {{ if .ID }}id="{{ .ID }}"{{ else }}id="{{ $.ID }}_{{ .Name }}"{{ end }} name="{{ .Name }}" type="{{ .Type }}"
           {{- with .Title }} title="{{ . }}"{{ end -}}
           {{- with .Placeholder }} placeholder="{{ . }}"{{ end -}}
           {{- with .Value }} value="{{ . }}"{{ end -}}
           {{- with .MinLength }} minlength="{{ . }}"{{ end -}}
           {{- with .MaxLength }} maxlength="{{ . }}"{{ end -}}
           {{- with .Required }} required{{ end -}}>
{{- end -}}
    <input type="submit">
</form>
`))

type schema Schema

// MarshalXML implements built-in `encoding/xml.Marshaler` interface.
func (s Schema) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "form"}
	return e.EncodeElement(schema(s), start)
}

// MarshalHTML encodes the schema to HTML.
func (s Schema) MarshalHTML() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := html.Execute(buf, s)
	return buf.Bytes(), err
}
