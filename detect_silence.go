package sc

// DetectSilence evaluates Done when input falls below a certain threshold.
type DetectSilence struct {
	In   Input // The input signal.
	Amp  Input // Amplitude threshold.
	Time Input // The minimum duration for which amplitude must be below Amp before Done triggers.
	Done int   // UGen done-action.
}

func (ds *DetectSilence) defaults() {
	if ds.Amp == nil {
		ds.Amp = C(0.0001)
	}
	if ds.Time == nil {
		ds.Time = C(0.1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (ds DetectSilence) Rate(rate int8) Input {
	if ds.In == nil {
		panic("DetectSilence expects In to not be nil")
	}
	CheckRate(rate)
	(&ds).defaults()
	return UgenInput("DetectSilence", rate, 0, 1, ds.In, ds.Amp, ds.Time, C(float32(ds.Done)))
}
