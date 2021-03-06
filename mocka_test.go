package mocka

import (
	"errors"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("mocka", func() {
	Describe("Function", func() {
		var (
			callCount        int
			fn               func(str string, num int) (int, error)
			failTestReporter *mockTestReporter
		)

		BeforeEach(func() {
			failTestReporter = &mockTestReporter{}
			callCount = 0
			fn = func(str string, num int) (int, error) {
				callCount++
				return len(str) + num, nil
			}
		})

		It("reports an error if passed a nil as the function pointer", func() {
			stub := Function(failTestReporter, nil)

			Expect(stub).To(BeNil())
			Expect(failTestReporter.messages).To(Equal([]string{
				"mocka: expected the second argument to be a pointer to a function, but received a nil",
			}))
		})

		It("reports an error if a non-pointer value is passed as the function pointer", func() {
			stub := Function(failTestReporter, 42)

			Expect(stub).To(BeNil())
			Expect(failTestReporter.messages).To(Equal([]string{
				"mocka: expected the second argument to be a pointer to a function, but received a int",
			}))
		})

		It("reports an error if a non-function value is passed as the function pointer", func() {
			num := 42
			stub := Function(failTestReporter, &num)

			Expect(stub).To(BeNil())
			Expect(failTestReporter.messages).To(Equal([]string{
				"mocka: expected the second argument to be a pointer to a function, but received a pointer to a int",
			}))
		})

		It("reports an error supplied out parameters are not of the same type", func() {
			stub := Function(failTestReporter, &fn, "42", nil)

			Expect(stub).To(BeNil())
			Expect(failTestReporter.messages).To(Equal([]string{
				"mocka: expected return values of type (int, error), but received (string, <nil>)",
			}))
		})

		It("reports an error if cloneValue returns an error", func() {
			_cloneValue = func(interface{}, interface{}) error {
				return errors.New("Ope")
			}
			defer func() {
				_cloneValue = cloneValue
			}()

			stub := Function(failTestReporter, &fn, 42, nil)

			Expect(stub).To(BeNil())
			Expect(failTestReporter.messages).To(Equal([]string{
				"mocka: could not clone function pointer to new memory address: Ope",
			}))
		})

		It("returns a stub with a reference to the original function", func() {
			stub := Function(GinkgoT(), &fn, 42, nil)

			Expect(stub).ToNot(BeNil())

			Expect(stub.originalFunc).ToNot(BeNil())

			_, _ = stub.originalFunc.(func(str string, num int) (int, error))("", 0)

			Expect(callCount).To(Equal(1))
		})

		It("returns a stub with properties initialized with zero values", func() {
			stub := Function(GinkgoT(), &fn, 42, nil)

			Expect(stub).ToNot(BeNil())
			Expect(stub.calls).To(BeNil())
			Expect(stub.customArgs).To(BeNil())
		})

		It("returns a stub with outParameters as supplied", func() {
			stub := Function(GinkgoT(), &fn, 42, nil)

			Expect(stub).ToNot(BeNil())
			Expect(stub.outParameters).To(Equal([]interface{}{42, nil}))
		})
	})

	Describe("CreateSandbox", func() {
		It("returns a sandbox with stub assigned as nil", func() {
			s := CreateSandbox(GinkgoT())

			Expect(s).ToNot(BeNil())
			Expect(s.stubs).To(BeNil())
		})
	})

	Describe("ensureTestReporter", func() {
		It("calls exit if testReporter is nil", func() {
			var message string
			exitFn := func(args ...interface{}) {
				Expect(args).To(HaveLen(1))
				msg, ok := args[0].(string)
				Expect(ok).To(BeTrue())
				message = msg
			}

			actual := ensureTestReporter(nil, exitFn)

			Expect(actual).To(BeNil())
			Expect(message).To(Equal("mocka: test reporter required to fail tests"))
		})

		It("returns the provided testReporter", func() {
			expected := GinkgoT()

			actual := ensureTestReporter(expected, log.Fatal)

			Expect(actual).To(Equal(expected))
		})
	})
})
