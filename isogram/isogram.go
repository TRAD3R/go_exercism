package isogram

import "unicode"

// IsIsogram - determines if a word or phrase is an isogram
func IsIsogram(str string) bool {
	uniqueChars := map[rune]rune{}
	for _, char := range []rune(str) {
		if unicode.IsLetter(char) {
			lowerChar := unicode.ToLower(char)
			if _, ok := uniqueChars[lowerChar]; ok {
				return false
			}
			uniqueChars[lowerChar] = char
		}
	}

	return true
}
