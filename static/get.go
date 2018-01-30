package static

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/kamilsk/form-api/domain"
)

// ErrorPageContext contains data for `error.html` template.
type ErrorPageContext struct {
	Schema   domain.Schema
	Error    domain.AccumulatedError
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
	data, err := ioutil.ReadFile(filepath.Join(base, tpl))
	if err != nil && os.IsNotExist(err) {
		return Asset("static/templates/" + tpl)
	}
	return data, err
}
