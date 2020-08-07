package anagram

import "strings"

// Detect - return the anagrams list
func Detect(subject string, candidates []string) []string {
	detected := []string{}

	for _, candidate := range candidates {
		if len(candidate) != len(subject) {
			continue
		}
		// words are not anagrams of themselves (case-insensitive)
		if strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		}

		if isAnagram(strings.ToLower(subject), strings.ToLower(candidate)) {
			detected = append(detected, candidate)
		}
	}

	return detected
}

// isAnagram - check is candidate is anagram for subject
func isAnagram(subject, candidate string) bool {
	mapSubject := getMapRune(subject)

	for _, r := range candidate {
		mapSubject[r] = mapSubject[r] - 1
		if mapSubject[r] == 0 {
			delete(mapSubject, r)
		}
	}

	return len(mapSubject) == 0
}

// getMapRune - convert string to map[rune]bool
func getMapRune(str string) map[rune]int {
	mapRune := map[rune]int{}

	for _, r := range str {
		if v, ok := mapRune[r]; !ok {
			mapRune[r] = 1
		} else {
			mapRune[r] = v + 1
		}
	}

	return mapRune
}
