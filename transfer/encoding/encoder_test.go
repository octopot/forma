package encoding_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/kamilsk/form-api/domen"
	"github.com/kamilsk/form-api/transfer/encoding"
	"github.com/stretchr/testify/assert"
)

const UUID domen.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestSupport(t *testing.T) {
	for _, tc := range []struct {
		name        string
		contentType string
		expected    bool
	}{
		{"supported, HTML", encoding.HTML, true},
		{"supported, JSON", encoding.JSON, true},
		{"supported, TEXT", encoding.TEXT, true},
		{"supported, XML", encoding.XML, true},
		{"not supported, TOML", "TOML", false},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, encoding.IsSupported(tc.contentType))
		})
	}
}

func TestEncoder(t *testing.T) {
	for _, tc := range []struct {
		name        string
		contentType string
		golden      string
		schema      domen.Schema
	}{
		{"email subscription, HTML", encoding.HTML, "./fixtures/email_subscription.html.golden", domen.Schema{
			ID:           UUID.String(),
			Title:        "Email subscription",
			Action:       "http://localhost:8080/api/v1/" + UUID.String(),
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domen.Input{
				{
					ID:        UUID.String() + "_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    UUID.String() + "__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
		{"email subscription, JSON", encoding.JSON, "./fixtures/email_subscription.json.golden", domen.Schema{
			ID:           UUID.String(),
			Title:        "Email subscription",
			Action:       "http://localhost:8080/api/v1/" + UUID.String(),
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domen.Input{
				{
					ID:        UUID.String() + "_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    UUID.String() + "__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
		{"email subscription, XML", encoding.XML, "./fixtures/email_subscription.xml.golden", domen.Schema{
			ID:           UUID.String(),
			Title:        "Email subscription",
			Action:       "http://localhost:8080/api/v1/" + UUID.String(),
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domen.Input{
				{
					ID:        UUID.String() + "_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    UUID.String() + "__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
		{"email subscription, YAML", encoding.TEXT, "./fixtures/email_subscription.yaml.golden", domen.Schema{
			ID:           UUID.String(),
			Title:        "Email subscription",
			Action:       "http://localhost:8080/api/v1/" + UUID.String(),
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domen.Input{
				{
					ID:        UUID.String() + "_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    UUID.String() + "__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			enc := encoding.NewEncoder(buf, tc.contentType)
			assert.Equal(t, tc.contentType, enc.ContentType())
			assert.NoError(t, enc.Encode(tc.schema))
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			assert.Equal(t, string(expected), string(buf.Bytes()))
		})
	}
	t.Run("unsupported content type", func(t *testing.T) {
		assert.Panics(t, func() { encoding.NewEncoder(bytes.NewBuffer(nil), "unsupported") })
	})
	t.Run("unsupported value by HTML encoder", func(t *testing.T) {
		enc := encoding.NewEncoder(bytes.NewBuffer(nil), encoding.HTML)
		assert.Error(t, enc.Encode("the value does not have `MarshalHTML` method"))
	})
	t.Run("problem writer", func(t *testing.T) {
		var enc encoding.Generic
		enc = encoding.NewEncoder(writer(func(p []byte) (n int, err error) { return 0, errors.New("problem writer") }), encoding.HTML)
		assert.Error(t, enc.Encode(domen.Schema{}))
		enc = encoding.NewEncoder(writer(func(p []byte) (n int, err error) { return 0, errors.New("problem writer") }), encoding.TEXT)
		assert.Error(t, enc.Encode(domen.Schema{}))
	})
}

type writer func(p []byte) (n int, err error)

func (fn writer) Write(p []byte) (n int, err error) {
	return fn(p)
}
