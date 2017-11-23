package static

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// LoadTemplate loads template from custom location or fallback it to `go-bindata`.
func LoadTemplate(base, tpl string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath.Join(base, tpl))
	if err != nil && os.IsNotExist(err) {
		return Asset(tpl)
	}
	return data, err
}
