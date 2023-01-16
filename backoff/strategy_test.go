package backoff_test

import (
	"math"
	"time"

	"github.com/dogmatiq/linger"
	. "github.com/dogmatiq/linger/backoff"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Exponential()", func() {
	It("returns a strategy that backs-off exponentially", func() {
		strategy := Exponential(3 * time.Second)

		Expect(strategy(nil, 4)).To(Equal(48 * time.Second))
		Expect(strategy(nil, 5)).To(Equal(96 * time.Second))
	})

	It("panics if the unit is zero", func() {
		Expect(func() {
			Exponential(0)
		}).To(Panic())
	})

	It("panics if the unit is negative", func() {
		Expect(func() {
			Exponential(-1)
		}).To(Panic())
	})

	It("does not overflow the time.Duration type", func() {
		strategy := Exponential(1)

		// No overflow at 2^62. We have 63-bits of positive magnitude available
		// because time.Duration is a signed 64-bit integer.
		Expect(strategy(nil, 62)).To(Equal(time.Duration(math.Pow(2, 62))))

		// Starts overflowing at 2^63.
		Expect(strategy(nil, 63)).To(Equal(linger.MaxDuration))

		// Continues to return the capped value as the exponent increases.
		Expect(strategy(nil, 100)).To(Equal(linger.MaxDuration))
	})

})

var _ = Describe("func Constant()", func() {
	It("returns a strategy that returns a fixed duration", func() {
		strategy := Constant(3 * time.Second)

		Expect(strategy(nil, 4)).To(Equal(3 * time.Second))
		Expect(strategy(nil, 5)).To(Equal(3 * time.Second))
	})
})

var _ = Describe("func Linear()", func() {
	It("returns a strategy that returns a linearly increasing duration", func() {
		strategy := Linear(3 * time.Second)

		Expect(strategy(nil, 4)).To(Equal(15 * time.Second))
		Expect(strategy(nil, 5)).To(Equal(18 * time.Second))
	})

	It("panics if the unit is zero", func() {
		Expect(func() {
			Linear(0)
		}).To(Panic())
	})

	It("panics if the unit is negative", func() {
		Expect(func() {
			Linear(-1)
		}).To(Panic())
	})

	It("does not overflow the time.Duration type", func() {
		unit := linger.MaxDuration - 1

		strategy := Linear(unit)

		// No overflow at 1 * unit, as it's slightly below the max.
		Expect(strategy(nil, 0)).To(Equal(unit))

		// Start overflowing at 2 * unit. Overflow is detected because the
		// result of the multiplication wraps negative.
		Expect(unit * 2).To(BeNumerically("<", 0)) // verify test inputs actually wrap
		Expect(strategy(nil, 1)).To(Equal(linger.MaxDuration))

		// Continue overflowing. Overflow is detected because the result of the
		// multiplication wraps negative, and continues on to be positive again.
		Expect(unit * 5).To(BeNumerically(">", 0)) // verify test inputs actually wrap
		Expect(strategy(nil, 4)).To(Equal(linger.MaxDuration))
	})
})

var _ = Describe("func WithTransform()", func() {
	It("returns a strategy that that transforms the result of the input strategy", func() {
		s := WithTransforms(
			Linear(10*time.Second),
			linger.Limiter(15*time.Second, linger.MaxDuration),
			linger.Limiter(0, 25*time.Second),
		)

		Expect(s(nil, 0)).To(Equal(15 * time.Second))
		Expect(s(nil, 1)).To(Equal(20 * time.Second))
		Expect(s(nil, 2)).To(Equal(25 * time.Second))
	})
})

var _ = Describe("func CoalesceStrategy()", func() {
	It("returns a strategy that yields the first positive duration", func() {

		stgyOne := func(_ error, n uint) time.Duration {
			if n == 1 {
				return 6 * time.Second
			}
			return 0
		}

		stgyTwo := func(_ error, n uint) time.Duration {
			if n == 2 {
				return 9 * time.Second
			}
			return 0
		}

		s := CoalesceStrategy(
			stgyOne,
			stgyTwo,
		)

		sC := CoalesceStrategy(
			s,
			Constant(3*time.Second),
		)

		Expect(s(nil, 0)).To(Equal(0 * time.Second))
		Expect(s(nil, 1)).To(Equal(6 * time.Second))
		Expect(s(nil, 2)).To(Equal(9 * time.Second))

		Expect(sC(nil, 0)).To(Equal(3 * time.Second))
		Expect(sC(nil, 1)).To(Equal(6 * time.Second))
		Expect(sC(nil, 2)).To(Equal(9 * time.Second))

	})
})
