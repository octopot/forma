package encoding

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"

	yaml "gopkg.in/yaml.v2"
)

const (
	// HTML MIME type.
	HTML = "text/html"
	// JSON MIME type.
	JSON = "application/json"
	// TEXT MIME type.
	TEXT = "text/plain"
	// XML MIME type.
	XML = "application/xml"
)

// This var is related to `Offers`.
var supported = map[string]func(io.Writer) Encoder{
	HTML: func(stream io.Writer) Encoder { return htmlEncoder{stream} },
	JSON: func(stream io.Writer) Encoder { return json.NewEncoder(stream) },
	TEXT: func(stream io.Writer) Encoder { return yamlEncoder{stream, yaml.Marshal} },
	XML:  func(stream io.Writer) Encoder { return xml.NewEncoder(stream) },
}

// Offers returns supported content types.
func Offers() []string {
	return []string{HTML, JSON, TEXT, XML}
}

// IsSupported returns true if the provided content type is supported by encoder.
func IsSupported(contentType string) bool {
	_, ok := supported[contentType]
	return ok
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

// NewEncoder returns encoder corresponding to the content type.
// It can raise panic if the content type is not supported.
// Use IsSupported first to check that.
func NewEncoder(stream io.Writer, contentType string) Generic {
	if !IsSupported(contentType) {
		panic(fmt.Errorf("not supported content type %q", contentType))
	}
	return encoder{contentType, supported[contentType](stream), stream}
}

type encoder struct {
	contentType string
	real        Encoder
	stream      io.Writer
}

// ContentType TODO issue#173
func (enc encoder) ContentType() string { return enc.contentType }

// Encode TODO issue#173
func (enc encoder) Encode(v interface{}) error { return enc.real.Encode(v) }

type htmlEncoder struct{ stream io.Writer }

// Encode TODO issue#173
func (enc htmlEncoder) Encode(v interface{}) error {
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

// Encode TODO issue#173
func (enc yamlEncoder) Encode(v interface{}) error {
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
