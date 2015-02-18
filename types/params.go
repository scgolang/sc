package types

// Params provides a way to add params to a synthdef
type Params interface {
	// Add adds a named parameter with an optional default value
	Add(name string, defaultValue ...float32)
	// Get gets all the parameters
	Get() []Param
}

type Param interface {
	Name() string
	DefaultValue() float32
}
