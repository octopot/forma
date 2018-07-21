package config_test

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var update = flag.Bool("update", false, "update .golden files")

func TestApplicationConfig_Dumping(t *testing.T) {
	testCases := []struct {
		name    string
		in      string
		out     string
		marshal func(interface{}) ([]byte, error)
	}{
		{"YAML dump", "fixtures/config.yml", "fixtures/dump.yml.golden", yaml.Marshal},
		{"JSON dump", "fixtures/config.yml", "fixtures/dump.json.golden", json.Marshal},
	}

	for _, test := range testCases {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			raw, err := ioutil.ReadFile(tc.in)
			assert.NoError(t, err)

			var cnf config.ApplicationConfig
			err = yaml.UnmarshalStrict(raw, &cnf)
			assert.NoError(t, err)

			actual, err := tc.marshal(cnf)
			assert.NoError(t, err)

			if *update {
				err = ioutil.WriteFile(tc.out, actual, os.ModePerm)
				assert.NoError(t, err)
			}
			expected, err := ioutil.ReadFile(tc.out)
			assert.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	}
}
