package examples

import (
	"fmt"
	"log"

	"github.com/MonsantoCo/mocka"
)

func ExampleFile() {
	f := mocka.File("file_name", "This is the body")
	b := make([]byte, 16)

	f.Read(b)

	fmt.Println(string(b))
	// Output: This is the body
}

func ExampleFunction() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fmt.Println(fn("1"))
	// Output: 20
}

func ExampleCreateSandbox() {
	var fn = func(str string) int {
		return len(str)
	}

	sandbox := mocka.CreateSandbox()
	defer sandbox.Restore()

	sandbox.StubFunction(&fn, 20)

	fmt.Println(fn("1"))
	// Output: 20
}

func ExampleCall_Arguments() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fn("123")

	fmt.Println(stub.GetFirstCall().Arguments())
	// Output: [123]
}

func ExampleCall_ReturnValues() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fn("123")

	fmt.Println(stub.GetFirstCall().ReturnValues())
	// Output: [20]
}

func ExampleOnCallReturner_Return() {
	var fn = func(str []string, n int) int {
		return len(str) + n
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	if err = stub.WithArgs([]string{"123", "456"}, 2).Return(5); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn([]string{"123", "456"}, 2))
	// Output: 5
}

func ExampleOnCallReturner_OnCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	withArgs123 := stub.WithArgs("123")

	if err = withArgs123.OnCall(1).Return(5); err != nil {
		log.Fatal(err.Error())
	}

	if err = withArgs123.OnCall(3).Return(3); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	// Output: 20
	// 5
	// 20
	// 3
}

func ExampleOnCallReturner_OnFirstCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	withArgs123 := stub.WithArgs("123")

	if err = withArgs123.OnFirstCall().Return(5); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	// Output: 5
	// 20
	// 20
}

func ExampleOnCallReturner_OnSecondCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	withArgs123 := stub.WithArgs("123")

	if err = withArgs123.OnSecondCall().Return(5); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	// Output: 20
	// 5
	// 20
}

func ExampleOnCallReturner_OnThirdCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	withArgs123 := stub.WithArgs("123")

	if err = withArgs123.OnThirdCall().Return(5); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	// Output: 20
	// 20
	// 5
}

func ExampleStub_Return() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fmt.Println(fn("123"))

	if err = stub.Return(5); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("123"))
	// Output: 20
	// 5
}

func ExampleStub_OnCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	if err = stub.OnCall(1).Return(5); err != nil {
		log.Fatal(err.Error())
	}

	if err = stub.OnCall(3).Return(3); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("asdf"))
	fmt.Println(fn("234"))
	fmt.Println(fn("12gsbs3"))
	fmt.Println(fn("adf"))
	// Output: 20
	// 5
	// 20
	// 3
}

func ExampleStub_OnFirstCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	if err = stub.OnFirstCall().Return(5); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	// Output: 5
	// 20
	// 20
}

func ExampleStub_OnSecondCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	if err = stub.OnSecondCall().Return(5); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	// Output: 20
	// 5
	// 20
}

func ExampleStub_OnThirdCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	if err = stub.OnThirdCall().Return(5); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	fmt.Println(fn("123"))
	// Output: 20
	// 20
	// 5
}

func ExampleStub_Restore() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(fn("123"))

	stub.Restore()

	fmt.Println(fn("123"))
	// Output: 20
	// 3
}

func ExampleStub_ExecOnCall() {
	var fn = func(in <-chan int) <-chan int {
		out := make(chan int, 1)
		go func() {
			out <- <-in
		}()
		return out
	}

	out := make(chan int, 1)
	stub, err := mocka.Function(&fn, out)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	stub.ExecOnCall(func(args []interface{}) {
		c := args[0].(<-chan int)
		out <- <-c
	})

	in := make(chan int, 1)
	in <- 10
	o := fn(in)
	fmt.Println(<-o)
	// Output: 10
}

func ExampleStub_GetCalls() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fn("first call")
	fn("second call")
	fn("third call")

	calls := stub.GetCalls()
	for _, call := range calls {
		fmt.Printf("Arguments: %+v; Return Values: %+v\n", call.Arguments(), call.ReturnValues())
	}
	// Output: Arguments: [first call]; Return Values: [20]
	// Arguments: [second call]; Return Values: [20]
	// Arguments: [third call]; Return Values: [20]
}

func ExampleStub_GetCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fn("first call")
	fn("second call")
	fn("third call")

	call := stub.GetCall(2)
	fmt.Printf("Arguments: %+v; Return Values: %+v\n", call.Arguments(), call.ReturnValues())
	// Output: Arguments: [third call]; Return Values: [20]
}

func ExampleStub_GetFirstCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fn("first call")
	fn("second call")
	fn("third call")

	call := stub.GetFirstCall()
	fmt.Printf("Arguments: %+v; Return Values: %+v\n", call.Arguments(), call.ReturnValues())
	// Output: Arguments: [first call]; Return Values: [20]
}

func ExampleStub_GetSecondCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fn("first call")
	fn("second call")
	fn("third call")

	call := stub.GetSecondCall()
	fmt.Printf("Arguments: %+v; Return Values: %+v\n", call.Arguments(), call.ReturnValues())
	// Output: Arguments: [second call]; Return Values: [20]
}

func ExampleStub_GetThirdCall() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fn("first call")
	fn("second call")
	fn("third call")

	call := stub.GetThirdCall()
	fmt.Printf("Arguments: %+v; Return Values: %+v\n", call.Arguments(), call.ReturnValues())
	// Output: Arguments: [third call]; Return Values: [20]
}

func ExampleStub_CallCount() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fn("first call")
	fn("second call")
	fn("third call")

	fmt.Println(stub.CallCount())
	// Output: 3
}

func ExampleStub_CalledOnce() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fmt.Println(stub.CalledOnce())

	fn("first call")
	fn("second call")
	fn("third call")

	fmt.Println(stub.CalledOnce())
	// Output: false
	// true
}

func ExampleStub_CalledTwice() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fmt.Println(stub.CalledTwice())

	fn("first call")
	fn("second call")
	fn("third call")

	fmt.Println(stub.CalledTwice())
	// Output: false
	// true
}

func ExampleStub_CalledThrice() {
	var fn = func(str string) int {
		return len(str)
	}

	stub, err := mocka.Function(&fn, 20)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stub.Restore()

	fmt.Println(stub.CalledThrice())

	fn("first call")
	fn("second call")
	fn("third call")

	fmt.Println(stub.CalledThrice())
	// Output: false
	// true
}