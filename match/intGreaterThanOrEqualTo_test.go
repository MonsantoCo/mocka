package match

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("intGreaterThanOrEqualTo", func() {
	Describe("IntGreaterThanOrEqualTo", func() {
		It("returns an intGreaterThanOrEqualTo struct", func() {
			actual := IntGreaterThanOrEqualTo(10)

			Expect(actual).To(BeAssignableToTypeOf(new(intGreaterThanOrEqualTo)))
		})
	})

	Describe("SupportedKinds", func() {
		It("returns all support kinds in go", func() {
			actual := IntGreaterThanOrEqualTo(5).SupportedKinds()

			Expect(actual).To(Equal(
				map[reflect.Kind]struct{}{
					reflect.Int:   {},
					reflect.Int8:  {},
					reflect.Int16: {},
					reflect.Int32: {},
					reflect.Int64: {},
				}))
		})
	})

	DescribeTable("Match returns true",
		func(expected int64, actual interface{}) {
			Expect(IntGreaterThanOrEqualTo(expected).Match(actual)).To(BeTrue())
		},
		Entry("with int", int64(5), int(10)),
		Entry("with int8", int64(10), int8(18)),
		Entry("with int16", int64(15), int16(22)),
		Entry("with int32", int64(20), int32(40)),
		Entry("with int64", int64(8), int64(15)),
		Entry("when actual(int) is the same as the expected", int64(5), int(5)),
		Entry("when actual(int8) is the same as the expected", int64(10), int8(10)),
		Entry("when actual(int16) is the same as the expected", int64(15), int16(15)),
		Entry("when actual(int32) is the same as the expected", int64(20), int32(20)),
		Entry("when actual(int64) is the same as the expected", int64(8), int64(8)),
	)

	DescribeTable("Match returns false",
		func(expected int64, actual interface{}) {
			Expect(IntGreaterThanOrEqualTo(expected).Match(actual)).To(BeFalse())
		},
		Entry("when actual is nil", int64(5), nil),
		Entry("when actual(int) is less than expected", int64(5), int(1)),
		Entry("when actual(int8) is less than expected", int64(10), int8(8)),
		Entry("when actual(int16) is less than expected", int64(15), int16(2)),
		Entry("when actual(int32) is less than expected", int64(20), int32(4)),
		Entry("when actual(int64) is less than expected", int64(8), int64(5)),
		Entry("when actual is not an int", int64(10), "10"),
	)
})
