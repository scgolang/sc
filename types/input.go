package types

type Input interface {
	IsConstant() bool
}

type ConstantInput interface {
	Value() float32
}

type UgenInput interface {
	Value() UgenNode
}
