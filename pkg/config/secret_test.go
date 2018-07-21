package config_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"testing"

	"github.com/kamilsk/form-api/pkg/config"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSecret_Printing(t *testing.T) {
	type password struct {
		XMLName struct{}      `json:"-" xml:"password" yaml:"-"`
		Value   config.Secret `json:"password" xml:"value,attr" yaml:"password"`
	}
	secret := password{Value: "secret"}

	testCases := []struct {
		name    string
		marshal func(password) ([]byte, error)
	}{
		{"print by `%#v`", func(pass password) ([]byte, error) {
			str := fmt.Sprintf("%#v", pass)
			return []byte(str), nil
		}},
		{"print by `%s`", func(pass password) ([]byte, error) {
			str := fmt.Sprintf("%s", pass.Value)
			return []byte(str), nil
		}},
		{"json marshal", func(pass password) ([]byte, error) {
			return json.Marshal(pass)
		}},
		{"xml marshal", func(pass password) ([]byte, error) {
			return xml.Marshal(pass)
		}},
		{"yaml marshal", func(pass password) ([]byte, error) {
			return yaml.Marshal(pass)
		}},
	}

	for _, test := range testCases {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			raw, err := tc.marshal(secret)
			assert.NoError(t, err)
			assert.False(t, strings.Contains(string(raw), string(secret.Value)))
		})
	}
}
