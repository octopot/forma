package domain_test

import (
	"testing"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	for _, test := range []struct {
		name    string
		uuid    domain.ID
		isValid bool
	}{
		{"empty", "", false},
		{"invalid", "abc-def-ghi", false},
		{"not v4", "41ca5e09-3ce2-3094-b108-3ecc257c6fa4", false},
		{"v4 [lower]", "41ca5e09-3ce2-4094-b108-3ecc257c6fa4", true},
		{"v4 [upper]", "41CA5E09-3CE2-4094-B108-3ECC257C6FA4", true},
	} {
		assert.Equal(t, test.uuid == "", test.uuid.IsEmpty(), test.name)
		assert.Equal(t, test.isValid, test.uuid.IsValid(), test.name)
		assert.Equal(t, test.uuid, domain.ID(test.uuid.String()), test.name)
	}
}
