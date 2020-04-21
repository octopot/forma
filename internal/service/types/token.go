package types

// Token represents user access token.
type Token string

// IsEmpty returns true if the Token is empty.
func (token Token) IsEmpty() bool {
	return token == ""
}

// IsValid returns true if the Token is not empty and has the appropriate format.
func (token Token) IsValid() bool {
	return !token.IsEmpty() && uuid.MatchString(string(token))
}

// String implements built-in `fmt.Stringer` interface and returns the underlying string value.
func (token Token) String() string {
	return string(token)
}
