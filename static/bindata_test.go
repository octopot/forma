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
	tests := []struct {
		name   string
		asset  string
		golden string
	}{
		{"init migration", "static/migrations/1_initial.sql", "./migrations/1_initial.sql"},
		{"error template", "static/templates/error.html", "./templates/error.html"},
		{"redirect template", "static/templates/redirect.html", "./templates/redirect.html"},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			obtained, err := static.Asset(filepath.Join(tc.asset))
			assert.NoError(t, err)
			assert.Equal(t, expected, obtained)
		})
	}
}

func TestMustAsset(t *testing.T) {
}

func TestAssetInfo(t *testing.T) {
}

func TestAssetNames(t *testing.T) {
}

func TestAssetDir(t *testing.T) {
	tests := []struct {
		name     string
		assetDir string
		expected []string
	}{
		{"root", "static", []string{"migrations", "templates"}},
		{"migrations", "static/migrations", []string{"1_initial.sql"}},
		{"templates", "static/templates", []string{"error.html", "redirect.html"}},
		{"not found", "static/templates/unknown", nil},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			files, err := static.AssetDir(tc.assetDir)
			sort.Strings(tc.expected)
			sort.Strings(files)
			assert.Equal(t, tc.expected, files)
			if len(files) == 0 {
				assert.Error(t, err)
			}
		})
	}
}

func TestRestoreAsset(t *testing.T) {
}

func TestRestoreAssets(t *testing.T) {
}
