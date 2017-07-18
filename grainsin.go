package sc

// GrainSin implements granular synthesis with sine tones.
// All args except numChannels and trigger are polled at grain creation time.
type GrainSin struct {
	// NumChannels is the number of channels to output.
	// If 1, mono is returned and pan is ignored.
	NumChannels int

	// Trigger is a KR or AR trigger to start a new grain.
	// If AR, grains after the start of the synth are
	// sample-accurate.
	Trigger Input

	// Dur is the size of the grain in seconds.
	Dur Input

	// Freq is the input to granulate.
	Freq Input

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

func (g *GrainSin) defaults() {
	if g.NumChannels == 0 {
		g.NumChannels = 1
	}
	if g.Trigger == nil {
		g.Trigger = C(0)
	}
	if g.Dur == nil {
		g.Dur = C(1)
	}
	if g.Freq == nil {
		g.Freq = C(440)
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
func (g GrainSin) Rate(rate int8) Input {
	CheckRate(rate)
	(&g).defaults()

	in := NewInput("GrainSin", rate, 0, 1, g.Trigger, g.Dur, g.Freq, g.Pan, g.EnvBuf, g.MaxGrains)
	if g.NumChannels == 1 {
		return in
	}
	mult := make([]Input, g.NumChannels)
	for i := range mult {
		mult[i] = in
	}
	return Multi(mult...)
}
