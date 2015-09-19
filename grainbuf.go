package sc

import (
	"fmt"
)

const (
	GrainBufHanningEnv       = -1
	GrainBufNoInterp         = 1
	GrainBufLinearInterp     = 2
	GrainBufCubicInterp      = 4
	GrainBufDefaultMaxGrains = 512
)

// GrainBuf is a table-lookup sinewave oscillator
type GrainBuf struct {
	// NumChannels is the number of channels to output.
	// If 1, mono is returned and pan is ignored.
	NumChannels int
	// Trigger is a KR or AR trigger to start a new grain.
	// If AR, grains after the start of the synth are
	// sample-accurate.
	Trigger Input
	// Dur is the size of the grain (in seconds)
	Dur Input
	// BufNum is the buffer holding a mono audio signal
	BufNum Input
	// Speed is the playback speed of the grain
	Speed Input
	// Pos is the position in the audio buffer where
	// the grain will start. This is in the range [0, 1].
	Pos Input
	// Interp is the interpolation method used for
	// pitch-shifting grains.
	// GrainBufNoInterp is no interpolation,
	// GrainBufLinearInterp is linear,
	// and GrainBufCubicInterp is cubic.
	Interp Input
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
	// when you initialize GrainBuf and can't be modified.
	// Default is 512, but lower values may result in more
	// efficient use of memory.
	MaxGrains Input
}

func (self *GrainBuf) defaults() {
	if self.NumChannels == 0 {
		self.NumChannels = 1
	}
	if self.Trigger == nil {
		self.Trigger = C(0)
	}
	if self.Dur == nil {
		self.Dur = C(1)
	}
	if self.Speed == nil {
		self.Speed = C(1)
	}
	if self.Pos == nil {
		self.Pos = C(0)
	}
	if self.Interp == nil {
		self.Interp = C(GrainBufLinearInterp)
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
func (self GrainBuf) Rate(rate int8) Input {
	CheckRate(rate)
	if self.BufNum == nil {
		panic(fmt.Errorf("BufNum can not be nil"))
	}
	(&self).defaults()
	return UgenInput("GrainBuf", rate, 0, self.NumChannels, self.Trigger, self.Dur, self.BufNum, self.Speed, self.Pos, self.Interp, self.Pan, self.EnvBuf, self.MaxGrains)
}
