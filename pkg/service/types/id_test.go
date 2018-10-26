package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/form-api/pkg/service/types"
)

func TestID(t *testing.T) {
	type entity struct {
		ID
		valid bool
	}

	tests := []struct {
		name   string
		entity entity
	}{
		{"ID is empty", entity{"", false}},
		{"ID is invalid", entity{"abc-def-ghi", false}},
		{"ID is not UUID v4", entity{"41ca5e09-3ce2-3094-b108-3ecc257c6fa4", false}},
		{"ID in lowercase", entity{"41ca5e09-3ce2-4094-b108-3ecc257c6fa4", true}},
		{"ID in uppercase", entity{"41CA5E09-3CE2-4094-B108-3ECC257C6FA4", true}},
	}

	for _, test := range tests {
		assert.Equal(t, test.entity.valid, test.entity.IsValid(), test.name)
		assert.Equal(t, test.entity.ID, ID(test.entity.String()), test.name)
	}
}
