package sc

type SynthDef interface {
}

// NewSynthDef creates a new SynthDef from a UgenGraphFunc
func NewSynthDef(name string, f UgenGraphFunc) SynthDef {
	// TODO: pass an interface into the ugen graph func that allows
	//       users to define params for the synthdef
	graph, err := f()
	if err != nil {
		panic(err)
	}
	return graph.SynthDef()
}
