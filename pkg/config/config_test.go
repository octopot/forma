package config_test

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var update = flag.Bool("update", false, "update .golden files")

func TestYAMLSerialization(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  string
	}{
		{"simple configuration", "fixtures/simple.yml", "fixtures/simple.yml.golden"},
	}

	for _, test := range testCases {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			raw, err := ioutil.ReadFile(tc.in)
			assert.NoError(t, err)

			var cnf config.ApplicationConfig
			err = yaml.UnmarshalStrict(raw, &cnf)
			assert.NoError(t, err)

			actual, err := yaml.Marshal(cnf)
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
