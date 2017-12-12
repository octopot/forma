package static_test

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"testing"

	"github.com/kamilsk/form-api/static"
	"github.com/stretchr/testify/assert"
)

func TestAsset(t *testing.T) {
	for _, tc := range []struct {
		name   string
		asset  string
		golden string
	}{
		{"init migration", "static/migrations/1_initial.sql", "./migrations/1_initial.sql"},
		{"error template", "static/templates/error.html", "./templates/error.html"},
		{"redirect template", "static/templates/redirect.html", "./templates/redirect.html"},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			obtained, err := static.Asset(filepath.Join(tc.asset))
			assert.NoError(t, err)
			assert.Equal(t, string(expected), string(obtained))
		})
	}
}

func TestAssetDir(t *testing.T) {
	for _, tc := range []struct {
		name     string
		assetDir string
		expected struct {
			files []string
		}
	}{
		{"root", "static", struct {
			files []string
		}{[]string{"migrations", "templates"}}},
		{"migrations", "static/migrations", struct {
			files []string
		}{[]string{"1_initial.sql", "demo"}}},
		{"templates", "static/templates", struct {
			files []string
		}{[]string{"error.html", "redirect.html"}}},
		{"not found", "static/templates/unknown", struct {
			files []string
		}{}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			files, err := static.AssetDir(tc.assetDir)
			sort.Strings(tc.expected.files)
			sort.Strings(files)
			assert.Equal(t, tc.expected.files, files)
			if len(files) == 0 {
				assert.Error(t, err)
			}
		})

	}
}

func TestAssetInfo(t *testing.T) {
	// TODO v2: up code coverage
}

func TestAssetNames(t *testing.T) {
	// TODO v2: up code coverage
}

func TestMustAsset(t *testing.T) {
	// TODO v2: up code coverage
}

func TestRestoreAsset(t *testing.T) {
	// TODO v2: up code coverage
}

func TestRestoreAssets(t *testing.T) {
	// TODO v2: up code coverage
}
