package domain_test

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var update = flag.Bool("update", false, "update .golden files")

func TestHTML(t *testing.T) {
	tests := []struct {
		name   string
		golden string
		schema domain.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.html.golden", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
					html, err := tc.schema.MarshalHTML()
					if err != nil {
						return err
					}
					_, err = file.Write(html)
					return err
				}))
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
	tests := []struct {
		name   string
		schema domain.Schema
	}{
		{"email subscription", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
			assert.Equal(t, tc.schema, func() domain.Schema {
				var schema domain.Schema
				data, err := json.MarshalIndent(tc.schema, "", "  ")
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
	tests := []struct {
		name     string
		filename string
		schema   domain.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.json", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
			var schema domain.Schema
			file := reader(tc.filename)
			assert.NoError(t, closeAfter(file, func() error { return json.NewDecoder(file).Decode(&schema) }))
			assert.Equal(t, tc.schema, schema)
		})
	}
}

func TestJSON_Encode(t *testing.T) {
	tests := []struct {
		name   string
		golden string
		schema domain.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.json.golden", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
					data, err := json.MarshalIndent(tc.schema, "", "  ")
					if err != nil {
						return err
					}
					_, err = file.Write(data)
					return err
				}))
			}
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			data, err := json.MarshalIndent(tc.schema, "", "  ")
			assert.NoError(t, err)
			assert.Equal(t, golden, data)
		})
	}
}

func TestXML(t *testing.T) {
	tests := []struct {
		name   string
		schema domain.Schema
	}{
		{"email subscription", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
			assert.Equal(t, tc.schema, func() domain.Schema {
				var schema domain.Schema
				data, err := xml.MarshalIndent(tc.schema, "", "  ")
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
	tests := []struct {
		name     string
		filename string
		schema   domain.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.xml", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
			var schema domain.Schema
			file := reader(tc.filename)
			assert.NoError(t, closeAfter(file, func() error { return xml.NewDecoder(file).Decode(&schema) }))
			assert.Equal(t, tc.schema, schema)
		})
	}
}

func TestXML_Encode(t *testing.T) {
	tests := []struct {
		name   string
		golden string
		schema domain.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.xml.golden", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
					Title:     "Email",
					MaxLength: 64,
					Required:  true,
				},
			},
		}},
		{"stored in db", "./fixtures/stored_in_db.xml.golden", domain.Schema{
			Language: "en",
			Title:    "Email subscription",
			Action:   "https://kamil.samigullin.info/",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
					data, err := xml.MarshalIndent(tc.schema, "", "    ")
					if err != nil {
						return err
					}
					_, err = file.Write(data)
					return err
				}))
			}
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			data, err := xml.MarshalIndent(tc.schema, "", "    ")
			assert.NoError(t, err)
			assert.Equal(t, golden, data)
		})
	}
}

func TestYAML(t *testing.T) {
	tests := []struct {
		name   string
		schema domain.Schema
	}{
		{"email subscription", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
			assert.Equal(t, tc.schema, func() domain.Schema {
				var schema domain.Schema
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
	tests := []struct {
		name     string
		filename string
		schema   domain.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.yaml", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
			var schema domain.Schema
			file := reader(tc.filename)
			assert.NoError(t, closeAfter(file, func() error {
				return yaml.Unmarshal(func() []byte {
					data, err := ioutil.ReadAll(file)
					if err != nil {
						panic(err)
					}
					return data
				}(), &schema)
			}))
			assert.Equal(t, tc.schema, schema)
		})
	}
}

func TestYAML_Encode(t *testing.T) {
	tests := []struct {
		name   string
		golden string
		schema domain.Schema
	}{
		{"email subscription", "./fixtures/email_subscription.yaml.golden", domain.Schema{
			Language:     "en",
			Title:        "Email subscription",
			Action:       "https://kamil.samigullin.info/",
			Method:       "post",
			EncodingType: "application/x-www-form-urlencoded",
			Inputs: []domain.Input{
				{
					Name:      "email",
					Type:      domain.EmailType,
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
					data, err := yaml.Marshal(tc.schema)
					if err != nil {
						return err
					}
					_, err = file.Write(data)
					return err
				}))
			}
			golden, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			data, err := yaml.Marshal(tc.schema)
			assert.NoError(t, err)
			assert.Equal(t, string(golden), string(data))
		})
	}
}

func closeAfter(file *os.File, action func() error) error {
	defer file.Close()
	if err := action(); err != nil {
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
