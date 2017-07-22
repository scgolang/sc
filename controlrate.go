package sc

// ControlRate returns the current control rate of the server.
type ControlRate struct{}

// Rate creates a new ugen at a specific rate.
// If rate is not IR this method will cause a runtime panic.
func (s ControlRate) Rate(rate int8) Input {
	if rate != IR {
		panic("ControlRate only supports IR")
	}
	return NewInput("ControlRate", rate, 0, 1)
}
