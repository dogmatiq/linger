package linger

import "time"

// DurationTransform is a function that applies some transformation to
// time.Duration values.
type DurationTransform func(time.Duration) time.Duration

// Identity is a DurationTransform that returns the input duration unchanged.
func Identity(d time.Duration) time.Duration {
	return d
}
