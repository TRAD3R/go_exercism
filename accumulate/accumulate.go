package accumulate

// Accumulate - returns a new collection containing the result of applying that operation to each element of the input collection
func Accumulate(s []string, f func(string) string) []string {
	res := []string{}
	for _, word := range s {
		res = append(res, f(word))
	}

	return res
}
