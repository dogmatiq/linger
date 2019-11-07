package linger

import "time"

// DurationPredicate is a predicate function for time.Duration values.
type DurationPredicate func(time.Duration) bool

// TimePredicate is a predicate function for time.Time values.
type TimePredicate func(time.Time) bool

// NonZero is a DurationPredicate that returns true if d is non-zero.
func NonZero(d time.Duration) bool {
	return d != 0
}

// Positive is a DurationPredicate that returns true if d is positive.
func Positive(d time.Duration) bool {
	return d > 0
}

// Negative is a DurationPredicate that returns true if d is negative.
func Negative(d time.Duration) bool {
	return d < 0
}

// NonZeroT is a TimePredicate that returns true if t is non-zero.
func NonZeroT(t time.Time) bool {
	return !t.IsZero()
}
