package types

type Params interface {
	Add(name string) Param
	List() []Param
}

type Param interface {
	Name() string
	Index() int32
	IsConstant() bool // HACK
	GetDefault() float32
	SetDefault(val float32) Param
}
