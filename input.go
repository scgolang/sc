package sc

type Input interface {
	IsConstant() bool
	Value() interface{}
}
