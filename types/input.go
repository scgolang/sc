package types

type Input interface {
	IsConstant() bool
	Value() interface{}
}
