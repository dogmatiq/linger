package linger

import (
	"context"
	"math"
	"sync/atomic"
	"time"
)

// BackoffStrategy is a function for computing delays between attempts to
// perform some application-defined operation.
//
// n is the number of successive failures since the last success, including this
// one.
//
// err is the error describing the operation's failure, if known. A nil error
// does not indicate a success.
type BackoffStrategy func(n int, err error) time.Duration

// DefaultBackoffStrategy is the default strategy used by Backoff.
//
// It is a conservative policy favouring large delay times under the assumption
// that the operation is expensive.
var DefaultBackoffStrategy BackoffStrategy = ExponentialBackoff(3 * time.Second)

// Backoff introduces delays between atempts to perform some application-defined
// operation.
type Backoff struct {
	// Strategy is the backoff strategy used to compute the "fundamental wait
	// duration". If it is nil, DefaultBackoffStrategy is used.
	Strategy BackoffStrategy

	// Transform is applied to the "fundamental wait duration" to produce the
	// "transformed wait duration".
	//
	// If it is nil, FullJitter is used. The Identity transform can be used to
	// disable the transform.
	Transform DurationTransform

	// Min is a lower bound for the "transformed wait duration", in combination
	// with Max it is used to produce the "bounded wait duration".
	Min time.Duration

	// Min is a upper bound for the "transformed wait duration", in combination
	// with Min it is used to produce the "bounded wait duration".
	Max time.Duration

	// failures is the number of successive failures that have occurred.
	failures uint32
}

// Ok marks the most recent attempt as a success.
func (b *Backoff) Ok() {
	atomic.StoreUint32(&b.failures, 0)
}

// Fail marks the most recent attempt as a failure and returns the "bounded wait
// duration", which is the time to wait before retrying the operation.
//
// err is the error describing the operation's failure, if known. A nil error
// does not indicate a success.
func (b *Backoff) Fail(err error) time.Duration {
	min := Longest(b.Min, 0)
	max := MustCoalesce(b.Max, 1*time.Hour)

	strategy := b.Strategy
	if strategy == nil {
		strategy = DefaultBackoffStrategy
	}

	transform := b.Transform
	if transform == nil {
		transform = FullJitter
	}

	n := atomic.AddUint32(&b.failures, 1)

	d := strategy(int(n), err) // fundamental
	d = transform(d)           // transformed
	d = Limit(d, min, max)     // bounded

	return d
}

// Sleep marks the most recent attempt and pauses the current goroutine until
// the "bounded wait duration" has elapsed.
//
// It sleeps until the duration elapses or ctx is canceled, whichever is first.
// If ctx is canceled before the duration elapses it returns ctx.Err(),
// otherwise it returns nil.
//
// err is the error describing the operation's failure, if known. A nil error
// does not indicate a success.
func (b *Backoff) Sleep(ctx context.Context, err error) error {
	return Sleep(ctx, b.Fail(err))
}

// ExponentialBackoff returns a BackoffStrategy that uses binary exponential
// backoff (BEB).
//
// The unit delay is doubled after each successive failure.
func ExponentialBackoff(unit time.Duration) BackoffStrategy {
	if unit <= 0 {
		panic("the unit duration must be postive")
	}

	u := unit.Seconds()

	return func(n int, _ error) time.Duration {
		scale := math.Pow(2, float64(n-1))
		seconds := u * scale

		return FromSeconds(seconds)
	}
}

// ConstantBackoff returns a BackoffStrategy that returns a fixed wait duration.
func ConstantBackoff(d time.Duration) BackoffStrategy {
	return func(_ int, _ error) time.Duration {
		return d
	}
}

// LinearBackoff returns a BackoffStrategy that increases the wait duration
// linearly.
//
// The unit delay is multiplied by the number of successive failures.
func LinearBackoff(unit time.Duration) BackoffStrategy {
	return func(n int, _ error) time.Duration {
		return time.Duration(n) * unit
	}
}
