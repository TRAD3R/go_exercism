package raindrops

import "strconv"

// Convert convert a number into a string that contains
// raindrop sounds corresponding to certain potential factors
func Convert(num int) string {
	var res string
	if num%3 == 0 {
		res += "Pling"
	}
	if num%5 == 0 {
		res += "Plang"
	}
	if num%7 == 0 {
		res += "Plong"
	}

	if len(res) > 0 {
		return res
	}

	return strconv.Itoa(num)
}
