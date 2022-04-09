package rng

import (
	"math/rand"
	"time"
)

const percent = 100

// Seed the random number generator with the current time in nanoseconds.
func Seed() {
	rand.Seed(time.Now().UnixNano())
}

// Chance will return true n% of the time.
func Chance(n int) bool {
	Seed()

	return rand.Intn(percent) < n
}

// Select a random integer from the population p, excluding e.
func Select(p, e int) int {
	Seed()

	r := rand.Intn(p)

	if r == e {
		r = Select(p, e)
	}

	return r
}
