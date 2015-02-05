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
// Ugen (name=Out, index=1)
// |
// +--> Constant (value=0, index=1)
// |
// +--> Ugen (name=SinOsc, index=0)
//      |
//      +--> Constant(440) (index=0)
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
