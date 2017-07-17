package sc

// GrainIn granulates an input signal.
type GrainIn struct {
	// NumChannels is the number of channels to output.
	// If 1, mono is returned and pan is ignored.
	NumChannels int

	// Trigger is a KR or AR trigger to start a new grain.
	// If AR, grains after the start of the synth are
	// sample-accurate.
	Trigger Input

	// Dur is the size of the grain in seconds.
	Dur Input

	// In is the input signal.
	In Input

	// Pan determines where to position the output in a stereo
	// field. If NumChannels = 1, no panning is done. If
	// NumChannels = 2, behavior is similar to Pan2. If
	// NumChannels > 2, behavior is the same as PanAz.
	Pan Input

	// EnvBuf is the buffer holding a mono audio signal.
	EnvBuf Input

	// MaxGrains is the maximum number of overlapping grains
	// that can be used at a given time. This value is set
	// when you initialize GrainBuf and can't be modified.
	// Default is 512, but lower values may result in more
	// efficient use of memory.
	MaxGrains Input
}

func (g *GrainIn) defaults() {
	if g.NumChannels == 0 {
		g.NumChannels = 1
	}
	if g.Trigger == nil {
		g.Trigger = C(0)
	}
	if g.Dur == nil {
		g.Dur = C(1)
	}
	if g.Pan == nil {
		g.Pan = C(0)
	}
	if g.EnvBuf == nil {
		g.EnvBuf = C(GrainBufHanningEnv)
	}
	if g.MaxGrains == nil {
		g.MaxGrains = C(GrainBufDefaultMaxGrains)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
// If the trig signal is nil this method will panic.
func (g GrainIn) Rate(rate int8) Input {
	CheckRate(rate)
	if g.In == nil {
		panic("GrainIn requires an input signal")
	}
	(&g).defaults()

	in := NewInput("GrainIn", rate, 0, 1, g.Trigger, g.Dur, g.In, g.Pan, g.EnvBuf, g.MaxGrains)
	if g.NumChannels == 1 {
		return in
	}
	mult := make([]Input, g.NumChannels)
	for i := range mult {
		mult[i] = in
	}
	return Multi(mult...)
}
