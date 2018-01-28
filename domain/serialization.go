package domain

import (
	"bytes"
	"encoding/xml"
	"html/template"
)

var html = template.Must(template.New("form").
	Parse(`<form {{ with .ID }}id="{{ . }}" {{ end -}}
                 lang="{{ .Language }}" title="{{ .Title }}" action="{{ .Action }}"
                 {{- with .Method }} method="{{ . }}"{{ end -}}
                 {{- with .EncodingType }} enctype="{{ . }}"{{ end -}}>
{{- $ := . -}}
{{- range .Inputs -}}
    {{- if .Title }}<label for="{{ .ID }}">{{ .Title }}</label>{{ end -}}
    <input {{ with .ID }}id="{{ .ID }}" {{ end -}} name="{{ .Name }}" type="{{ .Type }}"
           {{- with .Title }} title="{{ . }}"{{ end -}}
           {{- with .Placeholder }} placeholder="{{ . }}"{{ end -}}
           {{- with .Value }} value="{{ . }}"{{ end -}}
           {{- with .MinLength }} minlength="{{ . }}"{{ end -}}
           {{- with .MaxLength }} maxlength="{{ . }}"{{ end -}}
           {{- with .Required }} required{{ end -}}>
{{- end }}<input type="submit"></form>
`))

type schema Schema

// MarshalXML implements built-in `encoding/xml.Marshaler` interface.
func (s Schema) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{Local: "form"}
	return e.EncodeElement(schema(s), start)
}

// MarshalHTML encodes the schema to HTML by default template.
func (s Schema) MarshalHTML() ([]byte, error) {
	var (
		blob = [1024]byte{}
		raw  = blob[:0]
	)
	buf := bytes.NewBuffer(raw)
	err := html.Execute(buf, s)
	return buf.Bytes(), err
}
