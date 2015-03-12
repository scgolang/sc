package types

// UGen http://doc.sccode.org/Classes/UGen.html
type Ugen interface {
	// Rate creates a ugen at one of the supported rates.
	// Supported rates are AR, KR, and IR.
	// Unsupported rates cause a runtime panic.
	Rate(int8) UgenNode
}
