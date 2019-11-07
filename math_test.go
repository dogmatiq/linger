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
