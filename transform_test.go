package linger_test

import (
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Identity()", func() {
	It("returns the input duration", func() {
		d := Identity(100 * time.Second)
		Expect(d).To(Equal(100 * time.Second))
	})
})
