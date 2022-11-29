package linger

import (
	"math"
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

	diff := int64(b - a)

	return a + time.Duration(
		inclusiveRand(diff),
	)
}

// inclusiveRand returns a positive unsigned integer less than or equal to n.
func inclusiveRand(n int64) int64 {
	if n == 0 {
		return 0
	}

	// If n is equal to the maximum value of an int64 we can't use rand.Int63n()
	// because it uses a half-open range. As a work-around, we produce a random
	// unsigned integer and discard the sign bit.
	if n == math.MaxInt64 {
		return int64(rand.Uint64() & math.MaxInt64)
	}

	return rand.Int63n(n + 1)
}
