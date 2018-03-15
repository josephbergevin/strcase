// Package strcase converts strings to snake_case or CamelCase
package strcase

import (
	"strings"
)

// Converts a string to lower case, inserting the given separator on the word breaks
func toLowerWithSeparator(s, sep string) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		// if i > 0 && v >= 'A' && v <= 'Z' && n[len(n)-1] != sep[0] {
		// 	n += sep + string(v)
		// } else if v == ' ' {
		// 	n += sep
		// } else {
		// 	n = n + string(v)
		// }

		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		nextIsCapital := false
		if i+1 < len(s) {
			w := s[i+1]
			nextIsCapital = w >= 'A' && w <= 'Z'
		}
		if i > 0 && v >= 'A' && v <= 'Z' && n[len(n)-1] != sep[0] && !nextIsCapital {
			// add underscore if next letter is a capital
			n += sep + string(v)
		} else if v == ' ' {
			// replace spaces with underscores
			n += sep
		} else {
			n = n + string(v)
		}
	}
	return strings.ToLower(n)
}

func ToKebab(s string) string {
	s = strings.TrimSpace(s)
	r := strings.NewReplacer("_", "-", " ", "-")
	return toLowerWithSeparator(r.Replace(s), "-")
}

func ToSentence(s string) string {
	s = strings.TrimSpace(s)
	r := strings.NewReplacer("_", " ", "-", " ")
	return toLowerWithSeparator(r.Replace(s), " ")
}

func ToSnake(s string) string {
	s = strings.TrimSpace(s)
	r := strings.NewReplacer(" ", "_", "-", "_")
	return toLowerWithSeparator(r.Replace(s), "_")
}
