package linger

import (
	"context"
	"time"
)

// FromContextDeadline returns the duration until the deadline of ctx is
// reached.
//
// ok is false if ctx does not have a deadline.
func FromContextDeadline(ctx context.Context) (d time.Duration, ok bool) {
	if dl, ok := ctx.Deadline(); ok {
		return time.Until(dl), true
	}

	return 0, false
}

// ContextWithTimeout returns a content with a deadline some duration after the
// current time.
//
// The timeout duration is computed by finding the first of the supplied
// durations that is positive. It uses a zero duration if none of the supplied
// durations are positive.
func ContextWithTimeout(
	ctx context.Context,
	durations ...time.Duration,
) (context.Context, func()) {
	return ContextWithTimeoutX(ctx, Identity, durations...)
}

// ContextWithTimeoutX returns a content with a deadline some duration after the
// current time.
//
// The timeout duration is computed by finding the first of the supplied
// durations that is positive, then applying the transform x. It uses a zero
// duration if none of the supplied durations are positive.
//
// The transform can be used to apply jitter to the timeout duration, for
// example, by using one of the built-in jitter transforms such as FullJitter()
// or ProportionalJitter().
func ContextWithTimeoutX(
	ctx context.Context,
	x DurationTransform,
	durations ...time.Duration,
) (context.Context, func()) {
	d, _ := Coalesce(durations...)
	return context.WithTimeout(ctx, x(d))
}
