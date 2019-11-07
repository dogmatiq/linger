package linger

import (
	"math/rand"
	"time"
)

// Rand returns a random duration between min and max, inclusive.
func Rand(min, max time.Duration) time.Duration {
	l := int64(min)
	u := int64(max) + 1
	r := l + rand.Int63n(u-l)
	return time.Duration(r)
}
