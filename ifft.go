package sc

// IFFT converts from frequency content to a signal.
// The fast fourier transform analyzes the frequency content of a signal.
// The IFFT UGen converts this frequency-domain information back into
// time-domain audio data.
// Most often this is used as the end of a process which begins with FFT,
// followed by frequency-domain processing using PV (phase-vocoder) UGens,
// followed by IFFT.
type IFFT struct {
	// The FFT "chain" signal coming originally from an FFT UGen,
	// perhaps via other PV UGens.
	Buffer Input

	// WinType defines how the data is windowed:
	//   -1 is rectangular windowing, simple but typically not recommended;
	//    0 (the default) is sine windowing, typically recommended for phase-vocoder work;
	//    1 is Hann windowing, typically recommended for analysis work.
	WinType Input

	// WinSize can be used to account for zero-padding,
	// in the same way as the FFT UGen.
	WinSize Input
}

func (ifft *IFFT) defaults() {
	if ifft.WinType == nil {
		ifft.WinType = C(0)
	}
	if ifft.WinSize == nil {
		ifft.WinSize = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (ifft IFFT) Rate(rate int8) Input {
	CheckRate(rate)
	(&ifft).defaults()
	return NewInput("IFFT", rate, 0, 1, ifft.Buffer, ifft.WinType, ifft.WinSize)
}
