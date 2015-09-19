package sc


// EnvGen plays back breakpoint envelopes.
// The envelopes must implement the Envelope interface.
// The envelope and the arguments for LevelScale, LevelBias,
// and TimeScale are polled when the EnvGen is triggered and
// remain constant for the duration of the envelope.
type EnvGen struct {
	// Env determines the shape of the envelope
	Env Envelope
	// Gate triggers the envelope and holds it open while > 0
	Gate Input
	// LevelScale scales the levels of the breakpoints
	LevelScale Input
	// LevelBias offsets the levels of the breakpoints
	LevelBias Input
	// TimeScale scales the durations of the segments
	TimeScale Input
	// Done is the ugen done action
	Done int
}

func (self *EnvGen) defaults() {
	if self.Gate == nil {
		self.Gate = C(1)
	}
	if self.LevelScale == nil {
		self.LevelScale = C(1)
	}
	if self.LevelBias == nil {
		self.LevelBias = C(0)
	}
	if self.TimeScale == nil {
		self.TimeScale = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self EnvGen) Rate(rate int8) Input {
	CheckRate(rate)
	(&self).defaults()
	ins := []Input{self.Gate, self.LevelScale, self.LevelBias}
	ins = append(ins, self.TimeScale, C(float32(self.Done)))
	ins = append(ins, self.Env.Inputs()...)
	return UgenInput("EnvGen", rate, 0, 1, ins...)
}
