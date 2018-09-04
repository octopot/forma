package strings

// Concat concatenates all passed strings.
func Concat(strings ...string) string {
	b := &Builder{}
	for _, str := range strings {
		b.WriteString(str)
	}
	return b.String()
}

// First returns first non-empty string.
func First(strings ...string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}
	return ""
}
