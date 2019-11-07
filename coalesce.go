package linger

import "time"

// Coalesce returns the first of its arguments that is positive.
//
// If none of the arguments are positive, v is th zero-value and ok is false.
func Coalesce(values ...time.Duration) (v time.Duration, ok bool) {
	return First(Positive, values...)
}

// MustCoalesce returns the first of its arguments that is positive.
//
// It panics if none of the arguments are positive.
func MustCoalesce(values ...time.Duration) time.Duration {
	return MustFirst(Positive, values...)
}

// CoalesceT returns the first of its arguments that is non-zero.
//
// If none of the arguments are non-zero, v is th zero-value and ok is false.
func CoalesceT(values ...time.Time) (v time.Time, ok bool) {
	return FirstT(NonZeroT, values...)
}

// MustCoalesceT returns the first of its arguments that is non-zero.
//
// It panics if none of the arguments are non-zero.
func MustCoalesceT(values ...time.Time) time.Time {
	return MustFirstT(NonZeroT, values...)
}
