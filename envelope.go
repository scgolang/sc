package sc

// Envelope is an interface that should be implemented by
// types that define envelopes for EnvGen.
type Envelope interface {
	// InputsArray provides EnvGen with the data it needs
	// to get a list of inputs.
	Inputs() []Input
}
