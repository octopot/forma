package static_test

import (
	"io/ioutil"
	"testing"

	"github.com/kamilsk/form-api/static"
	"github.com/stretchr/testify/assert"
)

func TestLoadTemplate(t *testing.T) {
	for _, tc := range []struct {
		name      string
		base, tpl string
		golden    string
	}{
		{"error template", "templates", "error.html", "./templates/error.html"},
		{"error template, bindata", "/", "error.html", "./templates/error.html"},
		{"redirect template", "templates", "redirect.html", "./templates/redirect.html"},
		{"redirect template, bindata", "/", "redirect.html", "./templates/redirect.html"},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			obtained, err := static.LoadTemplate(tc.base, tc.tpl)
			assert.NoError(t, err)
			assert.Equal(t, string(expected), string(obtained))
		})
	}
}
