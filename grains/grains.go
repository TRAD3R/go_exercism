package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Number must be greater then 0 and lower then 65")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	return (1<<63)*2 - 1
}
