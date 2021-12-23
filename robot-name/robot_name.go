package robotname

import (
	"fmt"
	"math/rand"
	"time"
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

	if len(existedNames) == 26*26*10*10*10 {
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
	r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
