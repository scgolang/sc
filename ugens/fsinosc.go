package ugens

import . "github.com/briansorahan/sc/types"

// FSinOsc is a very fast sine wave generator implemented using a
// ringing filter. This generates a much cleaner sine wave than a
// table lookup oscillator and is a lot faster. However, the
// amplitude of the wave will vary with frequency. Generally the
// amplitude will go down as you raise the frequency and go up
// as you lower the frequency.
type FSinOsc struct {
	// Freq is frequency in Hz
	Freq Input
	// Phase is the initial phase offset
	Phase Input
}

func (self *FSinOsc) defaults() {
	if self.Freq == nil {
		self.Freq = C(440)
	}
	if self.Phase == nil {
		self.Phase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self FSinOsc) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	return UgenInput("FSinOsc", rate, 0, self.Freq, self.Phase)
}
