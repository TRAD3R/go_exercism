package pangram

import (
	"strings"
	"unicode"
)

// IsPangram - check is the string a panagram
func IsPangram(str string) bool {
	var alphabet = map[rune]bool{
		'a': true,
		'b': true,
		'c': true,
		'd': true,
		'e': true,
		'f': true,
		'g': true,
		'h': true,
		'i': true,
		'j': true,
		'k': true,
		'l': true,
		'm': true,
		'n': true,
		'o': true,
		'p': true,
		'q': true,
		'r': true,
		's': true,
		't': true,
		'u': true,
		'v': true,
		'w': true,
		'x': true,
		'y': true,
		'z': true,
	}
	for _, r := range strings.ToLower(str) {
		if unicode.IsLetter(r) {
			delete(alphabet, r)
		}
	}

	return len(alphabet) == 0
}
