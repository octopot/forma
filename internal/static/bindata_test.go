package static_test

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/ecosystem/forma/internal/static"
)

func TestAsset(t *testing.T) {
	tests := []struct {
		name   string
		asset  string
		golden string
	}{
		{"prepare migration", "static/migrations/1_prepare.sql", "./migrations/1_prepare.sql"},
		{"account migration", "static/migrations/2_account.sql", "./migrations/2_account.sql"},
		{"domain migration", "static/migrations/3_domain.sql", "./migrations/3_domain.sql"},
		{"audit migration", "static/migrations/4_audit.sql", "./migrations/4_audit.sql"},
		{"error template", "static/templates/error.html", "./templates/error.html"},
		{"redirect template", "static/templates/redirect.html", "./templates/redirect.html"},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			expected, err := ioutil.ReadFile(tc.golden)
			assert.NoError(t, err)
			obtained, err := Asset(filepath.Join(tc.asset))
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
		{"migrations", "static/migrations", []string{
			"1_prepare.sql",
			"2_account.sql",
			"3_domain.sql",
			"4_audit.sql",
		}},
		{"templates", "static/templates", []string{"error.html", "redirect.html"}},
		{"not found", "static/templates/unknown", nil},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			files, err := AssetDir(tc.assetDir)
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
