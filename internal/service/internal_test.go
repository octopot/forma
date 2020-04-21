package service

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extend(t *testing.T) {
	tests := []struct {
		name     string
		url      url.URL
		paths    []string
		expected string
	}{
		{"without paths", url.URL{Path: "/"}, nil, "/"},
		{"with some paths", url.URL{Path: "/with"}, []string{"some", "paths"}, "/with/some/paths"},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, extend(tc.url, tc.paths...))
		})
	}
}

func Test_must(t *testing.T) {
	tests := []struct {
		name      string
		base, tpl string
		panicked  bool
	}{
		{"non-existent template", "/", "_.tpl", true},
	}

	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			if tc.panicked {
				assert.Panics(t, func() { must(tc.base, tc.tpl) })
			} else {
				assert.NotEmpty(t, must(tc.base, tc.tpl))
			}
		})
	}
}
