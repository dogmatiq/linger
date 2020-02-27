package linger

import "time"

// First returns the first of its arguments for which the predicate
// function p returns true.
//
// If the p returns false for all values, v is the zero-value and ok is false.
func First(p DurationPredicate, values ...time.Duration) (v time.Duration, ok bool) {
	for _, v := range values {
		if p(v) {
			return v, true
		}
	}

	return 0, false
}

// MustFirst returns the first of its arguments for which the predicate
// function p returns true.
//
// It panics if p returns false for all values.
func MustFirst(p DurationPredicate, values ...time.Duration) time.Duration {
	if t, ok := First(p, values...); ok {
		return t
	}

	panic("the predicate did not match any input values")
}

// FirstT returns the first of its arguments for which the predicate
// function p returns true.
//
// If the p returns false for all values, v is the zero-value and ok is false.
func FirstT(p TimePredicate, values ...time.Time) (v time.Time, ok bool) {
	for _, v := range values {
		if p(v) {
			return v, true
		}
	}

	return time.Time{}, false
}

// MustFirstT returns the first of its arguments for which the predicate
// function p returns true.
//
// It panics if p returns false for all values.
func MustFirstT(p TimePredicate, values ...time.Time) time.Time {
	if t, ok := FirstT(p, values...); ok {
		return t
	}

	panic("the predicate did not match any input values")
}

// Defaulter returns a DurationTransform that returns the first value for which
// the predicate function p returns true. The transform input value is checked
// first, then each of the given values.
//
// It panics if p returns false for all values.
func Defaulter(p DurationPredicate, values ...time.Duration) DurationTransform {
	return func(v time.Duration) time.Duration {
		if p(v) {
			return v
		}

		return MustFirst(p, values...)
	}
}
