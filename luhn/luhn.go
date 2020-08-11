package luhn

import (
	"errors"
	"strconv"
	"unicode"
)

// Valid - check for valid number by the Luhn algorithm
func Valid(number string) bool {
	numWithoutSpace, err := convertToSolidNumber(number)
	if err != nil || len(numWithoutSpace) < 2 {
		return false
	}

	sum, err := getSum(numWithoutSpace)
	if err != nil {
		return false
	}

	return sum%10 == 0

}

// convertToSolidNumber - remove spaces and check fo non-digits numbers
func convertToSolidNumber(number string) ([]rune, error) {
	res := []rune{}
	for _, r := range number {
		if unicode.IsSpace(r) {
			continue
		}

		if unicode.IsDigit(r) {
			res = append(res, r)
		} else {
			return nil, errors.New("Non-digit symbol")
		}
	}

	return res, nil
}

// getSum - get digit sum
func getSum(num []rune) (int, error) {
	sum := 0

	isEven := false
	for i := len(num) - 1; i >= 0; i-- {
		digit := 0
		var err error
		if isEven {
			digit, err = getDouble(string(num[i]))
		} else {
			digit, err = strconv.Atoi(string(num[i]))
		}

		if err != nil {
			return 0, err
		}

		sum += digit
		isEven = !isEven
	}

	return sum, nil
}

// getDouble - double digit
func getDouble(r string) (int, error) {
	num, err := strconv.Atoi(r)
	if err != nil {
		return 0, err
	}

	num *= 2
	if num > 9 {
		num = num - 9
	}

	return num, nil
}
