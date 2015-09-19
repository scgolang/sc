package sc

import (
	"fmt"
)

// COsc
// Chorusing wavetable lookup oscillator. Produces sum of 2 signals at
//     freq +- (beats / 2)
// Due to summing, the peak amplitude is twice that of the wavetable.
type COsc struct {
	// BufNum the number of a buffer filled in wavetable format
	BufNum Input
	// Freq frequency in Hz
	Freq Input
	// Beats beat frequency in Hz
	Beats Input
}

func (self *COsc) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.Beats == nil {
		self.Beats = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
// There will also be a runtime panic if BufNum is nil.
func (self COsc) Rate(rate int8) Input {
	CheckRate(rate)
	if self.BufNum == nil {
		panic(fmt.Errorf("BufNum can not be nil"))
	}
	(&self).defaults()
	return UgenInput("COsc", rate, 0, 1, self.BufNum, self.Freq, self.Beats)
}
