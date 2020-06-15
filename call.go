package mocka

// Call describes a specific call the to the stub function
type Call interface {
	Arguments() []interface{}
	ReturnValues() []interface{}
}

type call struct {
	args []interface{}
	out  []interface{}
}

// Arguments returns the arguments that stub was called with.
func (c call) Arguments() []interface{} {
	return c.args
}

// ReturnValues returns the return values that the stubbed implementation returned.
func (c call) ReturnValues() []interface{} {
	return c.out
}
