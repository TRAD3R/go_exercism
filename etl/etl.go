package etl

import (
	"strings"
)

// Transform - extracts some Scrabble scores from a legacy system
func Transform(input map[int][]string) map[string]int {
	res := map[string]int{}

	for score, chars := range input {
		for _, letter := range chars {
			l := strings.ToLower(letter)
			res[l] = score
		}
	}

	return res
}
