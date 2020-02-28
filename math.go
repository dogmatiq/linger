package linger

import (
	"math"
	"time"
)

const (
	// MinDuration is the smallest duration that can be represented with the
	// time.Duration type.
	MinDuration = time.Duration(math.MinInt64)

	// MaxDuration is the largest duration that can be represented with the
	// time.Duration type.
	MaxDuration = time.Duration(math.MaxInt64)
)

// Multiply returns the result of multiplying d by v.
func Multiply(d time.Duration, v float64) time.Duration {
	return FromSeconds(d.Seconds() * v)
}

// Multiplier returns a DurationTransform that multiplies the input duration by v.
func Multiplier(v float64) DurationTransform {
	return func(d time.Duration) time.Duration {
		return Multiply(d, v)
	}
}

// Divide returns the result of dividing d by v.
func Divide(d time.Duration, v float64) time.Duration {
	return FromSeconds(d.Seconds() / v)
}

// Divider returns a DurationTransform that divides the input duration by v.
func Divider(v float64) DurationTransform {
	return func(d time.Duration) time.Duration {
		return Divide(d, v)
	}
}

// Shortest returns the smallest of the given durations.
//
// It returns MaxDuration if no durations are supplied.
func Shortest(durations ...time.Duration) time.Duration {
	min := MaxDuration

	for _, d := range durations {
		if d < min {
			min = d
		}
	}

	return min
}

// Longest returns the largest of the given durations.
//
// It returns MinDuration if no durations are supplied.
func Longest(durations ...time.Duration) time.Duration {
	max := MinDuration

	for _, d := range durations {
		if d > max {
			max = d
		}
	}

	return max
}

// Earliest returns the earliest of the given times.
//
// It returns the zero-value if no times are supplied.
func Earliest(times ...time.Time) time.Time {
	var min time.Time

	for i, t := range times {
		if i == 0 || t.Before(min) {
			min = t
		}
	}

	return min
}

// Latest returns the latest of the given times.
//
// It returns the zero-value if no times are supplied.
func Latest(times ...time.Time) time.Time {
	var max time.Time

	for _, t := range times {
		if t.After(max) {
			max = t
		}
	}

	return max
}

// Limit returns the duration d, capped between a and b, inclusive.
func Limit(d, a, b time.Duration) time.Duration {
	if b < a {
		a, b = b, a
	}

	if d < a {
		return a
	}

	if d > b {
		return b
	}

	return d
}

// Limiter returns a DurationTransform that limits the input duration
// between a and b, inclusive.
func Limiter(a, b time.Duration) DurationTransform {
	return func(d time.Duration) time.Duration {
		return Limit(d, a, b)
	}
}

// LimitT returns the time t, capped between a and b, inclusive.
func LimitT(t, a, b time.Time) time.Time {
	if b.Before(a) {
		a, b = b, a
	}

	if t.Before(a) {
		return a
	}

	if t.After(b) {
		return b
	}

	return t
}
