package sc

// VOsc3 is a wavetable lookup oscillator which can be swept smoothly across wavetables.
// All the wavetables must be allocated to the same size.
// Fractional values of table will interpolate between two adjacent tables.
// This unit generator contains three oscillators at different frequencies, mixed together.
// This oscillator requires at least two buffers to be filled with a wavetable format signal.
// This preprocesses the Signal into a form which can be used efficiently by the Oscillator.
// The buffer size must be a power of 2.
// This can be achieved by creating a Buffer object and sending it one of
// the "b_gen" messages ( sine1, sine2, sine3 ) with the wavetable flag set to true.
// This can also be achieved by creating a Signal object and sending it
// the asWavetable message, saving it to disk, and having the server load it from there.
// If you use Buffer objects to manage buffer numbers, you can use the [*allocConsecutive] method
// to allocate a continuous block of buffers. See the Buffer helpfile for details.
type VOsc3 struct {
	// Buffer index. Can be swept continuously among adjacent wavetable buffers of the same size.
	BufNum Input

	// Frequency in Hz of oscillator 1.
	Freq1 Input

	// Frequency in Hz of oscillator 2.
	Freq2 Input

	// Frequency in Hz of oscillator 3.
	Freq3 Input
}

func (o *VOsc3) defaults() {
	if o.Freq1 == nil {
		o.Freq1 = C(110)
	}
	if o.Freq2 == nil {
		o.Freq2 = C(220)
	}
	if o.Freq3 == nil {
		o.Freq3 = C(440)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If BufNum is nil this method will panic.
func (o VOsc3) Rate(rate int8) Input {
	CheckRate(rate)
	if o.BufNum == nil {
		panic("VOsc3 requires a buffer number")
	}
	(&o).defaults()
	return NewInput("VOsc3", rate, 0, 1, o.BufNum, o.Freq1, o.Freq2, o.Freq3)
}
