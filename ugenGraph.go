package sc

// UgenGraphFunc create a UgenGraph
type UgenGraphFunc func(ugen Ugen) UgenGraph

// UgenGraph
type UgenGraph interface {
	// Add adds a constant to the output of a UgenGraph
	Add(x float32) UgenGraph
	// Mul multiplies the output of a UgenGraph by a constant
	Mul(x float32) UgenGraph
	// SynthDef converts a UgenGraph to a SynthDef
	SynthDef() SynthDef
}

//
// Out
// |
// ---- Constant(0)
// |
// ---- SinOsc
//      |
//      ---- Constant(440)
//
type ugenGraph struct {
	root *Node
}

// TODO: implement
func (self *ugenGraph) Add(x float32) UgenGraph {
	return self
}

// TODO: implement
func (self *ugenGraph) Mul(x float32) UgenGraph {
	return self
}

// TODO: implement
func (self *ugenGraph) SynthDef() SynthDef {
	return nil
}

func NewUgenGraph(root *Node) UgenGraph {
	ugg := ugenGraph{root}
	return &ugg
}
