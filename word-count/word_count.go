package wordcount

import (
	"strings"
	"unicode"
)

// Frequency - resulting map
type Frequency map[string]int

// WordCount - word counter in sentence
func WordCount(str string) Frequency {
	str = strings.ToLower(str)
	res := Frequency{}

	subs := strings.FieldsFunc(str, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})
	for _, key := range subs {
		if len(key) == 0 {
			continue
		}
		key = strings.Trim(key, "'")
		if v, ok := res[key]; ok {
			res[key] = v + 1
		} else {
			res[key] = 1
		}
	}

	return res
}
