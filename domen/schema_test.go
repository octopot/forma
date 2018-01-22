package domen_test

import (
	"testing"

	"github.com/kamilsk/form-api/domen"
	"github.com/stretchr/testify/assert"
)

func TestSchema_Apply(t *testing.T) {
	for _, tc := range []struct {
		name     string
		schema   domen.Schema
		values   map[string][]string
		expected struct {
			error bool
			panic bool
			data  map[string][]string
		}
	}{
		{"normal case", domen.Schema{Inputs: []domen.Input{
			{Name: "name1", Type: domen.EmailType, MinLength: 6, MaxLength: 255, Required: true}}},
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
				err      domen.ValidationError
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
		schema   domen.Schema
		values   map[string][]string
		expected map[string][]string
	}{
		{"nil inputs", domen.Schema{},
			map[string][]string{"name1": {"val1", "val2"}},
			nil,
		},
		{"empty inputs", domen.Schema{Inputs: []domen.Input{}},
			map[string][]string{"name1": {"val1", "val2"}},
			nil,
		},
		{"nil values", domen.Schema{},
			nil,
			nil,
		},
		{"empty values", domen.Schema{},
			nil,
			nil,
		},
		{"normal case", domen.Schema{Inputs: []domen.Input{{Name: "name1"}}},
			map[string][]string{"name1": {"val1"}, "name2": {"val2"}},
			map[string][]string{"name1": {"val1"}},
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
		schema   domen.Schema
		values   map[string][]string
		expected map[string][]string
	}{
		{"input with spaces", domen.Schema{},
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
		schema   domen.Schema
		values   map[string][]string
		expected struct {
			error bool
			panic bool
			data  map[domen.Input][]string
		}
	}{
		{"nil inputs", domen.Schema{},
			map[string][]string{"name1": {"val1", "val2"}},
			struct {
				error bool
				panic bool
				data  map[domen.Input][]string
			}{false, false, nil},
		},
		{"empty inputs", domen.Schema{Inputs: []domen.Input{}},
			map[string][]string{"name1": {"val1", "val2"}},
			struct {
				error bool
				panic bool
				data  map[domen.Input][]string
			}{false, false, nil},
		},
		{"nil values", domen.Schema{},
			nil,
			struct {
				error bool
				panic bool
				data  map[domen.Input][]string
			}{false, false, nil},
		},
		{"empty values", domen.Schema{},
			nil,
			struct {
				error bool
				panic bool
				data  map[domen.Input][]string
			}{false, false, nil},
		},
		{"invalid length", domen.Schema{Inputs: []domen.Input{
			{Name: "name1", Type: domen.TextType, MinLength: 5},
			{Name: "name2", Type: domen.TextType, MaxLength: 2}}},
			map[string][]string{"name1": {"val1"}, "name2": {"val2"}},
			struct {
				error bool
				panic bool
				data  map[domen.Input][]string
			}{true, false, map[domen.Input][]string{
				{Name: "name1", Type: domen.TextType, MinLength: 5}: {
					`value "val1" at position 0 is invalid: value length is less than 5`},
				{Name: "name2", Type: domen.TextType, MaxLength: 2}: {
					`value "val2" at position 0 is invalid: value length is greater than 2`},
			}},
		},
		{"empty required value", domen.Schema{Inputs: []domen.Input{
			{Name: "name1", Type: domen.TextType, Required: true},
			{Name: "name2", Type: domen.TextType, Required: true}}},
			map[string][]string{"name1": {"val1", ""}, "name2": {}},
			struct {
				error bool
				panic bool
				data  map[domen.Input][]string
			}{true, false, map[domen.Input][]string{
				{Name: "name1", Type: domen.TextType, Required: true}: {"value at position 1 is invalid: value is empty"},
				{Name: "name2", Type: domen.TextType, Required: true}: {"values are empty"},
			}},
		},
		{"invalid type", domen.Schema{Inputs: []domen.Input{{Name: "name1", Type: "unknown"}}},
			map[string][]string{"name1": {"val1", "val2"}},
			struct {
				error bool
				panic bool
				data  map[domen.Input][]string
			}{false, true, map[domen.Input][]string{}},
		},
		{"invalid email", domen.Schema{Inputs: []domen.Input{{Name: "name1", Type: domen.EmailType}}},
			map[string][]string{"name1": {"val1", "val2"}},
			struct {
				error bool
				panic bool
				data  map[domen.Input][]string
			}{true, false, map[domen.Input][]string{
				{Name: "name1", Type: domen.EmailType}: {
					`value "val1" at position 0 is invalid: value is not a valid email`},
			}},
		},
		{"normal case", domen.Schema{Inputs: []domen.Input{
			{Name: "name1", Type: domen.EmailType, MinLength: 6, MaxLength: 255, Required: true}}},
			map[string][]string{"name1": {"test@my.email"}, "not_filtered": {"val2"}},
			struct {
				error bool
				panic bool
				data  map[domen.Input][]string
			}{},
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var err domen.ValidationError
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
