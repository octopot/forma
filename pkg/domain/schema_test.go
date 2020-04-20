package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/form-api/pkg/domain"
)

func TestSchema_Apply(t *testing.T) {
	tests := []struct {
		name     string
		schema   Schema
		values   map[string][]string
		expected struct {
			error bool
			panic bool
			data  map[string][]string
		}
	}{
		{"normal case", Schema{Inputs: []Input{
			{Name: "name1", Type: EmailType, MinLength: 6, MaxLength: 255, Required: true}}},
			map[string][]string{"name1": {"test@my.email"}, "not_filtered": {"val2"}},
			struct {
				error bool
				panic bool
				data  map[string][]string
			}{false, false, map[string][]string{"name1": {"test@my.email"}}},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			var (
				obtained map[string][]string
				err      ValidationError
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
	tests := []struct {
		name     string
		schema   Schema
		values   map[string][]string
		expected map[string][]string
	}{
		{"nil inputs", Schema{},
			map[string][]string{"name1": {"val1", "val2"}},
			nil,
		},
		{"empty inputs", Schema{Inputs: []Input{}},
			map[string][]string{"name1": {"val1", "val2"}},
			nil,
		},
		{"nil values", Schema{},
			nil,
			nil,
		},
		{"empty values", Schema{},
			nil,
			nil,
		},
		{"normal case", Schema{Inputs: []Input{{Name: "name1"}}},
			map[string][]string{"name1": {"val1"}, "name2": {"val2"}},
			map[string][]string{"name1": {"val1"}},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.schema.Filter(tc.values))
		})
	}
}

func TestSchema_Normalize(t *testing.T) {
	tests := []struct {
		name     string
		schema   Schema
		values   map[string][]string
		expected map[string][]string
	}{
		{"input with spaces", Schema{},
			map[string][]string{"name1": {string([]rune{'\u200B', '\u200C'}) + " val1 " + string([]rune{'\u200D', '\u2060'})}},
			map[string][]string{"name1": {"val1"}},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.schema.Normalize(tc.values))
		})
	}
}

func TestSchema_Validate(t *testing.T) {
	tests := []struct {
		name     string
		schema   Schema
		values   map[string][]string
		expected struct {
			error bool
			panic bool
			data  map[Input][]string
		}
	}{
		{"nil inputs", Schema{},
			map[string][]string{"name1": {"val1", "val2"}},
			struct {
				error bool
				panic bool
				data  map[Input][]string
			}{false, false, nil},
		},
		{"empty inputs", Schema{Inputs: []Input{}},
			map[string][]string{"name1": {"val1", "val2"}},
			struct {
				error bool
				panic bool
				data  map[Input][]string
			}{false, false, nil},
		},
		{"nil values", Schema{},
			nil,
			struct {
				error bool
				panic bool
				data  map[Input][]string
			}{false, false, nil},
		},
		{"empty values", Schema{},
			nil,
			struct {
				error bool
				panic bool
				data  map[Input][]string
			}{false, false, nil},
		},
		{"invalid length", Schema{Inputs: []Input{
			{Name: "name1", Type: TextType, MinLength: 5},
			{Name: "name2", Type: TextType, MaxLength: 2}}},
			map[string][]string{"name1": {"val1"}, "name2": {"val2"}},
			struct {
				error bool
				panic bool
				data  map[Input][]string
			}{true, false, map[Input][]string{
				{Name: "name1", Type: TextType, MinLength: 5}: {
					`value "val1" at position 0 is invalid: value length is less than 5`},
				{Name: "name2", Type: TextType, MaxLength: 2}: {
					`value "val2" at position 0 is invalid: value length is greater than 2`},
			}},
		},
		{"empty required value", Schema{Inputs: []Input{
			{Name: "name1", Type: TextType, Required: true},
			{Name: "name2", Type: TextType, Required: true}}},
			map[string][]string{"name1": {"val1", ""}, "name2": {}},
			struct {
				error bool
				panic bool
				data  map[Input][]string
			}{true, false, map[Input][]string{
				{Name: "name1", Type: TextType, Required: true}: {"value at position 1 is invalid: value is empty"},
				{Name: "name2", Type: TextType, Required: true}: {"values are empty"},
			}},
		},
		{"invalid type", Schema{Inputs: []Input{{Name: "name1", Type: "unknown"}}},
			map[string][]string{"name1": {"val1", "val2"}},
			struct {
				error bool
				panic bool
				data  map[Input][]string
			}{false, true, map[Input][]string{}},
		},
		{"invalid email", Schema{Inputs: []Input{{Name: "name1", Type: EmailType}}},
			map[string][]string{"name1": {"val1", "val2"}},
			struct {
				error bool
				panic bool
				data  map[Input][]string
			}{true, false, map[Input][]string{
				{Name: "name1", Type: EmailType}: {
					`value "val1" at position 0 is invalid: value is not a valid email`},
			}},
		},
		{"normal case", Schema{Inputs: []Input{
			{Name: "name1", Type: EmailType, MinLength: 6, MaxLength: 255, Required: true}}},
			map[string][]string{"name1": {"test@my.email"}, "not_filtered": {"val2"}},
			struct {
				error bool
				panic bool
				data  map[Input][]string
			}{},
		},
	}

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			var err ValidationError
			action := func() { _, err = tc.schema.Validate(tc.values) }
			if tc.expected.panic {
				assert.Panics(t, action)
			} else {
				assert.NotPanics(t, action)
			}
			if tc.expected.error {
				assert.EqualError(t, err, "validation error")
				obtained := err.InputWithErrors()
				for input, errors := range obtained {
					expected := tc.expected.data[input]
					assert.True(t, err.HasError(input))
					assert.Equal(t, expected, func() []string {
						converted := make([]string, 0, len(errors))
						for _, e := range errors {
							converted = append(converted, e.Error())
						}
						return converted
					}())
				}
				assert.False(t, err.HasError(Input{}))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSchema_Input(t *testing.T) {
	tests := []struct {
		name     string
		schema   Schema
		what     string
		expected int
	}{
		{
			name:     "expect the first occurrence",
			schema:   Schema{Inputs: []Input{{Name: "First"}, {Name: "fiRst"}, {Name: "firSt"}}},
			what:     "first",
			expected: 0,
		},
		{
			name:     "expect not find anything",
			schema:   Schema{Inputs: []Input{{Name: "First"}, {Name: "fiRst"}, {Name: "firSt"}}},
			what:     "second",
			expected: -1,
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			input := test.schema.Input(test.what)
			if test.expected == -1 {
				assert.Nil(t, input, test.name)
			} else {
				assert.Equal(t, input, &test.schema.Inputs[test.expected], test.name)
			}
		})
	}
}
