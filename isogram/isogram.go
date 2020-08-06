package isogram

import "unicode"

// IsIsogram - determines if a word or phrase is an isogram
func IsIsogram(str string) bool {
	uniqueChars := map[rune]bool{}
	for _, char := range str {
		if unicode.IsLetter(char) {
			lowerChar := unicode.ToLower(char)
			if uniqueChars[lowerChar] {
				return false
			}
			uniqueChars[lowerChar] = true
		}
	}

	return true
}
