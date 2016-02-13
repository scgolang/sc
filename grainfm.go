package sc

// GrainFM is a table-lookup sinewave oscillator
type GrainFM struct {
	// NumChannels is the number of channels to output.
	// If 1, mono is returned and pan is ignored.
	NumChannels int
	// Trigger is a KR or AR trigger to start a new grain.
	// If AR, grains after the start of the synth are
	// sample-accurate.
	Trigger Input
	// Dur is the size of the grain (in seconds)
	Dur Input
	// CarFreq the carrier frequency of the grain generator's internal oscillator
	CarFreq Input
	// ModFreq the modulator frequency of the grain generator's internal oscillator
	ModFreq Input
	// ModIndex the index of modulation
	ModIndex Input
	// Pan determines where to position the output in a stereo
	// field. If NumChannels = 1, no panning is done. If
	// NumChannels = 2, behavior is similar to Pan2. If
	// NumChannels > 2, behavior is the same as PanAz.
	Pan Input
	// EnvBuf is the buffer number containing a signal to use
	// for each grain's amplitude envelope. If set to
	// GrainBufHanningEnv, a built-in Hanning envelope is used.
	EnvBuf Input
	// MaxGrains is the maximum number of overlapping grains
	// that can be used at a given time. This value is set
	// when you initialize GrainFM and can't be modified.
	// Default is 512, but lower values may result in more
	// efficient use of memory.
	MaxGrains Input
}

func (gfm *GrainFM) defaults() {
	if gfm.NumChannels == 0 {
		gfm.NumChannels = 1
	}
	if gfm.Trigger == nil {
		gfm.Trigger = C(0)
	}
	if gfm.Dur == nil {
		gfm.Dur = C(1)
	}
	if gfm.CarFreq == nil {
		gfm.CarFreq = C(440)
	}
	if gfm.ModFreq == nil {
		gfm.ModFreq = C(200)
	}
	if gfm.ModIndex == nil {
		gfm.ModIndex = C(1)
	}
	if gfm.Pan == nil {
		gfm.Pan = C(0)
	}
	if gfm.EnvBuf == nil {
		gfm.EnvBuf = C(GrainBufHanningEnv)
	}
	if gfm.MaxGrains == nil {
		gfm.MaxGrains = C(GrainBufDefaultMaxGrains)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
// There will also be a runtime panic if BufNum is nil.
func (gfm GrainFM) Rate(rate int8) Input {
	CheckRate(rate)
	(&gfm).defaults()
	return UgenInput("GrainFM", rate, 0, gfm.NumChannels, gfm.Trigger, gfm.Dur, gfm.CarFreq, gfm.ModFreq, gfm.ModIndex, gfm.Pan, gfm.EnvBuf, gfm.MaxGrains)
}
