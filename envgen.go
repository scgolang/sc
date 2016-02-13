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

func (envgen *EnvGen) defaults() {
	if envgen.Gate == nil {
		envgen.Gate = C(1)
	}
	if envgen.LevelScale == nil {
		envgen.LevelScale = C(1)
	}
	if envgen.LevelBias == nil {
		envgen.LevelBias = C(0)
	}
	if envgen.TimeScale == nil {
		envgen.TimeScale = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (envgen EnvGen) Rate(rate int8) Input {
	CheckRate(rate)
	(&envgen).defaults()
	ins := []Input{
		envgen.Gate,
		envgen.LevelScale,
		envgen.LevelBias,
		envgen.TimeScale,
		C(float32(envgen.Done)),
	}
	ins = append(ins, envgen.Env.Inputs()...)
	return UgenInput("EnvGen", rate, 0, 1, ins...)
}
