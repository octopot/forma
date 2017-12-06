package form_test

import (
	"testing"

	"github.com/kamilsk/form-api/data/form"
	"github.com/stretchr/testify/assert"
)

func TestSchema_Apply(t *testing.T) {
	for _, tc := range []struct {
		name     string
		schema   form.Schema
		values   map[string][]string
		expected struct {
			error bool
			panic bool
			data  map[string][]string
		}
	}{
		{"normal case", form.Schema{Inputs: []form.Input{
			{Name: "name1", Type: form.EmailType, MinLength: 6, MaxLength: 255, Required: true},
		}},
			map[string][]string{"name1": {"test@my.email"}, "not_filtered": {"val2"}},
			struct {
				error bool
				panic bool
				data  map[string][]string
			}{false, false, map[string][]string{"name1": {"test@my.email"}}},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var (
				obtained map[string][]string
				err      form.ValidationError
			)
			action := func() { obtained, err = tc.schema.Apply(tc.values) }
			if tc.expected.panic {
				assert.Panics(t, action)
			} else {
				assert.NotPanics(t, action)
			}
			if tc.expected.error {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected.data, obtained)
			}
		})
	}
}

func TestSchema_Filter(t *testing.T) {
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
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.schema.Filter(tc.values))
		})
	}
}

func TestSchema_Normalize(t *testing.T) {
	for _, tc := range []struct {
		name     string
		schema   form.Schema
		values   map[string][]string
		expected map[string][]string
	}{
		{"input with spaces", form.Schema{},
			map[string][]string{"name1": {string([]rune{'\u200B', '\u200C'}) + " val1 " + string([]rune{'\u200D', '\u2060'})}},
			map[string][]string{"name1": {"val1"}},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.schema.Normalize(tc.values))
		})
	}
}

func TestSchema_Validate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		schema   form.Schema
		values   map[string][]string
		expected struct {
			error bool
			panic bool
			data  map[form.Input][]string
		}
	}{
		{
			name:   "nil inputs",
			schema: form.Schema{},
			values: map[string][]string{"name1": {"val1", "val2"}},
			expected: struct {
				error bool
				panic bool
				data  map[form.Input][]string
			}{false, false, nil},
		},
		{
			name:   "empty inputs",
			schema: form.Schema{Inputs: []form.Input{}},
			values: map[string][]string{"name1": {"val1", "val2"}},
			expected: struct {
				error bool
				panic bool
				data  map[form.Input][]string
			}{false, false, nil},
		},
		{
			name:   "nil values",
			schema: form.Schema{},
			values: nil,
			expected: struct {
				error bool
				panic bool
				data  map[form.Input][]string
			}{false, false, nil},
		},
		{
			name:   "empty values",
			schema: form.Schema{},
			values: nil,
			expected: struct {
				error bool
				panic bool
				data  map[form.Input][]string
			}{false, false, nil},
		},
		{
			name: "invalid length",
			schema: form.Schema{Inputs: []form.Input{
				{Name: "name1", Type: form.TextType, MinLength: 5},
				{Name: "name2", Type: form.TextType, MaxLength: 2},
			}},
			values: map[string][]string{"name1": {"val1"}, "name2": {"val2"}},
			expected: struct {
				error bool
				panic bool
				data  map[form.Input][]string
			}{true, false, map[form.Input][]string{
				{Name: "name1", Type: form.TextType, MinLength: 5}: {
					`value "val1" at position 0 is invalid: value length is less than 5`},
				{Name: "name2", Type: form.TextType, MaxLength: 2}: {
					`value "val2" at position 0 is invalid: value length is greater than 2`},
			}},
		},
		{
			name: "empty required value",
			schema: form.Schema{Inputs: []form.Input{
				{Name: "name1", Type: form.TextType, Required: true},
				{Name: "name2", Type: form.TextType, Required: true},
			}},
			values: map[string][]string{"name1": {"val1", ""}, "name2": {}},
			expected: struct {
				error bool
				panic bool
				data  map[form.Input][]string
			}{true, false, map[form.Input][]string{
				{Name: "name1", Type: form.TextType, Required: true}: {"value at position 1 is invalid: value is empty"},
				{Name: "name2", Type: form.TextType, Required: true}: {"values are empty"},
			}},
		},
		{
			name:   "invalid type",
			schema: form.Schema{Inputs: []form.Input{{Name: "name1", Type: "unknown"}}},
			values: map[string][]string{"name1": {"val1", "val2"}},
			expected: struct {
				error bool
				panic bool
				data  map[form.Input][]string
			}{false, true, map[form.Input][]string{}},
		},
		{
			name:   "invalid email",
			schema: form.Schema{Inputs: []form.Input{{Name: "name1", Type: form.EmailType}}},
			values: map[string][]string{"name1": {"val1", "val2"}},
			expected: struct {
				error bool
				panic bool
				data  map[form.Input][]string
			}{true, false, map[form.Input][]string{
				{Name: "name1", Type: form.EmailType}: {
					`value "val1" at position 0 is invalid: value is not a valid email`},
			}},
		},
		{
			name: "normal case",
			schema: form.Schema{Inputs: []form.Input{
				{Name: "name1", Type: form.EmailType, MinLength: 6, MaxLength: 255, Required: true},
			}},
			values: map[string][]string{"name1": {"test@my.email"}, "not_filtered": {"val2"}},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var err form.ValidationError
			action := func() { _, err = tc.schema.Validate(tc.values) }
			if tc.expected.panic {
				assert.Panics(t, action)
			} else {
				assert.NotPanics(t, action)
			}
			if tc.expected.error {
				assert.EqualError(t, err, "input data has error")
				obtained := err.InputWithErrors()
				for input, errors := range obtained {
					expected := tc.expected.data[input]
					assert.Equal(t, expected, func() []string {
						converted := make([]string, 0, len(errors))
						for _, err := range errors {
							converted = append(converted, err.Error())
						}
						return converted
					}())
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
