package sc

// Warp1 is a granular time stretcher and pitchshifter.
// Inspired by Chad Kirby's SuperCollider2 Warp1 class, which was inspired by Richard Karpen's sndwarp for CSound.
type Warp1 struct {
	// Number of output channels.
	NumChannels int
	// The index of the buffer to use. It must be a one channel (mono) buffer.
	BufNum Input
	// The position in the buffer. The value should be between 0 and 1, with 0 being the beginning of the buffer, and 1 the end.
	Pointer Input
	// The amount of frequency shift. 1.0 is normal, 0.5 is one octave down, 2.0 is one octave up. Negative values play the soundfile backwards.
	FreqScale Input
	// The size of each grain window.
	WindowSize Input
	// The buffer number containing a signal to use for the grain envelope. -1 uses a built-in Hanning envelope.
	EnvBufNum Input
	// The number of overlapping windows.
	Overlaps Input
	// The amount of randomness to the windowing function. Must be between 0 (no randomness) to 1.0 (probably too random actually)
	WindowRandRatio Input
	// The interpolation method used for pitchshifting grains. 1 = no interpolation. 2 = linear. 4 = cubic interpolation (more computationally intensive).
	Interp Input
}

func (wrp *Warp1) defaults() {
	if wrp.NumChannels == 0 {
		wrp.NumChannels = 1
	}
	if wrp.BufNum == nil {
		wrp.BufNum = C(0)
	}
	if wrp.Pointer == nil {
		wrp.Pointer = C(0)
	}
	if wrp.FreqScale == nil {
		wrp.FreqScale = C(1)
	}
	if wrp.WindowSize == nil {
		wrp.WindowSize = C(0.2)
	}
	if wrp.EnvBufNum == nil {
		wrp.EnvBufNum = C(-1)
	}
	if wrp.Overlaps == nil {
		wrp.Overlaps = C(8)
	}
	if wrp.WindowRandRatio == nil {
		wrp.WindowRandRatio = C(0)
	}
	if wrp.Interp == nil {
		wrp.Interp = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (wrp Warp1) Rate(rate int8) Input {
	CheckRate(rate)
	(&wrp).defaults()
	return NewInput("Warp1", rate, 0, wrp.NumChannels, wrp.BufNum, wrp.Pointer, wrp.FreqScale, wrp.WindowSize, wrp.EnvBufNum, wrp.Overlaps, wrp.WindowRandRatio, wrp.Interp)
}
