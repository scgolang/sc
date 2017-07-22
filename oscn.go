package sc

// OscN is a noninterpolating wavetable lookup oscillator with frequency
// and phase modulation inputs.
// It is usually better to use the interpolating oscillator Osc.
type OscN struct {
	BufNum Input
	Freq   Input
	Phase  Input
}

func (o *OscN) defaults() {
	if o.Freq == nil {
		o.Freq = C(440)
	}
	if o.Phase == nil {
		o.Phase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If BufNum is nil this method will panic.
func (o OscN) Rate(rate int8) Input {
	CheckRate(rate)
	if o.BufNum == nil {
		panic("OscN requires a buffer number")
	}
	(&o).defaults()
	return NewInput("OscN", rate, 0, 1, o.BufNum, o.Freq, o.Phase)
}
