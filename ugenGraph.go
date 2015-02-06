package sc

// UgenGraphFunc create a UgenGraph
type UgenGraphFunc func() (*UgenGraph, error)

//
// example graph for SineTone synth
//
// Ugen (name=Out, index=1, numInputs=2, numOutputs=0)
// |
// +--> Constant (index=1, value=0)
// |
// +--> Ugen (name=SinOsc, index=0, numInputs=1, numOutputs=1)
//      |
//      +--> Constant (index=0, value=440)
//
type UgenGraph struct {
	root *Ugen
}

// TODO: implement
func (self *UgenGraph) SynthDef() SynthDef {
	return nil
}

func NewUgenGraph(root *Ugen) *UgenGraph {
	return &UgenGraph{root}
}
