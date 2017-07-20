package sc

// Klank is a bank of fixed frequency resonators which can be used to simulate the resonant modes of an object.
// Each mode is given a ring time, which is the time for the mode to decay by 60 dB.
type Klank struct {
	// In is the excitation input to the resonant filter bank.
	In Input

	// Spec is three Input slices (which should all have the same length).
	// The first slice is the filter frequencies.
	// The second slice is the filter amplitudes.
	// The third slice is the 60dB decay times in seconds for the filters.
	Spec Input

	// FreqScale is a scale factor multiplied by all frequencies at initialization time.
	FreqScale Input

	// FreqOffset is an offset added to all frequencies at initialization time.
	FreqOffset Input

	// DecayScale is a scale factor multiplied by all ring times at initialization time.
	DecayScale Input
}

func (k *Klank) defaults() {
	if k.FreqScale == nil {
		k.FreqScale = C(1)
	}
	if k.FreqOffset == nil {
		k.FreqOffset = C(0)
	}
	if k.DecayScale == nil {
		k.DecayScale = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If k.In is nil this method will cause a runtime panic.
// If k.Spec is not ArraySpec or Multi(ArraySpec...) then this method will panic.
func (k Klank) Rate(rate int8) Input {
	if k.In == nil {
		panic("Klank expects In to not be nil")
	}
	CheckRate(rate)
	(&k).defaults()

	specs := getArraySpecInputs(k.Spec)

	if len(specs) == 1 {
		ins := []Input{k.In, k.FreqScale, k.FreqOffset, k.DecayScale}
		ins = append(ins, specs[0].inputs(true)...)
		return NewInput("Klank", rate, 0, 1, ins...)
	}
	var klanks []Input
	for _, spec := range specs {
		ins := []Input{k.In, k.FreqScale, k.FreqOffset, k.DecayScale}
		ins = append(ins, spec.inputs(true)...)
		klanks = append(klanks, NewInput("Klank", rate, 0, 1, ins...))
	}
	return Multi(klanks...)
}
