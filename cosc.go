package sc

import (
	"fmt"
)

// COsc is a chorusing wavetable lookup oscillator.
// Produces sum of 2 signals at
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

func (cosc *COsc) defaults() {
	if cosc.Freq == nil {
		cosc.Freq = C(440)
	}
	if cosc.Beats == nil {
		cosc.Beats = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
// There will also be a runtime panic if BufNum is nil.
func (cosc COsc) Rate(rate int8) Input {
	CheckRate(rate)
	if cosc.BufNum == nil {
		panic(fmt.Errorf("BufNum can not be nil"))
	}
	(&cosc).defaults()
	return UgenInput("COsc", rate, 0, 1, cosc.BufNum, cosc.Freq, cosc.Beats)
}
