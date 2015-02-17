package types

// Synthdef
type Synthdef interface {
	// Name returns the name of the synthdef.
	Name() string

	// Constants returns the constants that appear
	// in a synthdef
	Constants() []float32
}
