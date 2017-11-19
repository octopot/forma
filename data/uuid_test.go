package data_test

import (
	"testing"

	"github.com/kamilsk/form-api/data"
)

func TestUUID_IsValid(t *testing.T) {
	for _, tc := range []struct {
		uuid     data.UUID
		expected bool
	}{
		{uuid: "invalid", expected: false},
		{uuid: "a0eebc99-9c0b-1ef8-bb6d-6bb9bd380a11", expected: true},
		{uuid: "a0eebc99-9c0b-2ef8-bb6d-6bb9bd380a11", expected: true},
		{uuid: "a0eebc99-9c0b-3ef8-bb6d-6bb9bd380a11", expected: true},
		{uuid: "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", expected: true},
		{uuid: "a0eebc99-9c0b-5ef8-bb6d-6bb9bd380a11", expected: true},
		{uuid: "a0eebc99-9c0b-0ef8-bb6d-6bb9bd380a11", expected: false},
		{uuid: "a0eebc99-9c0b-6ef8-bb6d-6bb9bd380a11", expected: false},
		{uuid: "A0EEBC99-9C0B-1EF8-BB6D-6BB9BD380A11", expected: true},
	} {
		if tc.expected != tc.uuid.IsValid() {
			if tc.expected {
				t.Errorf("expected valid UUID, obtained invalid: %s", tc.uuid)
			} else {
				t.Errorf("expected invalid UUID, obtained valid: %s", tc.uuid)
			}
		}
	}
}
