package sc

import (
	"fmt"
)

// PlayBuf plays back a sample from memory.
// If Buf is nil a runtime panic will occur.
type PlayBuf struct {
	// NumChannels number of playback channels
	NumChannels int
	// BufNum buffer to use
	BufNum Input
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
// There will also be a runtime panic if BufNum is nil.
func (self PlayBuf) Rate(rate int8) Input {
	CheckRate(rate)
	if self.BufNum == nil {
		panic(fmt.Errorf("BufNum can not be nil"))
	}
	(&self).defaults()
	done := C(float32(self.Done))
	return UgenInput("PlayBuf", rate, 0, self.NumChannels, self.BufNum, self.Speed, self.Trigger, self.Start, self.Loop, done)
}
