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
		{"a0eebc99-9c0b-0ef8-bb6d-6bb9bd380a11", false},
		{"a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11", false},
		{"a0eebc99-9c0b-2ef8-bb6d-6bb9bd380a11", false},
		{"a0eebc99-9c0b-3ef8-bb6d-6bb9bd380a11", false},
		{"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", true},
		{"a0eebc99-9c0b-5ef8-bb6d-6bb9bd380a11", false},
		{"a0eebc99-9c0b-6ef8-bb6d-6bb9bd380a11", false},
		{"A0EEBC99-9C0B-4EF8-BB6D-6BB9BD380A11", true},
	} {
		assert.Equal(t, tc.expected, tc.uuid.IsValid())
		assert.Equal(t, tc.uuid, data.UUID(tc.uuid.String()))
	}
}
