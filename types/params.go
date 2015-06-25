package types

type Param interface {
	// Name returns the name of the synthdef param
	Name() string
	// Index returns the index of the synthdef param
	Index() int32
	// InitialValue returns the initial value of the synthdef param
	InitialValue() float32
}

type Params interface {
	// Add adds a named parameter to a synthdef, with an initial value
	Add(name string, initialValue float32) Input
	// List returns a list of the params that have been added to a synthdef
	List() []Param
	// Control returns a Ugen that should be used as the first ugen
	// of any synthdef that has parameters
	Control() Ugen
}
