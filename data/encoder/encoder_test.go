package encoder_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/kamilsk/form-api/data"
	"github.com/kamilsk/form-api/data/encoder"
	"github.com/kamilsk/form-api/data/form"
	"github.com/stretchr/testify/assert"
)

const UUID data.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

func TestSupport(t *testing.T) {
	for _, tc := range []struct {
		name        string
		contentType string
		expected    bool
	}{
		{"supported, HTML", encoder.HTML, true},
		{"supported, JSON", encoder.JSON, true},
		{"supported, TEXT", encoder.TEXT, true},
		{"supported, XML", encoder.XML, true},
		{"not supported, TOML", "TOML", false},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, encoder.Support(tc.contentType))
		})
	}
}

func TestEncoder(t *testing.T) {
	for _, tc := range []struct {
		name        string
		contentType string
		golden      string
		schema      form.Schema
	}{
		{"email subscription, HTML", encoder.HTML, "./fixtures/email_subscription.html.golden", form.Schema{
			ID:      UUID.String(),
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/" + UUID.String(),
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
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
		{"email subscription, JSON", encoder.JSON, "./fixtures/email_subscription.json.golden", form.Schema{
			ID:      UUID.String(),
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/" + UUID.String(),
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
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
		{"email subscription, XML", encoder.XML, "./fixtures/email_subscription.xml.golden", form.Schema{
			ID:      UUID.String(),
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/" + UUID.String(),
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
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
		{"email subscription, YAML", encoder.TEXT, "./fixtures/email_subscription.yaml.golden", form.Schema{
			ID:      UUID.String(),
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/" + UUID.String(),
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
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
			enc := encoder.New(buf, tc.contentType)
			assert.Equal(t, tc.contentType, enc.ContentType())
			assert.NoError(t, enc.Encode(tc.schema))
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			assert.Equal(t, string(expected), string(buf.Bytes()))
		})
	}
	t.Run("unsupported content type", func(t *testing.T) {
		assert.Panics(t, func() { encoder.New(bytes.NewBuffer(nil), "unsupported") })
	})
	t.Run("unsupported value by HTML encoder", func(t *testing.T) {
		enc := encoder.New(bytes.NewBuffer(nil), encoder.HTML)
		assert.Error(t, enc.Encode("the value does not have `MarshalHTML` method"))
	})
	t.Run("problem writer", func(t *testing.T) {
		var enc encoder.Generic
		enc = encoder.New(writer(func(p []byte) (n int, err error) { return 0, errors.New("problem writer") }), encoder.HTML)
		assert.Error(t, enc.Encode(form.Schema{}))
		enc = encoder.New(writer(func(p []byte) (n int, err error) { return 0, errors.New("problem writer") }), encoder.TEXT)
		assert.Error(t, enc.Encode(form.Schema{}))
	})
}

type writer func(p []byte) (n int, err error)

func (fn writer) Write(p []byte) (n int, err error) {
	return fn(p)
}
