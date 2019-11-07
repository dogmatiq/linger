package linger

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Rand()", func() {
	It("returns a value between min and max", func() {
		d := Rand(1*time.Second, 100*time.Second)
		Expect(d).To(BeNumerically(">=", 1*time.Second))
		Expect(d).To(BeNumerically("<=", 100*time.Second))
	})
})
