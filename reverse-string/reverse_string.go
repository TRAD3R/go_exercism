package reverse

// Reverse - reverse a string
func Reverse(str string) string {
	runes := []rune(str)
	reverseRune := []rune{}

	for i := len(runes) - 1; i >= 0; i-- {
		reverseRune = append(reverseRune, runes[i])
	}

	return string(reverseRune)
}
