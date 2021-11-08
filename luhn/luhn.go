package luhn

import (
	"strconv"
	"strings"
)

// Valid - check for valid number by the Luhn algorithm
func Valid(number string) bool {
	if len(number) < 2 {
		return false
	}

	numWithoutSpace := strings.ReplaceAll(number, " ", "")
	if len(numWithoutSpace) < 2 {
		return false
	}

	sum, err := getSum(numWithoutSpace)
	if err != nil {
		return false
	}

	return sum%10 == 0

}

// getSum - get digit sum
func getSum(num string) (int, error) {
	sum := 0

	isEven := len(num) % 2 == 0
	for _, r := range num{
		digit, err := strconv.Atoi(string(r))

		if err != nil {
			return 0, err
		}

		if isEven {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}



		sum += digit
		isEven = !isEven
	}

	return sum, nil
}
