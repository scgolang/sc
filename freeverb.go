package sc


// FreeVerb reverb implemented with faust
type FreeVerb struct {
	// In the input signal
	In Input
	// Mix dry/wet balance [0, 1]
	Mix Input
	// Room room size [0, 1]
	Room Input
	// Damp high frequency damping [0, 1]
	Damp Input
}

func (self *FreeVerb) defaults() {
	if self.Mix == nil {
		self.Mix = C(0.33)
	}
	if self.Room == nil {
		self.Room = C(0.5)
	}
	if self.Damp == nil {
		self.Damp = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal is not provided this method will
// trigger a runtime panic.
func (self FreeVerb) Rate(rate int8) Input {
	CheckRate(rate)
	if self.In == nil {
		panic("FreeVerb expects In to not be nil")
	}
	(&self).defaults()
	return UgenInput("FreeVerb", rate, 0, 1, self.In, self.Mix, self.Room, self.Damp)
}
