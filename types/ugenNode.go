package types

type UgenNode interface {
	// Name returns the name of the ugen node
	Name() string
	// Rate returns the rate of the ugen node
	Rate() int8
	// Inputs returns the inputs of the ugen node
	Inputs() []Input
	// Outputs returns the outputs of the ugen node
	Outputs() []Output
}
