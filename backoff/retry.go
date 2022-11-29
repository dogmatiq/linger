package backoff

import (
	"context"

	"github.com/dogmatiq/linger"
)

// Retry calls the given function until it succeeds.
//
// Each subsequent call is delayed according to the given backoff strategy.
// If s is nil, DefaultStrategy is used.
//
// It returns ctx.Err() if ctx is canceled before fn() succeeds.
// n is the number of times that fn() failed, even if err is non-nil.
func Retry(
	ctx context.Context,
	s Strategy,
	fn func(ctx context.Context) error,
) (n uint, err error) {
	if s == nil {
		s = DefaultStrategy
	}

	for {
		err := fn(ctx)
		if err == nil {
			return n, nil
		}

		d := s(err, n)
		n++

		if err := linger.Sleep(ctx, d); err != nil {
			return n, err
		}
	}
}
