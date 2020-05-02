package hamming

import "fmt"

// Distance - Calculate the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	dna1, dna2 := []rune(a), []rune(b)

	if len(dna1) != len(dna2) {
		return 0, fmt.Errorf("DNAs have different length")
	}

	distance := 0
	for i, r := range dna1 {
		if r != dna2[i] {
			distance++
		}
	}

	return distance, nil
}
