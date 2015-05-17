package ugens

import . "github.com/scgolang/sc/types"

const (
	GrainBufHanningEnv       = -1
	GrainBufNoInterp         = 1
	GrainBufLinearInterp     = 1
	GrainBufCubicInterp      = 1
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
	// Buf is the buffer holding a mono audio signal
	Buf int
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
	Interp int
	// Pan determines where to position the output in a stereo
	// field. If NumChannels = 1, no panning is done. If
	// NumChannels = 2, behavior is similar to Pan2. If
	// NumChannels > 2, behavior is the same as PanAz.
	Pan Input
	// EnvBuf is the buffer number containing a signal to use
	// for each grain's amplitude envelope. If set to
	// GrainBufHanningEnv, a built-in Hanning envelope is used.
	EnvBuf int
	// MaxGrains is the maximum number of overlapping grains
	// that can be used at a given time. This value is set
	// when you initialize GrainBuf and can't be modified.
	// Default is 512, but lower values may result in more
	// efficient use of memory.
	MaxGrains int
}

func (self *GrainBuf) defaults() {
	if self.NumChannels == 0 {
		self.NumChannels = 1
	}
	if self.Trigger == nil {
		self.Trigger = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self GrainBuf) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	return UgenInput("GrainBuf", rate, 0, self.NumChannels, self.Trigger)
}
