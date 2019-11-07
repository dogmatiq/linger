package linger_test

import (
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func First()", func() {
	It("returns the first value that matches the predicate", func() {
		v, ok := First(Positive, 0*time.Second, -1*time.Second, 1*time.Second, 2*time.Second)
		Expect(ok).To(BeTrue())
		Expect(v).To(Equal(1 * time.Second))
	})

	It("returns false if no values match", func() {
		_, ok := First(Positive, 0*time.Second, -1*time.Second)
		Expect(ok).To(BeFalse())
	})
})

var _ = Describe("func MustFirst()", func() {
	It("returns the first value that matches the predicate", func() {
		v := MustFirst(Positive, 0*time.Second, -1*time.Second, 1*time.Second, 2*time.Second)
		Expect(v).To(Equal(1 * time.Second))
	})

	It("panics if no values match", func() {
		Expect(func() {
			MustFirst(Positive, 0*time.Second, -1*time.Second)
		}).To(Panic())
	})
})

var _ = Describe("func FirstT()", func() {
	epoch := time.Unix(0, 0)
	now := time.Now()

	It("returns the first value that matches the predicate", func() {
		v, ok := FirstT(NonZeroT, time.Time{}, epoch, now)
		Expect(ok).To(BeTrue())
		Expect(v).To(BeTemporally("==", epoch))
	})

	It("returns false if no values match", func() {
		_, ok := FirstT(NonZeroT, time.Time{})
		Expect(ok).To(BeFalse())
	})
})

var _ = Describe("func MustFirstT()", func() {
	epoch := time.Unix(0, 0)
	now := time.Now()

	It("returns the first value that matches the predicate", func() {
		v := MustFirstT(NonZeroT, time.Time{}, epoch, now)
		Expect(v).To(BeTemporally("==", epoch))
	})

	It("panics if no values match", func() {
		Expect(func() {
			MustFirstT(NonZeroT, time.Time{})
		}).To(Panic())
	})
})
