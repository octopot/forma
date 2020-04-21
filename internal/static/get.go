package static

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"go.octolab.org/ecosystem/forma/internal/domain"
)

// ErrorPageContext contains data for `error.html` template.
type ErrorPageContext struct {
	Schema   domain.Schema
	Error    domain.ValidationError
	Delay    time.Duration
	Redirect string
}

// RedirectPageContext contains data for `redirect.html` template.
type RedirectPageContext struct {
	Schema   domain.Schema
	Delay    time.Duration
	Redirect string
}

// LoadTemplate loads the template from a custom location or fallback it to `bindata`.
func LoadTemplate(base, tpl string) ([]byte, error) {
	path := filepath.Join(base, tpl)
	data, err := ioutil.ReadFile(path)
	if err != nil && os.IsNotExist(err) {
		return Asset(path)
	}
	return data, err
}
