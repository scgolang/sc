package sc

// GVerb is a two-channel reverb UGen.
// Based on the "GVerb" LADSPA effect by Juhana Sadeharju (kouhia at nic.funet.fi).
type GVerb struct {
	// Mono Input
	In Input
	// In squared meters
	RoomSize Input
	// In seconds
	RevTime Input
	// 0 to 1, high frequency rolloff, 0 damps the reverb signal completely, 1 not at all
	Damping Input
	// 0 to 1, same as damping control, but on the input signal.
	InputBW Input
	// A control on the stereo spread and diffusion of the reverb signal.
	Spread Input
	// Amount of dry signal.
	DryLevel Input
	// Amount of early reflection level.
	EarlyRefLevel Input
	// Amount of tail level.
	TailLevel Input
	// To set the size of the delay lines. Defaults to roomsize + 1.
	MaxRoomSize Input
}

func (gv *GVerb) defaults() {
	if gv.RoomSize == nil {
		gv.RoomSize = C(10)
	}
	if gv.RevTime == nil {
		gv.RevTime = C(3)
	}
	if gv.Damping == nil {
		gv.Damping = C(0.5)
	}
	if gv.InputBW == nil {
		gv.InputBW = C(0.5)
	}
	if gv.Spread == nil {
		gv.Spread = C(15)
	}
	if gv.DryLevel == nil {
		gv.DryLevel = C(1)
	}
	if gv.EarlyRefLevel == nil {
		gv.EarlyRefLevel = C(0.7)
	}
	if gv.TailLevel == nil {
		gv.TailLevel = C(0.5)
	}
	if gv.MaxRoomSize == nil {
		gv.MaxRoomSize = C(300)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (gv GVerb) Rate(rate int8) Input {
	CheckRate(rate)
	if gv.In == nil {
		panic("GVerb expects In to not be nil")
	}
	(&gv).defaults()
	return NewInput("GVerb", rate, 0, 2, gv.In, gv.RoomSize, gv.RevTime, gv.Damping, gv.InputBW, gv.Spread, gv.DryLevel, gv.EarlyRefLevel, gv.TailLevel, gv.MaxRoomSize)
}
