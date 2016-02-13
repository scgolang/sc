package sc

// ClipNoise generates noise whose values are either -1 or 1.
// This produces the maximum energy for the least peak to peak amplitude.
type ClipNoise struct{}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (cn ClipNoise) Rate(rate int8) Input {
	return UgenInput("ClipNoise", rate, 0, 1)
}
