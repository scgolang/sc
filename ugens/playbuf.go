package ugens

import . "github.com/scgolang/sc/types"

// PlayBuf plays back a sample from memory.
// If Buf is nil a runtime panic will occur.
type PlayBuf struct {
	// NumChannels number of playback channels
	NumChannels int
	// Buf buffer to use
	Buf Buffer
	// Speed is the playback speed: 1.0 for original speed,
	// 2.0 for Chipmunks, 0.5 for DJ Screw
	Speed Input
	// Trigger when this input changes from negative to positive
	// playback will jump to the beginning of the buffer
	Trigger Input
	// Start is the sample frame to start playback
	Start Input
	// Loop 1 is on, 0 is off
	Loop Input
	// Done action to take when done playing buffer.
	// See http://doc.sccode.org/Reference/UGen-doneActions.html
	Done int
}

func (self *PlayBuf) defaults() {
	if self.NumChannels == 0 {
		self.NumChannels = 1
	}
	if self.Speed == nil {
		self.Speed = C(1.0)
	}
	if self.Trigger == nil {
		self.Trigger = C(1)
	}
	if self.Start == nil {
		self.Start = C(0)
	}
	if self.Loop == nil {
		self.Loop = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self PlayBuf) Rate(rate int8) Input {
	checkRate(rate)
	if self.Buf == nil {
		panic("PlayBuf expects Buf to not be nil")
	}
	(&self).defaults()
	nchans := C(float32(self.NumChannels))
	done := C(float32(self.Done))
	return UgenInput("PlayBuf", rate, 0, nchans, self.Speed, self.Trigger, self.Start, self.Loop, done)
}
