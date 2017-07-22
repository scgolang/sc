package sc

// Osc is a Linear interpolating wavetable lookup oscillator with frequency and phase modulation inputs.
// This oscillator requires a buffer to be filled with a wavetable format signal.
// This preprocesses the Signal into a form which can be used efficiently by the Oscillator.
// The buffer size must be a power of 2.
// This can be achieved by creating a Buffer object and sending it
// one of the "b_gen" messages ( Buffer: -sine1, Buffer: -sine2, Buffer: -sine3 )
// with the wavetable flag set to true.
// This can also be achieved by creating a Signal object and sending it the 'asWavetable' message,
// thereby creating a Wavetable object in the required format.
// Then, the wavetable data may be transmitted to the server
// using the Buffer: *sendCollection or Buffer: *loadCollection methods.
type Osc struct {
	BufNum Input
	Freq   Input
	Phase  Input
}

func (o *Osc) defaults() {
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
func (o Osc) Rate(rate int8) Input {
	CheckRate(rate)
	if o.BufNum == nil {
		panic("Osc requires a buffer number")
	}
	(&o).defaults()
	return NewInput("Osc", rate, 0, 1, o.BufNum, o.Freq, o.Phase)
}
