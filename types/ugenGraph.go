package types

type UgenGraphFunc func(params Params) UgenNode

type UgenGraph interface {
	Root() UgenNode
}
