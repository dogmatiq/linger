package linger_test

import (
	"context"
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func FromContextDeadline()", func() {
	It("returns the time until the deadline of the context", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		d, ok := FromContextDeadline(ctx)
		Expect(ok).To(BeTrue())
		Expect(d).To(BeNumerically("~", 10*time.Second, 1*time.Second))
	})

	It("returns false if the context does not have a deadline", func() {
		_, ok := FromContextDeadline(context.Background())
		Expect(ok).To(BeFalse())
	})
})

var _ = Describe("func ContextWithTimeout()", func() {
	It("sets a timeout for the first positive duration", func() {
		expect := time.Now().Add(10 * time.Second)
		ctx, cancel := ContextWithTimeout(context.Background(), 0*time.Second, -1*time.Second, 10*time.Second)
		defer cancel()

		dl, ok := ctx.Deadline()
		Expect(ok).To(BeTrue())
		Expect(dl).To(BeTemporally("~", expect))
	})

	It("times out 'immediately' if none of the durations are positive", func() {
		ctx, cancel := ContextWithTimeout(context.Background(), 0*time.Second, -1*time.Second)
		defer cancel()

		err := ctx.Err()
		Expect(err).To(Equal(context.DeadlineExceeded))
	})
})

var _ = Describe("func ContextWithTimeoutX()", func() {
	It("applies the transform", func() {
		x := func(t time.Duration) time.Duration {
			return t / 2
		}

		expect := time.Now().Add(10 * time.Second)
		ctx, cancel := ContextWithTimeoutX(context.Background(), x, 20*time.Second)
		defer cancel()

		dl, ok := ctx.Deadline()
		Expect(ok).To(BeTrue())
		Expect(dl).To(BeTemporally("~", expect))
	})
})
