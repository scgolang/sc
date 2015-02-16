package sc

type UgenNode interface {
	Name() string
	Rate() int8
	Inputs() []Input
	Outputs() []Output
}
