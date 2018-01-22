package domen_test

import (
	"testing"

	"github.com/kamilsk/form-api/domen"
	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	for _, test := range []struct {
		name    string
		uuid    domen.UUID
		isValid bool
	}{
		{"empty", "", false},
		{"invalid", "abc-def-ghi", false},
		{"not v4 [0]", "41ca5e09-3ce2-0094-b108-3ecc257c6fa4", false},
		{"not v4 [1]", "41ca5e09-3ce2-1094-b108-3ecc257c6fa4", false},
		{"not v4 [2]", "41ca5e09-3ce2-2094-b108-3ecc257c6fa4", false},
		{"not v4 [3]", "41ca5e09-3ce2-3094-b108-3ecc257c6fa4", false},
		{"v4 [lower]", "41ca5e09-3ce2-4094-b108-3ecc257c6fa4", true},
		{"v4 [upper]", "41CA5E09-3CE2-4094-B108-3ECC257C6FA4", true},
		{"not v4 [5]", "41ca5e09-3ce2-5094-b108-3ecc257c6fa4", false},
	} {
		assert.Equal(t, test.uuid == "", test.uuid.IsEmpty(), test.name)
		assert.Equal(t, test.isValid, test.uuid.IsValid(), test.name)
		assert.Equal(t, test.uuid, domen.UUID(test.uuid.String()), test.name)
	}
}
