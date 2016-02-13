package sc

// Impulse non-band-limited single-sample impulses
type Impulse struct {
	// Freq frequency in Hz
	Freq Input
	// Phase offset in cycles [0, 1]
	Phase Input
}

func (impulse *Impulse) defaults() {
	if impulse.Freq == nil {
		impulse.Freq = C(440)
	}
	if impulse.Phase == nil {
		impulse.Phase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (impulse Impulse) Rate(rate int8) Input {
	CheckRate(rate)
	(&impulse).defaults()
	return UgenInput("Impulse", rate, 0, 1, impulse.Freq, impulse.Phase)
}
