package sc

// Median returns the median of the last length input points.
// This non-linear filter is good at reducing impulse noise from a signal.
type Median struct {
	In     Input
	Length Input
}

func (m *Median) defaults() {
	if m.Length == nil {
		m.Length = C(3)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
func (m Median) Rate(rate int8) Input {
	CheckRate(rate)
	if m.In == nil {
		panic("Median requires an input signal")
	}
	(&m).defaults()
	return NewInput("Median", rate, 0, 1, m.Length, m.In)
}
