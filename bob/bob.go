// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

const (
	Question 		string = "Sure."
	Yell 			string = "Whoa, chill out!"
	Yell_Question 	string = "Calm down, I know what I'm doing!"
	Empty 			string = "Fine. Be that way!"
	Other			string = "Whatever."
)
// Hey should have a comment documenting it.
func Hey(remark string) string {
	if isEmpty(remark) {
		return Empty
	}

	if isQuestion(remark) {
		if isYell(remark) {
			return Yell_Question
		} else {
			return Question
		}
	}

	if isYell(remark) {
		return Yell
	}


	return Other
}

func isEmpty(remark string) bool {
	for _, c := range []rune(remark) {
		if unicode.IsGraphic(c) && !unicode.IsSpace(c){
			return false
		}
	}

	return true
}

func isYell(remark string) bool {
	hasLetter := false
	for _, c := range []rune(remark) {
		if unicode.IsLetter(c) {
			hasLetter = true
			if unicode.IsLower(c) {
				return false
			}
		}
	}

	return hasLetter
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(strings.TrimSpace(remark), "?")
}
