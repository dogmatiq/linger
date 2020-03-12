package backoff

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/dogmatiq/linger"
)

// Counter keeps track of the number of times an operation has failed to
// introduce delays between retries.
type Counter struct {
	// Strategy is used to calculate the delay duration.
	// If it is nil, DefaultStrategy is used.
	Strategy Strategy

	// failures is the number of successive failures that have occurred.
	failures uint32 // atomic
}

// Reset marks the most recent attempt as a success, resetting the counter.
func (c *Counter) Reset() {
	atomic.StoreUint32(&c.failures, 0)
}

// Fail marks the most recent attempt as a failure and returns the duration to
// wait before the operation should be retried.
//
// err is the error describing the operation's failure condition, if known. A
// nil error does not indicate a success.
func (c *Counter) Fail(err error) time.Duration {
	n := atomic.AddUint32(&c.failures, 1)

	s := c.Strategy
	if s == nil {
		s = DefaultStrategy
	}

	return s(err, uint(n-1))
}

// Sleep marks the most recent attempt as a failure and pauses the current
// goroutine until the wait duration has elapsed.
//
// It sleeps until the duration elapses or ctx is canceled, whichever is first.
// If ctx is canceled before the duration elapses it returns ctx.Err(),
// otherwise it returns nil.
//
// err is the error describing the operation's failure condition, if known. A
// nil error does not indicate a success.
func (c *Counter) Sleep(ctx context.Context, err error) error {
	return linger.Sleep(ctx, c.Fail(err))
}
