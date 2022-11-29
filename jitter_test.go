package linger_test

import (
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func FullJitter()", func() {
	It("returns a value between 0 and d", func() {
		d := FullJitter(100 * time.Second)
		Expect(d).To(BeNumerically(">=", 0*time.Second))
		Expect(d).To(BeNumerically("<=", 100*time.Second))
	})
})

var _ = Describe("func ProportionalJitter()", func() {
	It("adds to the input duration when the proportion is positive", func() {
		d := ProportionalJitter(0.25)(100 * time.Second)
		Expect(d).To(BeNumerically(">=", 100*time.Second))
		Expect(d).To(BeNumerically("<=", 125*time.Second))
	})

	It("subtracts from the input duration when the proportion is negative", func() {
		d := ProportionalJitter(-0.25)(100 * time.Second)
		Expect(d).To(BeNumerically(">=", 75*time.Second))
		Expect(d).To(BeNumerically("<=", 100*time.Second))
	})
})

var _ = Describe("func Rand()", func() {
	It("returns a value between the given values", func() {
		d := Rand(1*time.Second, 100*time.Second)
		Expect(d).To(BeNumerically(">=", 1*time.Second))
		Expect(d).To(BeNumerically("<=", 100*time.Second))
	})

	It("does not require the arguments in any particular order", func() {
		d := Rand(100*time.Second, 1*time.Second)
		Expect(d).To(BeNumerically(">=", 1*time.Second))
		Expect(d).To(BeNumerically("<=", 100*time.Second))
	})

	It("returns the value when the arguments are equal", func() {
		d := Rand(100*time.Second, 100*time.Second)
		Expect(d).To(Equal(100 * time.Second))
	})

	// See https://github.com/dogmatiq/linger/issues/18
	It("supports MaxDuration", func() {
		d := Rand(0, MaxDuration)
		Expect(d).To(BeNumerically(">=", 0))
		Expect(d).To(BeNumerically("<=", MaxDuration))

		d = Rand(MaxDuration, 0)
		Expect(d).To(BeNumerically(">=", 0))
		Expect(d).To(BeNumerically("<=", MaxDuration))

		d = Rand(MaxDuration, MaxDuration)
		Expect(d).To(Equal(MaxDuration))
	})
})
