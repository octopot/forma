package encoder

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"

	"gopkg.in/yaml.v2"
)

const (
	// HTML MIME type.
	HTML = "text/html"
	// JSON MIME type.
	JSON = "application/json"
	// TEXT MIME type.
	TEXT = "text/plain"
	// XML MIME type.
	XML = "text/xml"
)

var supported = []string{HTML, JSON, TEXT, XML}

// Support returns true if provided content type is supported by encoder.
func Support(contentType string) bool {
	for _, available := range Supported() {
		if available == contentType {
			return true
		}
	}
	return false
}

// Supported returns acceptable content types.
func Supported() []string {
	return supported
}

// Encoder defines basic behavior of encoders.
type Encoder interface {
	// Encode writes the encoding of the value to the stream.
	Encode(interface{}) error
}

// Generic defines basic behavior of the application encoder.
type Generic interface {
	Encoder
	// ContentType returns a content type of the encoder.
	ContentType() string
}

// New returns encoder corresponding to the content type.
// It can raise the panic if the content type is unsupported.
func New(stream io.Writer, contentType string) Generic {
	enc := &encoder{cType: contentType, stream: stream}
	switch contentType {
	case HTML:
		enc.real = &htmlEncoder{stream}
	case JSON:
		enc.real = json.NewEncoder(stream)
	case TEXT:
		enc.real = &yamlEncoder{stream, yaml.Marshal}
	case XML:
		enc.real = xml.NewEncoder(stream)
	default:
		panic(fmt.Sprintf("not supported content type %q", contentType))
	}
	return enc
}

type encoder struct {
	cType  string
	stream io.Writer
	real   Encoder
}

func (enc *encoder) ContentType() string { return enc.cType }

func (enc *encoder) Encode(v interface{}) error { return enc.real.Encode(v) }

type htmlEncoder struct{ stream io.Writer }

func (enc *htmlEncoder) Encode(v interface{}) error {
	marshaler, compatible := v.(interface {
		MarshalHTML() ([]byte, error)
	})
	if !compatible {
		return fmt.Errorf("html encode: the value does not have `MarshalHTML` method")
	}
	b, err := marshaler.MarshalHTML()
	if err != nil {
		return err
	}
	n, err := enc.stream.Write(b)
	if err != nil {
		return err
	}
	if n != len(b) {
		return fmt.Errorf("html encode: data loss when recording")
	}
	return nil
}

type yamlEncoder struct {
	stream  io.Writer
	marshal func(interface{}) ([]byte, error)
}

func (enc *yamlEncoder) Encode(v interface{}) error {
	b, err := enc.marshal(v)
	if err != nil {
		return err
	}
	n, err := enc.stream.Write(b)
	if err != nil {
		return err
	}
	if n != len(b) {
		return fmt.Errorf("yaml encode: data loss when recording")
	}
	return nil
}
