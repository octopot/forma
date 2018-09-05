package domain

// Template TODO issue#173
type Template string

// IsEmpty TODO issue#173
func (t Template) IsEmpty() bool {
	return t == ""
}
