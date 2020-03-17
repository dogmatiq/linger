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

	u := float64(unit)

	return func(_ error, n uint) time.Duration {
		scale := math.Pow(2, float64(n))
		nanos := u * scale

		// Overflow check.
		if nanos >= float64(linger.MaxDuration) {
			return linger.MaxDuration
		}

		return time.Duration(nanos)
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
	if unit <= 0 {
		panic("the unit duration must be postive")
	}

	return func(_ error, n uint) time.Duration {
		mult := time.Duration(n) + 1
		delay := mult * unit

		// Overflow check: If delay is negative, there was clearly an overflow
		// because both unit and mult are positive.
		if delay < 0 {
			return linger.MaxDuration
		}

		// Overflow check: Dividing the delay by the unit value should give us
		// back the multiplier that we used. If not, there was an overflow.
		if delay/unit != mult {
			return linger.MaxDuration
		}

		return delay
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
