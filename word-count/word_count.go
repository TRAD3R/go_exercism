package wordcount

import (
	"regexp"
	"strings"
)

// Frequency - resulting map
type Frequency map[string]int

// WordCount - word counter in sentence
func WordCount(str string) Frequency {
	str = strings.ToLower(str)
	res := Frequency{}
	str = replaceWrongChars(str, "[^a-z0-9']", " ")

	subs := strings.Split(str, " ")
	for _, key := range subs {
		if len(key) == 0 {
			continue
		}
		key = replaceWrongChars(key, "^'", "")
		key = replaceWrongChars(key, "'$", "")
		if v, ok := res[key]; ok {
			res[key] = v + 1
		} else {
			res[key] = 1
		}
	}

	return res
}

// replaceWrongChars - regexp replacer
func replaceWrongChars(source, expr, repl string) string {
	reg := regexp.MustCompile(expr)
	return reg.ReplaceAllString(source, repl)
}
