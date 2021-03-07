package pkg

import (
	"math/rand"
	"time"
)

func realRandom(i int) int {
	// Generate seed for random number
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(i)
}
