package sc

type Ugen interface {
	Ar(args ...interface{}) UgenNode
	Kr(args ...interface{}) UgenNode
	Ir(args ...interface{}) UgenNode
}
