package sc

// FreeVerb reverb implemented with faust
type FreeVerb struct {
	// In the input signal
	In Input
	// Mix dry/wet balance [0, 1]
	Mix Input
	// Room room size [0, 1]
	Room Input
	// Damp high frequency damping [0, 1]
	Damp Input
}

func (fv *FreeVerb) defaults() {
	if fv.Mix == nil {
		fv.Mix = C(0.33)
	}
	if fv.Room == nil {
		fv.Room = C(0.5)
	}
	if fv.Damp == nil {
		fv.Damp = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (fv FreeVerb) Rate(rate int8) Input {
	CheckRate(rate)
	if fv.In == nil {
		panic("FreeVerb expects In to not be nil")
	}
	(&fv).defaults()
	return UgenInput("FreeVerb", rate, 0, 1, fv.In, fv.Mix, fv.Room, fv.Damp)
}
