package sc

// DecodeB2 Decodes a two dimensional ambisonic B-format signal to
// a set of speakers in a regular polygon.
// The outputs will be in clockwise order.
// The position of the first speaker is either center or left of center.
type DecodeB2 struct {
	// NumChans is the number of output speakers. Typically 4 to 8.
	NumChans int

	// The B-format signal.
	// This should be a ugen with 3 outputs: W, X, and Y.
	In Input

	// Orientation should be zero if the front is a vertex of the polygon.
	// The first speaker will be directly in front.
	// Should be 0.5 if the front bisects a side of the polygon.
	// Then the first speaker will be the one left of center.
	Orientation Input
}

func (d *DecodeB2) defaults() {
	if d.Orientation == nil {
		d.Orientation = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (d DecodeB2) Rate(rate int8) Input {
	CheckRate(rate)
	if d.In == nil {
		panic("DecodeB2 requires an input signal")
	}
	(&d).defaults()
	return NewInput("DecodeB2", rate, 0, d.NumChans, d.In, d.Orientation)
}
