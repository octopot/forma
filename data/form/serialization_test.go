package form_test

import (
	"bytes"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kamilsk/form-api/data/form"
	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update .golden files")

func TestXML_Decode(t *testing.T) {
	for _, tc := range []struct {
		filename string
		schema   form.Schema
	}{
		{filename: "./fixtures/email_subscription.xml", schema: form.Schema{
			ID:    "",
			Title: "Email subscription",
			Inputs: []form.Input{
				{
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					Name:  "redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		var schema form.Schema
		file := reader(tc.filename)
		assert.NoError(t, dryClose(file, func() error { return xml.NewDecoder(file).Decode(&schema) }, false))
		assert.Equal(t, tc.schema, schema)
	}
}

func TestXML_Encode(t *testing.T) {
	for _, tc := range []struct {
		golden string
		schema form.Schema
	}{
		{golden: "./fixtures/email_subscription.golden", schema: form.Schema{
			ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title: "Email subscription",
			Inputs: []form.Input{
				{
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					Name:  "redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		if *update {
			file := writer(tc.golden)
			assert.NoError(t, dryClose(file, func() error { return tc.schema.MarshalTo(file) }, false))
		}
		buf := bytes.NewBuffer(nil)
		tc.schema.MarshalTo(buf)
		golden, err := ioutil.ReadFile(tc.golden)
		assert.NoError(t, err)
		assert.Equal(t, golden, buf.Bytes())
	}
}

func dryClose(file *os.File, action func() error, closeIfError bool) error {
	if !closeIfError {
		defer file.Close()
	}
	if err := action(); err != nil {
		if closeIfError {
			file.Close()
		}
		return err
	}
	return nil
}

func reader(file string) *os.File {
	f, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	return f
}

func writer(file string) *os.File {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	return f
}
