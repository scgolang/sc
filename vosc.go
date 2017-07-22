package sc

// VOsc is a wavetable lookup oscillator which can be swept smoothly across wavetables.
// All the wavetables must be allocated to the same size.
// Fractional values of table will interpolate between two adjacent tables.
// This oscillator requires at least two buffers to be filled with a wavetable format signal.
// This preprocesses the Signal into a form which can be used efficiently by the Oscillator.
// The buffer size must be a power of 2.
// This can be achieved by creating a Buffer object and sending it one of the
// "b_gen" messages ( sine1, sine2, sine3 ) with the wavetable flag set to true.
// This can also be achieved by creating a Signal object and sending it the asWavetable message,
// saving it to disk, and having the server load it from there.
// If you use Buffer objects to manage buffer numbers, you can use the [*allocConsecutive] method
// to allocate a continuous block of buffers. See the Buffer helpfile for details.
type VOsc struct {
	// Buffer index. Can be swept continuously among adjacent wavetable buffers of the same size.
	BufNum Input

	// Frequency in Hz.
	Freq Input

	// Phase (0..1)
	Phase Input
}

func (o *VOsc) defaults() {
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
func (o VOsc) Rate(rate int8) Input {
	CheckRate(rate)
	if o.BufNum == nil {
		panic("VOsc requires a buffer number")
	}
	(&o).defaults()
	return NewInput("VOsc", rate, 0, 1, o.BufNum, o.Freq, o.Phase)
}
