package data_test

import (
	"testing"

	"github.com/kamilsk/form-api/data"
	"github.com/stretchr/testify/assert"
)

func TestUUID_IsValid(t *testing.T) {
	for _, tc := range []struct {
		uuid     data.UUID
		expected bool
	}{
		{"invalid", false},
		{"41ca5e09-3ce2-0094-b108-3ecc257c6fa4", false},
		{"41ca5e09-3ce2-1094-b108-3ecc257c6fa4", false},
		{"41ca5e09-3ce2-2094-b108-3ecc257c6fa4", false},
		{"41ca5e09-3ce2-3094-b108-3ecc257c6fa4", false},
		{"41ca5e09-3ce2-4094-b108-3ecc257c6fa4", true},
		{"41ca5e09-3ce2-5094-b108-3ecc257c6fa4", false},
		{"41ca5e09-3ce2-6094-b108-3ecc257c6fa4", false},
		{"41CA5E09-3CE2-4094-B108-3ECC257C6FA4", true},
	} {
		assert.Equal(t, tc.expected, tc.uuid.IsValid())
		assert.Equal(t, tc.uuid, data.UUID(tc.uuid.String()))
	}
}
