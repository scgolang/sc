package types

type Param interface {
	Name() string
	GetDefault() float32
	SetDefault(val float32) Param
}
