package linger

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Backoff", func() {
	var (
		config  *BackoffConfig
		backoff *Backoff
	)

	BeforeEach(func() {
		config = &BackoffConfig{
			Strategy:  LinearBackoff(10 * time.Millisecond),
			Transform: Identity,
		}

		backoff = &Backoff{
			Config: config,
		}
	})

	Describe("func Ok()", func() {
		It("resets the failure count", func() {
			backoff.Fail(nil)
			backoff.Fail(nil)
			backoff.Ok()

			Expect(backoff.Fail(nil)).To(Equal(10 * time.Millisecond))
		})
	})

	Describe("func Fail()", func() {
		It("uses the default configuration", func() {
			backoff.Config = nil

			// The default configuration uses FullJitter, so this is hard to
			// test well, but essentially we're ensuring it doesn't panic.

			Expect(backoff.Fail(nil)).To(BeNumerically("<=", 3*time.Second))
			Expect(backoff.Fail(nil)).To(BeNumerically("<=", 6*time.Second))
			Expect(backoff.Fail(nil)).To(BeNumerically("<=", 12*time.Second))
		})

		It("uses the backoff strategy", func() {
			Expect(backoff.Fail(nil)).To(Equal(10 * time.Millisecond))
			Expect(backoff.Fail(nil)).To(Equal(20 * time.Millisecond))
		})

		It("uses the default backoff strategy if none is specified", func() {
			config.Strategy = nil

			Expect(backoff.Fail(nil)).To(Equal(3 * time.Second))
			Expect(backoff.Fail(nil)).To(Equal(6 * time.Second))
			Expect(backoff.Fail(nil)).To(Equal(12 * time.Second))
		})

		It("applies the transform", func() {
			config.Transform = func(d time.Duration) time.Duration {
				return d * 2
			}

			Expect(backoff.Fail(nil)).To(Equal(20 * time.Millisecond))
			Expect(backoff.Fail(nil)).To(Equal(40 * time.Millisecond))
		})

		It("uses a full-jitter transform by default", func() {
			config.Transform = nil

			for attempts := 0; attempts < 10000; attempts++ {
				backoff.Ok()

				d := backoff.Fail(nil)
				Expect(d).To(BeNumerically(">=", 0*time.Millisecond))
				Expect(d).To(BeNumerically("<=", 10*time.Millisecond))

				if d < 3*time.Second {
					// run the test repeatedly until a value less than the
					// fundamental value is returned.
					return
				}
			}

			Fail("exhausted attempts to detect jitter")
		})

		It("respects the lower bound", func() {
			config.Min = 25 * time.Millisecond

			Expect(backoff.Fail(nil)).To(Equal(25 * time.Millisecond))
			Expect(backoff.Fail(nil)).To(Equal(25 * time.Millisecond))
			Expect(backoff.Fail(nil)).To(Equal(30 * time.Millisecond))
		})

		It("respects the upper bound", func() {
			config.Max = 15 * time.Millisecond

			Expect(backoff.Fail(nil)).To(Equal(10 * time.Millisecond))
			Expect(backoff.Fail(nil)).To(Equal(15 * time.Millisecond))
			Expect(backoff.Fail(nil)).To(Equal(15 * time.Millisecond))
		})

		It("uses a default upper bound of one hour", func() {
			config.Strategy = LinearBackoff(25 * time.Minute)

			Expect(backoff.Fail(nil)).To(Equal(25 * time.Minute))
			Expect(backoff.Fail(nil)).To(Equal(50 * time.Minute))
			Expect(backoff.Fail(nil)).To(Equal(1 * time.Hour))
		})

		It("checks the lower bound after the transform", func() {
			config.Min = 25 * time.Millisecond
			config.Transform = func(d time.Duration) time.Duration {
				return d * 2
			}

			Expect(backoff.Fail(nil)).To(Equal(25 * time.Millisecond))
			Expect(backoff.Fail(nil)).To(Equal(40 * time.Millisecond))
		})

		It("checks the upper bound after the transform", func() {
			config.Max = 25 * time.Millisecond
			config.Transform = func(d time.Duration) time.Duration {
				return d * 2
			}

			Expect(backoff.Fail(nil)).To(Equal(20 * time.Millisecond))
			Expect(backoff.Fail(nil)).To(Equal(25 * time.Millisecond))
		})
	})

	Describe("func Sleep()", func() {
		It("sleeps for the computed wait duration", func() {
			start := time.Now()
			err := backoff.Sleep(context.Background(), nil)
			stop := time.Now()
			elapsed := stop.Sub(start)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(elapsed).To(BeNumerically(">=", 10*time.Millisecond))
		})
	})
})

var _ = Describe("func ExponentialBackoff()", func() {
	It("returns a strategy that backs-off exponentially", func() {
		strategy := ExponentialBackoff(3 * time.Second)

		Expect(strategy(5, nil)).To(Equal(48 * time.Second))
		Expect(strategy(6, nil)).To(Equal(96 * time.Second))
	})

	It("panics if the unit is zero", func() {
		Expect(func() {
			ExponentialBackoff(0)
		}).To(Panic())
	})

	It("panics if the unit is negative", func() {
		Expect(func() {
			ExponentialBackoff(-1)
		}).To(Panic())
	})
})

var _ = Describe("func ConstantBackoff()", func() {
	It("returns a strategy that returns a fixed duration", func() {
		strategy := ConstantBackoff(3 * time.Second)

		Expect(strategy(5, nil)).To(Equal(3 * time.Second))
		Expect(strategy(6, nil)).To(Equal(3 * time.Second))
	})
})

var _ = Describe("func LinearBackoff()", func() {
	It("returns a strategy that returns a fixed duration", func() {
		strategy := LinearBackoff(3 * time.Second)

		Expect(strategy(5, nil)).To(Equal(15 * time.Second))
		Expect(strategy(6, nil)).To(Equal(18 * time.Second))
	})
})
