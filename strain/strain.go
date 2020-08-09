package strain

//Ints - type for int slice
type Ints []int

//Lists - type for slice of int slice
type Lists [][]int

//Strings - type for string slice
type Strings []string

// Discard - return values are not match predicate
func (in Ints) Discard(predicate func(int) bool) Ints {
	if in == nil {
		return nil
	}
	out := Ints{}

	for _, i := range in {
		if predicate(i) {
			continue
		}

		out = append(out, i)
	}
	return out
}

// Keep - return values are match predicate
func (in Ints) Keep(predicate func(int) bool) Ints {
	if in == nil {
		return nil
	}
	out := Ints{}

	for _, i := range in {
		if !predicate(i) {
			continue
		}

		out = append(out, i)
	}
	return out
}

// Discard - return values are not match predicate
func (in Lists) Discard(predicate func([]int) bool) Lists {
	if in == nil {
		return nil
	}
	out := Lists{}

	for _, i := range in {
		if predicate(i) {
			continue
		}

		out = append(out, i)
	}
	return out
}

// Keep - return values are match predicate
func (in Lists) Keep(predicate func([]int) bool) Lists {
	if in == nil {
		return nil
	}
	out := Lists{}

	for _, i := range in {
		if !predicate(i) {
			continue
		}

		out = append(out, i)
	}
	return out
}

// Discard - return values are not match predicate
func (in Strings) Discard(predicate func(string) bool) Strings {
	if in == nil {
		return nil
	}
	out := Strings{}

	for _, i := range in {
		if predicate(i) {
			continue
		}

		out = append(out, i)
	}
	return out
}

// Keep - return values are match predicate
func (in Strings) Keep(predicate func(string) bool) Strings {
	if in == nil {
		return nil
	}
	out := Strings{}

	for _, i := range in {
		if !predicate(i) {
			continue
		}

		out = append(out, i)
	}
	return out
}
