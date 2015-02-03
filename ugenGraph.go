package gosc

type UgenGraph interface {
	SynthDef() *SynthDef
}

// Out
// |
// ---- Constant(0)
// |
// ---- SinOsc
//      |
//      ---- Constant(440)
type ugenGraph struct {
	root *Node
}

func (self *ugenGraph) SynthDef() *SynthDef {
	return nil
}

func NewUgenGraph(root *Node) UgenGraph {
	ugg := ugenGraph{root}
	return &ugg
}

type UgenGraphFunc func(args ...interface{}) UgenGraph
