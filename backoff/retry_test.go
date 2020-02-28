package backoff_test

import (
	"context"
	"errors"
	"time"

	. "github.com/dogmatiq/linger/backoff"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Retry()", func() {
	It("returns nil if the function succeeds", func() {
		count := 0

		n, err := Retry(
			context.Background(),
			Constant(1*time.Nanosecond),
			func(context.Context) error {
				if count == 3 {
					return nil
				}

				count++
				return errors.New("<error>")
			},
		)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(n).To(BeNumerically("==", 3))
	})

	It("returns an error if the context is canceled", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		n, err := Retry(
			ctx,
			Linear(10*time.Millisecond),
			func(context.Context) error {
				return errors.New("<error>")
			},
		)

		Expect(err).To(Equal(context.DeadlineExceeded))
		Expect(n).To(BeNumerically("==", 1))
	})
})
