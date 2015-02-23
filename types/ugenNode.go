package types

type UgenNode interface {
	// Name returns the name of the ugen node
	Name() string
	// Rate returns the rate of the ugen node
	Rate() int8
	// SpecialIndex returns the special index of the ugen node
	SpecialIndex() int16
	// Inputs returns the inputs of the ugen node
	Inputs() []Input
	// Outputs returns the outputs of the ugen node
	Outputs() []Output
	// Mul multiplies the output of a ugen node by a constant
	// and returns a new ugen node representing the resulting audio stream
	Mul(f float32) UgenNode
	// Add adds a constant to the output of a ugen node
	// and returns a new ugen node representing the resulting audio stream
	Add(f float32) UgenNode
}
