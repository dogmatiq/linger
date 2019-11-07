package linger

import (
	"context"
	"time"
)

// Sleep pauses the current goroutine until some duration has elapsed.
//
// The sleep duration computed by finding the first of the supplied durations
// that is positive. It returns immediately if none of the supplied durations
// are positive.
//
// It sleeps until the duration elapses or ctx is canceled, whichever is first.
// If ctx is canceled before the duration elapses it returns ctx.Err(),
// otherwise it returns nil.
func Sleep(ctx context.Context, durations ...time.Duration) error {
	return SleepX(ctx, Identity, durations...)
}

// SleepX pauses the current goroutine until some duration has elapsed.
//
// The sleep duration computed by finding the first of the supplied durations
// that is positive, then applying the transform x. It returns immediately if
// none of the supplied durations are positive.
//
// The transform can be used to apply jitter to the sleep duration, for example,
// by using one of the built-in jitter transforms such as FullJitter() or
// ProportionalJitter().
//
// It sleeps until the duration elapses or ctx is canceled, whichever is first.
// If ctx is canceled before the duration elapses it returns ctx.Err(),
// otherwise it returns nil.
func SleepX(
	ctx context.Context,
	x DurationTransform,
	durations ...time.Duration,
) error {
	d, _ := Coalesce(durations...)
	return sleep(ctx, x, d)
}

// SleepUntil pauses the current goroutine until a specific time.
//
// The sleep duration is computed by finding the amount of time until the
// earliest of the supplied times. It returns immediately if the earliest of the
// times is in the past.
//
// It sleeps until the time is reached or ctx is canceled, whichever is first.
// If ctx is canceled before the time is reached it returns ctx.Err(),
// otherwise it returns nil.
func SleepUntil(ctx context.Context, times ...time.Time) error {
	return SleepUntilX(ctx, Identity, times...)
}

// SleepUntilX pauses the current goroutine until a specific time.
//
// The sleep duration is computed by finding the amount of time until the
// earliest of the supplied times, then applying the transform x. It returns
// immediately if the earliest of the times is in the past.
//
// The transform can be used to apply jitter to the sleep duration, for example,
// by using one of the built-in jitter transforms such as FullJitter() or
// ProportionalJitter().
//
// It sleeps until the time is reached or ctx is canceled, whichever is first.
// If ctx is canceled before the time is reached it returns ctx.Err(),
// otherwise it returns nil.
func SleepUntilX(
	ctx context.Context,
	x DurationTransform,
	times ...time.Time,
) error {
	d := time.Until(Earliest(times...))
	return sleep(ctx, x, d)
}

func sleep(ctx context.Context, x DurationTransform, d time.Duration) error {
	if d <= 0 {
		return ctx.Err()
	}

	t := time.NewTimer(x(d))
	defer t.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
