package sc

// DC creates a constant amplitude signal.
type DC struct {
	In Input
}

func (dc *DC) defaults() {
	if dc.In == nil {
		dc.In = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (dc DC) Rate(rate int8) Input {
	CheckRate(rate)
	(&dc).defaults()
	return NewInput("DC", rate, 0, 1, dc.In)
}
