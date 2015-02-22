package types

type UgenNode interface {
	// Name returns the name of the ugen node
	Name() string
	// Rate returns the rate of the ugen node
	Rate() int8
	// IsOutput tells the ugen node that it is being
	// used as the input for another ugen
	IsOutput()
	// Outputs returns the outputs of the ugen node
	Outputs() []Output
	// Inputs returns the inputs of the ugen node
	Inputs() []Input
}
