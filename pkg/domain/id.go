package domain

import "regexp"

// [4] means that supported only v4.
var uuid = regexp.MustCompile(`(?i:^[0-9A-F]{8}-[0-9A-F]{4}-[4][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$)`)

// ID wraps built-in `string` type and provides useful methods above it.
type ID string

// IsEmpty returns true if the ID is empty.
func (s ID) IsEmpty() bool {
	return s == ""
}

// IsValid returns true if the ID is not empty and compatible with RFC 4122.
func (s ID) IsValid() bool {
	return !(s == "") && uuid.MatchString(string(s)) // IsEmpty and String were inlined manually
}

// String implements built-in `fmt.Stringer` interface and returns the underlying string value.
func (s ID) String() string {
	return string(s)
}
