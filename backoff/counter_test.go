package backoff_test

import (
	"context"
	"time"

	. "github.com/dogmatiq/linger/backoff"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Counter", func() {
	var (
		strategy Strategy
		counter  *Counter
	)

	BeforeEach(func() {
		strategy = Linear(10 * time.Millisecond)

		counter = &Counter{
			Strategy: strategy,
		}
	})

	Describe("func Reset()", func() {
		It("resets the failure count", func() {
			counter.Fail(nil)
			counter.Fail(nil)
			counter.Reset()

			Expect(counter.Fail(nil)).To(Equal(10 * time.Millisecond))
		})
	})

	Describe("func Fail()", func() {
		It("uses the default strategy if none is specified", func() {
			counter.Strategy = nil

			// The default strategy uses FullJitter, so this is hard to
			// test well, but essentially we're ensuring it doesn't panic.

			Expect(counter.Fail(nil)).To(BeNumerically("<=", 3*time.Second))
			Expect(counter.Fail(nil)).To(BeNumerically("<=", 6*time.Second))
			Expect(counter.Fail(nil)).To(BeNumerically("<=", 12*time.Second))
		})

		It("uses the specified strategy", func() {
			Expect(counter.Fail(nil)).To(Equal(10 * time.Millisecond))
			Expect(counter.Fail(nil)).To(Equal(20 * time.Millisecond))
		})
	})

	Describe("func Sleep()", func() {
		It("sleeps for the computed wait duration", func() {
			start := time.Now()
			err := counter.Sleep(context.Background(), nil)
			stop := time.Now()
			elapsed := stop.Sub(start)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(elapsed).To(BeNumerically(">=", 10*time.Millisecond))
		})
	})
})
