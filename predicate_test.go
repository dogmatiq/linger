package linger_test

import (
	"time"

	. "github.com/dogmatiq/linger"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable(
	"type DurationPredicate",
	func(p DurationPredicate, v time.Duration, expect bool) {
		Expect(p(v)).To(Equal(expect))
	},
	Entry("non-zero with zero value", NonZero, 0*time.Second, false),
	Entry("non-zero with non-zero value", NonZero, 1*time.Second, true),
	Entry("positive with zero value", Positive, 0*time.Second, false),
	Entry("positive with positive value", Positive, +1*time.Second, true),
	Entry("positive with negative value", Positive, -1*time.Second, false),
	Entry("negative with zero value", Negative, 0*time.Second, false),
	Entry("negative with positive value", Negative, +1*time.Second, false),
	Entry("negative with negative value", Negative, -1*time.Second, true),
)

var _ = DescribeTable(
	"type TimePredicate",
	func(p TimePredicate, v time.Time, expect bool) {
		Expect(p(v)).To(Equal(expect))
	},
	Entry("non-zero with zero value", NonZeroT, time.Time{}, false),
	Entry("non-zero with non-zero value", NonZeroT, time.Now(), true),
)
