package strings

// FirstNotEmpty returns a first non-empty string.
func FirstNotEmpty(strings ...string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}
	return ""
}

// NotEmpty filters empty strings in-place.
func NotEmpty(strings []string) []string {
	filtered := strings[:0]
	for _, str := range strings {
		if str != "" {
			filtered = append(filtered, str)
		}
	}
	return filtered
}
