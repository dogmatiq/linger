package backoff

import (
	"time"

	"github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Exponential()", func() {
	It("returns a strategy that backs-off exponentially", func() {
		strategy := Exponential(3 * time.Second)

		Expect(strategy(5, nil)).To(Equal(48 * time.Second))
		Expect(strategy(6, nil)).To(Equal(96 * time.Second))
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
})

var _ = Describe("func Constant()", func() {
	It("returns a strategy that returns a fixed duration", func() {
		strategy := Constant(3 * time.Second)

		Expect(strategy(5, nil)).To(Equal(3 * time.Second))
		Expect(strategy(6, nil)).To(Equal(3 * time.Second))
	})
})

var _ = Describe("func Linear()", func() {
	It("returns a strategy that returns a fixed duration", func() {
		strategy := Linear(3 * time.Second)

		Expect(strategy(5, nil)).To(Equal(15 * time.Second))
		Expect(strategy(6, nil)).To(Equal(18 * time.Second))
	})
})

var _ = Describe("func WithTransform()", func() {
	It("returns a strategy that that transforms the result of the input strategy", func() {
		s := WithTransforms(
			Linear(10*time.Second),
			linger.Limiter(15*time.Second, linger.MaxDuration),
			linger.Limiter(0, 25*time.Second),
		)

		Expect(s(1, nil)).To(Equal(15 * time.Second))
		Expect(s(2, nil)).To(Equal(20 * time.Second))
		Expect(s(3, nil)).To(Equal(25 * time.Second))
	})
})
