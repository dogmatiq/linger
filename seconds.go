package linger

import "time"

// FromSeconds returns a duration equivalent to the given number of seconds.
func FromSeconds(s float64) time.Duration {
	return time.Duration(
		s * float64(time.Second),
	)
}
