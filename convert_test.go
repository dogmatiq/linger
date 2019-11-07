package linger_test

import (
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func FromSeconds()", func() {
	It("returns a duration representing the given number of seconds", func() {
		d := FromSeconds(120.0)
		Expect(d).To(Equal(2 * time.Minute))
	})
})
