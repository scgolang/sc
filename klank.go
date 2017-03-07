package sc

// Klank is a bank of fixed frequency resonators which can be used to simulate the resonant modes of an object.
// Each mode is given a ring time, which is the time for the mode to decay by 60 dB.
type Klank struct {
	// Spec is three float slices (which should all have the same length).
	// The first slice is the filter frequencies.
	// The second slice is the filter amplitudes.
	// The third slice is the 60dB decay times in seconds for the filters.
	Spec [3][]float32

	In         Input // In is the excitation input to the resonant filter bank.
	FreqScale  Input // FreqScale is a scale factor multiplied by all frequencies at initialization time.
	FreqOffset Input // FreqOffset is an offset added to all frequencies at initialization time.
	DecayScale Input // DecayScale is a scale factor multiplied by all ring times at initialization time.
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
// If In is nil this method will cause a runtime panic.
func (k Klank) Rate(rate int8) Input {
	if k.In == nil {
		panic("Klank expects In to not be nil")
	}
	CheckRate(rate)
	(&k).defaults()
	return UgenInput("Klank", rate, 0, 1, k.In, k.FreqScale, k.FreqOffset, k.DecayScale)
}
