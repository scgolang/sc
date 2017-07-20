package sc

// VarSaw is a sawtooth-triangle oscillator with variable duty.
type VarSaw struct {
	// Freq is the oscillator's frequency in Hz.
	Freq Input

	// IPhase is the oscillator's initial phase in cycles (0..1).
	IPhase Input

	// Width is the duty cycle from 0 to 1.
	Width Input
}

func (vs *VarSaw) defaults() {
	if vs.Freq == nil {
		vs.Freq = C(440)
	}
	if vs.IPhase == nil {
		vs.IPhase = C(0)
	}
	if vs.Width == nil {
		vs.Width = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (vs VarSaw) Rate(rate int8) Input {
	CheckRate(rate)
	(&vs).defaults()
	return NewInput("VarSaw", rate, 0, 1, vs.Freq, vs.IPhase, vs.Width)
}
