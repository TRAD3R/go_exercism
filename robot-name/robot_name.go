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
	if r.name != "" {
		return r.name, nil
	}

	if len(existedNames) == maxRandNames {
		return "", fmt.Errorf("Out of names")
	}

	r.name = newName()
	for existedNames[r.name] {
		r.name = newName()
	}

	existedNames[r.name] = true

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func newName() string {
	return fmt.Sprintf(
		"%s%s%d%d%d",
		randomLetter(),
		randomLetter(),
		randomInt(),
		randomInt(),
		randomInt(),
	)
}

func randomLetter() string {
	letter := random.Intn(letterMax) + letterShift

	return string(rune(letter))
}

func randomInt() int {
	return random.Intn(10)
}
