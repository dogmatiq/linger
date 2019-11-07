package linger_test

import (
	"context"
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("func Sleep()", func() {
	It("sleeps for the first positive duration", func() {
		start := time.Now()
		err := Sleep(context.Background(), 0*time.Second, -1*time.Second, 10*time.Millisecond)
		stop := time.Now()
		elapsed := stop.Sub(start)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(elapsed).To(BeNumerically(">=", 10*time.Millisecond))
	})

	It("returns 'immediately' if none of the durations are positive", func() {
		start := time.Now()
		err := Sleep(context.Background(), 0*time.Second, -1*time.Second)
		stop := time.Now()
		elapsed := stop.Sub(start)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(elapsed).To(BeNumerically("<", 5*time.Millisecond))
	})

	It("returns early if the context is canceled", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		start := time.Now()
		err := Sleep(ctx, 100*time.Millisecond)
		stop := time.Now()
		elapsed := stop.Sub(start)

		Expect(err).To(Equal(context.DeadlineExceeded))
		Expect(elapsed).To(BeNumerically(">=", 10*time.Millisecond))
		Expect(elapsed).To(BeNumerically("<", 100*time.Millisecond))
	})
})

var _ = Describe("func SleepX()", func() {
	It("applies the transform", func() {
		x := func(t time.Duration) time.Duration {
			return t / 2
		}

		start := time.Now()
		err := SleepX(context.Background(), x, 20*time.Millisecond)
		stop := time.Now()
		elapsed := stop.Sub(start)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(elapsed).To(BeNumerically(">=", 10*time.Millisecond))
		Expect(elapsed).To(BeNumerically("<", 20*time.Millisecond))
	})
})

var _ = Describe("func SleepUntil()", func() {
	It("sleeps until earliest time", func() {
		start := time.Now()
		err := SleepUntil(context.Background(), start.Add(20*time.Millisecond), start.Add(10*time.Millisecond))
		stop := time.Now()
		elapsed := stop.Sub(start)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(elapsed).To(BeNumerically(">=", 10*time.Millisecond))
		Expect(elapsed).To(BeNumerically("<", 20*time.Millisecond))
	})

	It("returns immediately if all of the times are in the past", func() {
		start := time.Now()
		err := SleepUntil(context.Background(), start.Add(-20*time.Millisecond), start.Add(-10*time.Millisecond))
		stop := time.Now()
		elapsed := stop.Sub(start)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(elapsed).To(BeNumerically("<", 5*time.Millisecond))
	})

	It("returns early if the context is canceled", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancel()

		start := time.Now()
		err := SleepUntil(ctx, start.Add(100*time.Millisecond))
		stop := time.Now()
		elapsed := stop.Sub(start)

		Expect(err).To(Equal(context.DeadlineExceeded))
		Expect(elapsed).To(BeNumerically(">=", 10*time.Millisecond))
		Expect(elapsed).To(BeNumerically("<", 100*time.Millisecond))
	})
})

var _ = Describe("func SleepUntilX()", func() {
	It("applies the transform", func() {
		x := func(t time.Duration) time.Duration {
			return t / 2
		}

		start := time.Now()
		err := SleepUntilX(context.Background(), x, start.Add(20*time.Millisecond))
		stop := time.Now()
		elapsed := stop.Sub(start)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(elapsed).To(BeNumerically(">=", 10*time.Millisecond))
		Expect(elapsed).To(BeNumerically("<", 20*time.Millisecond))
	})
})
