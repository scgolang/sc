package types

type UgenGraphFunc func() UgenNode

type UgenGraph interface {
	Root() UgenNode
}
