package form_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kamilsk/form-api/data/form"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var update = flag.Bool("update", false, "update .golden files")

func TestJSON(t *testing.T) {
	for _, tc := range []struct {
		name   string
		schema form.Schema
	}{
		{"email subscription", form.Schema{
			ID:      "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.schema, func() form.Schema {
				var schema form.Schema
				data, err := json.Marshal(tc.schema)
				if err != nil {
					panic(err)
				}
				if err := json.Unmarshal(data, &schema); err != nil {
					panic(err)
				}
				return schema
			}())
		})
	}
}

func TestJSON_Decode(t *testing.T) {
	for _, tc := range []struct {
		name     string
		filename string
		schema   form.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.json", form.Schema{
			ID:      "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var schema form.Schema
			file := reader(tc.filename)
			assert.NoError(t, dryClose(file, func() error { return json.NewDecoder(file).Decode(&schema) }, false))
			assert.Equal(t, tc.schema, schema)
		})
	}
}

func TestJSON_Encode(t *testing.T) {
	for _, tc := range []struct {
		name   string
		golden string
		schema form.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.json.golden", form.Schema{
			ID:      "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if *update {
				file := writer(tc.golden)
				assert.NoError(t, dryClose(file, func() error { return json.NewEncoder(file).Encode(tc.schema) }, false))
			}
			buf := bytes.NewBuffer(nil)
			json.NewEncoder(buf).Encode(tc.schema)
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			assert.Equal(t, string(golden), string(buf.Bytes()))
		})
	}
}

func TestXML(t *testing.T) {
	for _, tc := range []struct {
		name   string
		schema form.Schema
	}{
		{"email subscription", form.Schema{
			ID:      "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.schema, func() form.Schema {
				var schema form.Schema
				data, err := xml.Marshal(tc.schema)
				if err != nil {
					panic(err)
				}
				if err := xml.Unmarshal(data, &schema); err != nil {
					panic(err)
				}
				return schema
			}())
		})
	}
}

func TestXML_Decode(t *testing.T) {
	for _, tc := range []struct {
		name     string
		filename string
		schema   form.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.xml", form.Schema{
			ID:      "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var schema form.Schema
			file := reader(tc.filename)
			assert.NoError(t, dryClose(file, func() error { return xml.NewDecoder(file).Decode(&schema) }, false))
			assert.Equal(t, tc.schema, schema)
		})
	}
}

func TestXML_Encode(t *testing.T) {
	for _, tc := range []struct {
		name   string
		golden string
		schema form.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.xml.golden", form.Schema{
			ID:      "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
		{"stored in db", "./fixtures/stored_in_db.xml.golden", form.Schema{
			Title:  "Email subscription",
			Action: "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Inputs: []form.Input{
				{
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if *update {
				file := writer(tc.golden)
				assert.NoError(t, dryClose(file, func() error {
					if err := xml.NewEncoder(file).Encode(tc.schema); err != nil {
						return err
					}
					_, err := file.Write([]byte("\n"))
					return err
				}, false))
			}
			buf := bytes.NewBuffer(nil)
			xml.NewEncoder(buf).Encode(tc.schema)
			buf.Write([]byte("\n"))
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			assert.Equal(t, string(golden), string(buf.Bytes()))
		})
	}
}

func TestYAML(t *testing.T) {
	for _, tc := range []struct {
		name   string
		schema form.Schema
	}{
		{"email subscription", form.Schema{
			ID:      "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.schema, func() form.Schema {
				var schema form.Schema
				data, err := yaml.Marshal(tc.schema)
				if err != nil {
					panic(err)
				}
				if err := yaml.Unmarshal(data, &schema); err != nil {
					panic(err)
				}
				return schema
			}())
		})
	}
}

func TestYAML_Decode(t *testing.T) {
	for _, tc := range []struct {
		name     string
		filename string
		schema   form.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.yaml", form.Schema{
			ID:      "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var schema form.Schema
			file := reader(tc.filename)
			assert.NoError(t, dryClose(file, func() error {
				return yaml.Unmarshal(func() []byte {
					data, err := ioutil.ReadAll(file)
					if err != nil {
						panic(err)
					}
					return data
				}(), &schema)
			}, false))
			assert.Equal(t, tc.schema, schema)
		})
	}
}

func TestYAML_Encode(t *testing.T) {
	for _, tc := range []struct {
		name   string
		golden string
		schema form.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.yaml.golden", form.Schema{
			ID:      "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Title:   "Email subscription",
			Action:  "http://localhost:8080/api/v1/a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11",
			Method:  "post",
			EncType: "application/x-www-form-urlencoded",
			Inputs: []form.Input{
				{
					ID:        "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11_email",
					Name:      "email",
					Type:      "email",
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
				{
					ID:    "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11__redirect",
					Name:  "_redirect",
					Type:  "hidden",
					Value: "https://kamil.samigullin.info/",
				},
			},
		}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if *update {
				file := writer(tc.golden)
				assert.NoError(t, dryClose(file, func() error {
					data, err := yaml.Marshal(tc.schema)
					if err != nil {
						return err
					}
					_, err = file.Write(data)
					return err
				}, false))
			}
			data, err := yaml.Marshal(tc.schema)
			assert.NoError(t, err)
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			assert.Equal(t, string(golden), string(data))
		})
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
