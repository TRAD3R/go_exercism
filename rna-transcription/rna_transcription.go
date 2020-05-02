package strand

import "strings"

/**
G -> C
C -> G
T -> A
A -> U
 */
func ToRNA(dna string) string {
	rnaMap := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}

	var rna []string

	for _, nucleotide := range strings.Split(dna, "") {
		rna  = append(rna, rnaMap[nucleotide])
	}
	return strings.Join(rna, "")
}
