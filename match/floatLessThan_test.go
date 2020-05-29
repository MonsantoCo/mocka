package match

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("floatLessThan", func() {
	Describe("FloatLessThan", func() {
		It("returns an floatLessThan struct", func() {
			actual := FloatLessThan(10)

			Expect(actual).To(BeAssignableToTypeOf(new(floatLessThan)))
		})
	})

	Describe("SupportedKinds", func() {
		It("returns all support kinds in go", func() {
			actual := FloatLessThan(5).SupportedKinds()

			Expect(actual).To(Equal(
				map[reflect.Kind]struct{}{
					reflect.Float32: {},
					reflect.Float64: {},
				}))
		})
	})

	DescribeTable("Match returns true",
		func(expected float64, actual interface{}) {
			Expect(FloatLessThan(expected).Match(actual)).To(BeTrue())
		},
		Entry("with float32", float64(40), float32(20)),
		Entry("with float64", float64(15), float64(8)),
	)

	DescribeTable("Match returns false",
		func(expected float64, actual interface{}) {
			Expect(FloatLessThan(expected).Match(actual)).To(BeFalse())
		},
		Entry("when actual is nil", float64(5), nil),
		Entry("when actual(float32) is greater than expected", float64(4), float32(20)),
		Entry("when actual(float64) is greater than expected", float64(5), float64(8)),
		Entry("when actual(float32) is the same as the expected", float64(20), float32(20)),
		Entry("when actual(float64) is the same as the expected", float64(8), float64(8)),
		Entry("when actual is not an int", float64(10), "10"),
	)
})
