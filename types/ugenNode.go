package types

type UgenNode interface {
	Name() string
	Rate() int8
	Inputs() []Input
	Outputs() []Output
	// EnsureOutput is how we tell a ugen node
	// that its output is being used as the input to
	// another ugen node.
	// All it has to do is ensure that the node has
	// at least one output whose rate is that of the
	// node itself.
	EnsureOutput()
}
