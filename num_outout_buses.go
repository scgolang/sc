package sc

// NumOutputBuses provides the number of output buses in a synthdef.
type NumOutputBuses struct{}

// Rate creates a new ugen at a specific rate.
// If rate is not IR this method will cause a runtime panic.
func (n NumOutputBuses) Rate(rate int8) Input {
	if rate != IR {
		panic("NumOutputBuses must be rate IR")
	}
	return NewInput("NumOutputBuses", rate, 0, 1)
}
