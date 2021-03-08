package pkg

import (
	"math/rand"
	"time"
)

// RealRandom generates a random number from a seed
func RealRandom(i int) int {
	// Generate seed for random number
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(i)
}

// HasPosition checks whether a slice of positions contains a spesific position
func HasPosition(b []Position, p Position) bool {
	for _, pos := range b {
		if pos == p {
			return true
		}
	}

	return false
}
