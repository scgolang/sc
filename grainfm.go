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

func (self *GrainFM) defaults() {
	if self.NumChannels == 0 {
		self.NumChannels = 1
	}
	if self.Trigger == nil {
		self.Trigger = C(0)
	}
	if self.Dur == nil {
		self.Dur = C(1)
	}
	if self.CarFreq == nil {
		self.CarFreq = C(440)
	}
	if self.ModFreq == nil {
		self.ModFreq = C(200)
	}
	if self.ModIndex == nil {
		self.ModIndex = C(1)
	}
	if self.Pan == nil {
		self.Pan = C(0)
	}
	if self.EnvBuf == nil {
		self.EnvBuf = C(GrainBufHanningEnv)
	}
	if self.MaxGrains == nil {
		self.MaxGrains = C(GrainBufDefaultMaxGrains)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
// There will also be a runtime panic if BufNum is nil.
func (self GrainFM) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	return UgenInput("GrainFM", rate, 0, self.NumChannels, self.Trigger, self.Dur, self.CarFreq, self.ModFreq, self.ModIndex, self.Pan, self.EnvBuf, self.MaxGrains)
}
