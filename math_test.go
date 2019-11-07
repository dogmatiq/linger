package linger_test

import (
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
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

var _ = DescribeTable(
	"func Limit()",
	func(in, out time.Duration) {
		a := 20 * time.Second
		b := 30 * time.Second
		Expect(Limit(in, a, b)).To(Equal(out))
		Expect(Limit(in, b, a)).To(Equal(out))
	},
	Entry("limit to min", 15*time.Second, 20*time.Second),
	Entry("limit to max", 35*time.Second, 30*time.Second),
	Entry("do not limit", 25*time.Second, 25*time.Second),
)

var _ = DescribeTable(
	"func LimitT()",
	func(in, out time.Duration) {
		now := time.Now()

		a := now.Add(20 * time.Second)
		b := now.Add(30 * time.Second)
		inT := now.Add(in)
		outT := now.Add(out)

		Expect(LimitT(inT, a, b)).To(Equal(outT))
		Expect(LimitT(inT, b, a)).To(Equal(outT))
	},
	Entry("limit to min", 15*time.Second, 20*time.Second),
	Entry("limit to max", 35*time.Second, 30*time.Second),
	Entry("do not limit", 25*time.Second, 25*time.Second),
)
