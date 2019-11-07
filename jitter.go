package linger

import (
	"math/rand"
	"time"
)

// FullJitter is a DurationTransform that applies "full jitter" to the input
// duration.
//
// The output duration is a random duration between 0 and the input duration,
// inclusive.
func FullJitter(d time.Duration) time.Duration {
	return Rand(0, d)
}

// ProportionalJitter returns a DurationTransform that applies "proportional
// jitter" to the input duration.
//
// The returned transform generates a random jitter amount proportional to the
// input duration. This jitter amount is added to the input duration to produce
// the output duration.
//
// The jitter proportion is given by p. For example, a value of 0.1 adds up to
// 10% to the input duration. p may be negative to indicate that the jitter
// amount should be subtracted from the input duration.
func ProportionalJitter(p float64) DurationTransform {
	return func(d time.Duration) time.Duration {
		j := Multiply(d, p)
		return Rand(d, d+j)
	}
}

// Rand returns a random duration between a and b, inclusive.
func Rand(a, b time.Duration) time.Duration {
	if b < a {
		a, b = b, a
	}

	min := int64(a)
	max := int64(b) + 1
	r := min + rand.Int63n(max-min)

	return time.Duration(r)
}
