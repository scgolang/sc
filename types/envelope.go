package types

type Envelope interface {
	// InputsArray provides EnvGen with the data it needs
	// to get a list of inputs
	Inputs() []Input
}

