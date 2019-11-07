package linger_test

import (
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Multiply()", func() {
	It("returns the expected product", func() {
		d := Multiply(10*time.Second, 1.5)
		Expect(d).To(Equal(15 * time.Second))
	})
})

var _ = Describe("func Divide()", func() {
	It("returns the expected quotient", func() {
		d := Divide(15*time.Second, 1.5)
		Expect(d).To(Equal(10 * time.Second))
	})
})

var _ = Describe("func Shortest()", func() {
	It("returns the shortest duration", func() {
		d := Shortest(10*time.Second, 5*time.Second, 15*time.Second)
		Expect(d).To(Equal(5 * time.Second))
	})

	It("returns MaxDuration if no durations are supplied", func() {
		d := Shortest()
		Expect(d).To(Equal(MaxDuration))
	})
})

var _ = Describe("func Longest()", func() {
	It("returns the longest duration", func() {
		d := Longest(10*time.Second, 5*time.Second, 15*time.Second)
		Expect(d).To(Equal(15 * time.Second))
	})

	It("returns MinDuration if no durations are supplied", func() {
		d := Longest()
		Expect(d).To(Equal(MinDuration))
	})
})

var _ = Describe("func Earliest()", func() {
	It("returns the earliest time", func() {
		now := time.Now()
		t1 := now.Add(10 * time.Second)
		t2 := now.Add(5 * time.Second)
		t3 := now.Add(15 * time.Second)

		t := Earliest(t1, t2, t3)
		Expect(t).To(Equal(t2))
	})

	It("eturns the zero-value if no times are supplied", func() {
		t := Earliest()
		Expect(t).To(Equal(time.Time{}))
	})
})

var _ = Describe("func Latest()", func() {
	It("returns the latest time", func() {
		now := time.Now()
		t1 := now.Add(10 * time.Second)
		t2 := now.Add(5 * time.Second)
		t3 := now.Add(15 * time.Second)

		t := Latest(t1, t2, t3)
		Expect(t).To(Equal(t3))
	})

	It("returns the zero-value if no times are supplied", func() {
		t := Latest()
		Expect(t).To(Equal(time.Time{}))
	})
})
