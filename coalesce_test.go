package linger_test

import (
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Coalesce()", func() {
	It("returns the first positive value", func() {
		v, ok := Coalesce(0*time.Second, -1*time.Second, 1*time.Second, 2*time.Second)
		Expect(ok).To(BeTrue())
		Expect(v).To(Equal(1 * time.Second))
	})

	It("returns the zero-value and false if no values are positive", func() {
		v, ok := Coalesce(0*time.Second, -1*time.Second)
		Expect(ok).To(BeFalse())
		Expect(v).To(Equal(0 * time.Second))
	})
})

var _ = Describe("func MustCoalesce()", func() {
	It("returns the first positive value", func() {
		v := MustCoalesce(0*time.Second, -1*time.Second, 1*time.Second, 2*time.Second)
		Expect(v).To(Equal(1 * time.Second))
	})

	It("panics if no values are positive", func() {
		Expect(func() {
			MustCoalesce(0*time.Second, -1*time.Second)
		}).To(Panic())
	})
})

var _ = Describe("func CoalesceT()", func() {
	epoch := time.Unix(0, 0)
	now := time.Now()

	It("returns the first non-zero value", func() {
		v, ok := CoalesceT(time.Time{}, epoch, now)
		Expect(ok).To(BeTrue())
		Expect(v).To(BeTemporally("==", epoch))
	})

	It("returns the zero-value and false if no values are non-zero", func() {
		v, ok := CoalesceT(time.Time{})
		Expect(ok).To(BeFalse())
		Expect(v).To(Equal(time.Time{}))
	})
})

var _ = Describe("func MustCoalesceT()", func() {
	epoch := time.Unix(0, 0)
	now := time.Now()

	It("returns the first non-zero value", func() {
		v := MustCoalesceT(time.Time{}, epoch, now)
		Expect(v).To(BeTemporally("==", epoch))
	})

	It("panics if no values are non-zero", func() {
		Expect(func() {
			MustCoalesceT(time.Time{})
		}).To(Panic())
	})
})

var _ = Describe("func Coalescer()", func() {
	It("returns the first positive value", func() {
		v := Coalescer(-1*time.Second, 1*time.Second)(0 * time.Second)
		Expect(v).To(Equal(1 * time.Second))
	})

	It("returns the input value if it is positive", func() {
		v := Coalescer(-1*time.Second, 1*time.Second)(5 * time.Second)
		Expect(v).To(Equal(5 * time.Second))
	})

	It("panics if no values are positive", func() {
		Expect(func() {
			Coalescer(-1 * time.Second)(0 * time.Second)
		}).To(Panic())
	})
})
