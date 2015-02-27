package types

// Params provides a way to add parameters to a synthdef
type Params interface {
	// Add adds a parameter to the synthdef with an optional
	// initial value. The parameter's initial value will be
	// 0 if you do not provide one.
	Add(name string, initialValue ...float32) Param
	// List gets a list of all the synthdef parameters
	List() []Param
	// Control returns the ugen node that outputs the
	// synthdef parameters
	Control() UgenNode
}

// Param represents a synthdef parameter
type Param interface {
	Name() string
	Index() int32
	GetDefault() float32
	SetDefault(val float32) Param
}
