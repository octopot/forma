package static

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadTemplate loads the template from a custom location or fallback it to `bindata`.
// TODO v2: support by flags
func LoadTemplate(base, tpl string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath.Join(base, tpl))
	if err != nil && os.IsNotExist(err) {
		return Asset("static/templates/" + tpl)
	}
	return data, err
}
