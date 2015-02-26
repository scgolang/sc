package types

type Params interface {
	Add(name string) Param
	List() []Param
	Control() UgenNode
}

type Param interface {
	Name() string
	Index() int32
	GetDefault() float32
	SetDefault(val float32) Param
}
