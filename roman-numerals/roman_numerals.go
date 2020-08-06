package main

import (
	"errors"
	"strings"
)

// ArabToRoman -  correspondence of Arabic and Roman numerals
var ArabToRoman = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

// ToRomanNumeral - converter arabic to roman numeral
func ToRomanNumeral(number int) (string, error) {
	if number < 1 || number > 3000 {
		return "", errors.New("Not valid number")
	}
	romanNumber := []string{}
	for i := 1000; i >= 1; i /= 10 {
		digit := number / i
		if digit == 0 {
			continue
		}

		romanNumber = append(romanNumber, getDigit(digit, i))
		number = number % i
	}

	return strings.Join(romanNumber, ""), nil
}

// getDigit - convert current arabic digit to roman
func getDigit(number, multi int) (romans string) {
	if number == 4 || number == 9 {
		number++
		romans = ArabToRoman[multi] + ArabToRoman[number*multi]
	} else {
		if number >= 5 {
			romans = ArabToRoman[5*multi]
			number = number - 5
		}
		for i := 0; i < number; i++ {
			romans += ArabToRoman[multi]
		}
	}

	return romans
}
