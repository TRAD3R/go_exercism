package clock

import "fmt"

type Clock struct {
	Hour   int
	Minute int
}

func New(h, m int) Clock {
	var c Clock

	c.Hour, c.Minute = convertMin2Hour(h*60 + m)
	return c
}

func (c Clock) Add(m int) Clock {
	newTime := c.Hour*60 + c.Minute + m

	c.Hour, c.Minute = convertMin2Hour(newTime)

	return c
}

func (c Clock) Subtract(m int) Clock {
	newTime := c.Hour*60 + c.Minute - m

	c.Hour, c.Minute = convertMin2Hour(newTime)

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

func convertMin2Hour(min int) (h, m int) {
	for min < 0 {
		min = 1440 + min
	}

	hours := min / 60
	m = min % 60

	if hours > 24 {
		h = hours % 24
	} else {
		h = hours
	}

	if h == 24 {
		h = 0
	}

	return h, m
}
