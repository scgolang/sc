package types

// Synthdef
type Synthdef interface {
	// Name returns the name of the synthdef.
	Name() string

	// Constants returns the constants that appear
	// in a synthdef
	Constants() []float32

	// InitialParamValues returns the initial values
	// for the synthdef's parameters
	InitialParamValues() []float32

	// ParamNames returns the names of the parameters
	// of the synthdef
	ParamNames() []string

	// Ugens returns the list of ugens present in
	// the synthdef
	Ugens() []Ugen

	// Variants returns the list of variants for
	// a synthdef
	Variants() []Variant
}
