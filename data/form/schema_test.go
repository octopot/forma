package form_test

import (
	"fmt"
	"testing"

	"github.com/kamilsk/form-api/data/form"
	"github.com/stretchr/testify/assert"
)

func TestSchema_Apply(t *testing.T) {
	for _, tc := range []struct {
		name     string
		schema   form.Schema
		values   map[string][]string
		expected map[string][]string
	}{
		{
			name:     "nil inputs",
			schema:   form.Schema{},
			values:   map[string][]string{"name1": {"val1", "val2"}},
			expected: nil,
		},
		{
			name:     "empty inputs",
			schema:   form.Schema{Inputs: []form.Input{}},
			values:   map[string][]string{"name1": {"val1", "val2"}},
			expected: nil,
		},
		{
			name:     "nil values",
			schema:   form.Schema{},
			values:   nil,
			expected: nil,
		},
		{
			name:     "empty values",
			schema:   form.Schema{},
			values:   nil,
			expected: nil,
		},
		{
			name:     "normal case",
			schema:   form.Schema{Inputs: []form.Input{{Name: "name1"}}},
			values:   map[string][]string{"name1": {"val1"}, "name2": {"val2"}},
			expected: map[string][]string{"name1": {"val1"}},
		},
	} {
		filtered := tc.schema.Apply(tc.values)
		assert.Equal(t, tc.expected, filtered, fmt.Sprintf("test case %q failed", tc.name))
	}
}
