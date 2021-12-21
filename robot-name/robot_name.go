package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxRandNames = 26 * 26 * 10 * 10 * 10
	letterMax    = 26
	letterShift  = 65
)

var (
	random       = rand.New(rand.NewSource(time.Now().UnixNano()))
	existedNames = map[string]bool{}
)

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if len(r.name) == 0 {

		if len(existedNames) == maxRandNames {
			return "", fmt.Errorf("Out of names")
		}

		r.name = generateName()
		existedNames[r.name] = true
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	for {
		newName := fmt.Sprintf(
			"%s%s%d%d%d",
			randomLetter(),
			randomLetter(),
			randomInt(),
			randomInt(),
			randomInt(),
		)

		if _, ok := existedNames[newName]; ok {
			continue
		}

		return newName
	}
}

func randomLetter() string {
	letter := random.Intn(letterMax) + letterShift

	return string(rune(letter))
}

func randomInt() int {
	return random.Intn(10)
}
