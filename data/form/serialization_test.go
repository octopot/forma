package form_test

import (
	"encoding/xml"
	"os"
	"reflect"
	"testing"

	"github.com/kamilsk/form-api/data/form"
)

func TestXML_Decode(t *testing.T) {
	for _, tc := range []struct {
		filename string
		schema   form.Schema
	}{
		{filename: "./fixtures/email_subscription.xml", schema: form.Schema{
			ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title: "Email subscription",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
				},
				{
					ID:   "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_redirect",
					Name: "redirect",
					Type: "hidden",
				},
			},
		}},
	} {
		var schema form.Schema
		err := xml.NewDecoder(func() *os.File {
			f, err := os.Open(tc.filename)
			if err != nil {
				panic(err)
			}
			return f
		}()).Decode(&schema)
		if err != nil {
			t.Error("unexpected error", err)
		}
		if !reflect.DeepEqual(schema, tc.schema) {
			t.Errorf(`
expected form schema:
%+v
obtained:
%+v`, tc.schema, schema)
		}
	}
}
