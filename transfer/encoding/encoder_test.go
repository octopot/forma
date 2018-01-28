package encoding_test

import (
	"bytes"
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kamilsk/form-api/domain"
	"github.com/kamilsk/form-api/transfer/encoding"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update .golden files")

func TestSupport(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		expected    bool
	}{
		{"supported, HTML", encoding.HTML, true},
		{"supported, JSON", encoding.JSON, true},
		{"supported, TEXT", encoding.TEXT, true},
		{"supported, XML", encoding.XML, true},
		{"unsupported, TOML", "application/toml", false},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, encoding.IsSupported(tc.contentType))
		})
	}
}

func TestEncoder(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		golden      string
		schema      domain.Schema
	}{
		{"email subscription, HTML", encoding.HTML, "./fixtures/email_subscription.html.golden", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
			},
		}},
		{"email subscription, JSON", encoding.JSON, "./fixtures/email_subscription.json.golden", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
			},
		}},
		{"email subscription, XML", encoding.XML, "./fixtures/email_subscription.xml.golden", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
			},
		}},
		{"email subscription, YAML", encoding.TEXT, "./fixtures/email_subscription.yaml.golden", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
			},
		}},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if *update {
				file := writer(tc.golden)
				assert.NoError(t, closeAfter(file, func() error {
					return encoding.NewEncoder(file, tc.contentType).Encode(tc.schema)
				}))
			}

			buf := bytes.NewBuffer(nil)
			enc := encoding.NewEncoder(buf, tc.contentType)
			assert.Equal(t, tc.contentType, enc.ContentType())
			assert.NoError(t, enc.Encode(tc.schema))
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			assert.Equal(t, expected, buf.Bytes())
		})
	}

	t.Run("unsupported content type", func(t *testing.T) {
		assert.Panics(t, func() { encoding.NewEncoder(bytes.NewBuffer(nil), "application/toml") })
	})
	t.Run("unsupported value by HTML encoder", func(t *testing.T) {
		enc := encoding.NewEncoder(bytes.NewBuffer(nil), encoding.HTML)
		assert.Error(t, enc.Encode("the value does not have `MarshalHTML` method"))
	})
	t.Run("writer fails", func(t *testing.T) {
		var enc encoding.Generic
		enc = encoding.NewEncoder(
			writerFn(func(p []byte) (n int, err error) { return 0, errors.New("problem writerFn") }), encoding.HTML)
		assert.Error(t, enc.Encode(domain.Schema{}))
		enc = encoding.NewEncoder(
			writerFn(func(p []byte) (n int, err error) { return 0, errors.New("problem writerFn") }), encoding.TEXT)
		assert.Error(t, enc.Encode(domain.Schema{}))
	})
}

func closeAfter(file *os.File, action func() error) error {
	defer file.Close()
	if err := action(); err != nil {
		return err
	}
	return nil
}

func writer(file string) *os.File {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	return f
}

type writerFn func(p []byte) (n int, err error)

func (fn writerFn) Write(p []byte) (n int, err error) {
	return fn(p)
}
