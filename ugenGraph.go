package sc

// example graph for SineTone synth
//
// Ugen (name=Out, index=1, numInputs=2, numOutputs=0)
// |
// +--> Constant (index=1, value=0)
// |
// +--> Ugen (name=SinOsc, index=0, numInputs=1, numOutputs=1)
//      |
//      +--> Constant (index=0, value=440)

// UgenGraphFunc create a Ugen graph and return the root ugen
type UgenGraphFunc func() (*Ugen, error)

type UgenGraph interface {
}

type ugenGraph struct {
	root UgenNode
}

func NewUgenGraph(root UgenNode) UgenGraph {
	ugg := ugenGraph{root}
	return &ugg
}
