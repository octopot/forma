package strings

import "strings"

// Concat concatenates all passed strings.
func Concat(substrings ...string) string {
	b := strings.Builder{}
	for _, str := range substrings {
		b.WriteString(str)
	}
	return b.String()
}
