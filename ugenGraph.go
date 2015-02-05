package sc

// UgenGraphFunc create a UgenGraph
type UgenGraphFunc func(ugen Ugen) UgenGraph

// UgenGraph
type UgenGraph interface {
	Root() *Ugen
	// SynthDef converts a UgenGraph to a SynthDef
	SynthDef() *SynthDef
}

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
type ugenGraph struct {
	root *Ugen
}

func (self *ugenGraph) Root() *Ugen {
	return self.root
}

// TODO: implement
func (self *ugenGraph) SynthDef() *SynthDef {
	return nil
}
