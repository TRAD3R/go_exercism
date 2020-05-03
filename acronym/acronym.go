package acronym

import (
	"regexp"
	"unicode"
)

// Abbreviate convert a phrase to its acronym.
func Abbreviate(s string) string {
	split := regexp.MustCompile(`\s|-`)
	words := split.Split(s, -1)
	acronym := []rune{}

	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		chars := []rune(word)
		for _, char := range chars {
			if unicode.IsLetter(char) {
				acronym = append(acronym, unicode.ToUpper(char))
				break
			}
		}
	}

	return string(acronym)
}
