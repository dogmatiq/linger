package linger

import "time"

// Multiply returns the result of multiplying d by v.
func Multiply(d time.Duration, v float64) time.Duration {
	return FromSeconds(d.Seconds() * v)
}

// Divide returns the result of dividing d by v.
func Divide(d time.Duration, v float64) time.Duration {
	return FromSeconds(d.Seconds() / v)
}
