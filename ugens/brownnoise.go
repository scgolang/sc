package ugens

import . "github.com/scgolang/sc/types"

// Crackle
// A noise generator based on a chaotic function.
type Crackle struct {
	Chaos Input
}

func (self *Crackle) defaults() {
	if self.Chaos == nil {
		self.Chaos = C(1.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Crackle) Rate(rate int8) Input {
	checkRate(rate)
	(&self).defaults()
	return UgenInput("Crackle", rate, 0, 1, self.Chaos)
}
