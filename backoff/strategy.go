package backoff

import (
	"math"
	"time"

	"github.com/dogmatiq/linger"
)

// DefaultStrategy is the default strategy used if none is specified.
//
// It is a conservative policy favouring large delay times under the assumption
// that the operation is expensive.
var DefaultStrategy Strategy = WithTransforms(
	Exponential(3*time.Second),
	linger.FullJitter,
	linger.Limiter(0, 1*time.Hour),
)

// Strategy is a function for computing delays between attempts to perform some
// application-defined operation.
//
// err is the error describing the operation's failure, if known. A nil error
// does not indicate a success.
//
// n is the number of successive failures since the last success, not including
// the failure indicated by err.
type Strategy func(err error, n uint) time.Duration

// Exponential returns a Strategy that uses binary exponential backoff (BEB).
//
// The unit delay is doubled after each successive failure.
func Exponential(unit time.Duration) Strategy {
	if unit <= 0 {
		panic("the unit duration must be postive")
	}

	u := unit.Seconds()

	return func(_ error, n uint) time.Duration {
		scale := math.Pow(2, float64(n))
		seconds := u * scale
		return linger.FromSeconds(seconds)
	}
}

// Constant returns a Strategy that returns a fixed wait duration.
func Constant(d time.Duration) Strategy {
	return func(error, uint) time.Duration {
		return d
	}
}

// Linear returns a Strategy that increases the wait duration linearly.
//
// The unit delay is multiplied by the number of successive failures.
func Linear(unit time.Duration) Strategy {
	return func(_ error, n uint) time.Duration {
		return time.Duration(n+1) * unit
	}
}

// WithTransforms returns a strategy that transforms the result of s using each
// of the given transforms in order.
func WithTransforms(s Strategy, transforms ...linger.DurationTransform) Strategy {
	return func(err error, n uint) time.Duration {
		d := s(err, n)

		for _, x := range transforms {
			d = x(d)
		}

		return d
	}
}
