package sc

// Balance2 equal power panner
type Balance2 struct {
	// L is the left input signal
	L Input
	// R is the right input signal
	R Input
	// Pos stereo position where -1 is hard left and +1 is hard right
	Pos Input
	// Level gain [0, 1]
	Level Input
}

func (bal *Balance2) defaults() {
	if bal.Pos == nil {
		bal.Pos = C(0)
	}
	if bal.Level == nil {
		bal.Level = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (bal Balance2) Rate(rate int8) Input {
	if bal.L == nil || bal.R == nil {
		panic("Balance2 expects L and R to not be nil")
	}
	CheckRate(rate)
	(&bal).defaults()
	return UgenInput("Balance2", rate, 0, 2, bal.L, bal.R, bal.Pos, bal.Level)
}
