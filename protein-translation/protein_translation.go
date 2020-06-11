package protein

import (
	"errors"
)

var (
	// Codons - codons and resulting Amino Acids
	Codons = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	// ErrInvalidBase - custom error
	ErrInvalidBase = errors.New("error invalid base")

	// ErrStop - custom error
	ErrStop = errors.New("error stop sequence")
)

// FromCodon - translates codon into protein
func FromCodon(codon string) (string, error) {
	protein, ok := Codons[codon]
	if !ok {
		return "", ErrInvalidBase
	}

	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}

// FromRNA - translates RNA sequences into proteins
func FromRNA(rna string) ([]string, error) {
	proteins := []string{}
	runes := []rune(rna)

	for i := 0; i < len(runes); i += 3 {
		codon := string(runes[i : i+3])
		protein, err := FromCodon(codon)
		if err != nil && err != ErrStop {
			return proteins, err
		}

		if err == ErrStop {
			return proteins, nil
		}
		proteins = append(proteins, protein)
	}

	return proteins, nil
}
