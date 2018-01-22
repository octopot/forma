package domen_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kamilsk/form-api/domen"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

const UUID domen.UUID = "41ca5e09-3ce2-4094-b108-3ecc257c6fa4"

var update = flag.Bool("update", false, "update .golden files")

func TestHTML(t *testing.T) {
	for _, tc := range []struct {
		name   string
		golden string
		schema domen.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.html.golden", domen.Schema{
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
			if *update {
				file := writer(tc.golden)
				assert.NoError(t, dryClose(file, func() error {
					html, err := tc.schema.MarshalHTML()
					if err != nil {
						return err
					}
					_, err = file.Write(html)
					return err
				}, false))
			}
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			obtained, err := tc.schema.MarshalHTML()
			assert.NoError(t, err)
			assert.Equal(t, golden, obtained)
		})
	}
}

func TestJSON(t *testing.T) {
	for _, tc := range []struct {
		name   string
		schema domen.Schema
	}{
		{"email subscription", domen.Schema{
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
			assert.Equal(t, tc.schema, func() domen.Schema {
				var schema domen.Schema
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
		schema   domen.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.json", domen.Schema{
			ID:           UUID.String() + "",
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
			var schema domen.Schema
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
		schema domen.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.json.golden", domen.Schema{
			ID:           UUID.String() + "",
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
			if *update {
				file := writer(tc.golden)
				assert.NoError(t, dryClose(file, func() error { return json.NewEncoder(file).Encode(tc.schema) }, false))
			}
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			buf := bytes.NewBuffer(nil)
			json.NewEncoder(buf).Encode(tc.schema)
			assert.Equal(t, string(golden), string(buf.Bytes()))
		})
	}
}

func TestXML(t *testing.T) {
	for _, tc := range []struct {
		name   string
		schema domen.Schema
	}{
		{"email subscription", domen.Schema{
			ID:           UUID.String() + "",
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
			assert.Equal(t, tc.schema, func() domen.Schema {
				var schema domen.Schema
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
		schema   domen.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.xml", domen.Schema{
			ID:           UUID.String() + "",
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
			var schema domen.Schema
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
		schema domen.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.xml.golden", domen.Schema{
			ID:           UUID.String() + "",
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
		{"stored in db", "./fixtures/stored_in_db.xml.golden", domen.Schema{
			Title:  "Email subscription",
			Action: "http://localhost:8080/api/v1/" + UUID.String(),
			Inputs: []domen.Input{
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
					return nil
				}, false))
			}
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			buf := bytes.NewBuffer(nil)
			xml.NewEncoder(buf).Encode(tc.schema)
			assert.Equal(t, string(golden), string(buf.Bytes()))
		})
	}
}

func TestYAML(t *testing.T) {
	for _, tc := range []struct {
		name   string
		schema domen.Schema
	}{
		{"email subscription", domen.Schema{
			ID:           UUID.String() + "",
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
			assert.Equal(t, tc.schema, func() domen.Schema {
				var schema domen.Schema
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
		schema   domen.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.yaml", domen.Schema{
			ID:           UUID.String() + "",
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
			var schema domen.Schema
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
		schema domen.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.yaml.golden", domen.Schema{
			ID:           UUID.String() + "",
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
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			data, err := yaml.Marshal(tc.schema)
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