package sc

// FFT implements the Fast Fourier Transform.
// The fast fourier transform analyzes the frequency content of a signal,
// which can be useful for audio analysis or for frequency-domain sound processing (phase vocoder).
type FFT struct {
	// A buffer to store spectral data. The buffer's size must correspond to a power of 2.
	// LocalBuf is useful here, because processes should not share data between synths.
	// (Note: most PV UGens operate on this data in place. Use PV_Copy for parallel processing.)
	Buffer Input
	// The signal to be analyzed. The signal's rate determines the rate at which the input is read.
	In Input
	// The amount of offset from the beginning of one FFT analysis frame to the next, measured in multiples of the analysis frame size.
	// This can range between 1.0 and values close to (but larger than) 0.0,
	// and the default is 0.5 (meaning each frame has a 50% overlap with the preceding/following frames).
	Hop Input
	// Defines how the data is windowed:
	// -1 is rectangular windowing, simple but typically not recommended
	// 0 is (the default) Sine windowing, typically recommended for phase-vocoder work
	// 1 is Hann windowing, typically recommended for analysis work.
	WinType Input
	// A simple control allowing FFT analysis to be active (>0) or inactive (<=0).
	// This is mainly useful for signal analysis processes which are only intended to analyse at specific times rather than continuously
	Active Input
	// The windowed audio frames are usually the same size as the buffer.
	// If you wish the FFT to be zero-padded then you can specify a window size smaller than the actual buffer size
	// (e.g. window size 1024 with buffer size 2048).
	// Both values must still be a power of two. Leave this at its default of zero for no zero-padding.
	WinSize Input
}

func (fft *FFT) defaults() {
	if fft.Hop == nil {
		fft.Hop = C(0.5)
	}
	if fft.WinType == nil {
		fft.In = C(0)
	}
	if fft.Active == nil {
		fft.Active = C(1)
	}
	if fft.WinSize == nil {
		fft.WinSize = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If an In signal or Buffer is not provided this method will trigger a runtime panic.
func (fft FFT) Rate(rate int8) Input {
	CheckRate(rate)
	if fft.Buffer == nil {
		panic("FFT expects Buffer to not be nil")
	}
	if fft.In == nil {
		panic("FFT expects In to not be nil")
	}
	(&fft).defaults()
	return NewInput("FFT", rate, 0, 1, fft.Buffer, fft.In, fft.Hop, fft.WinType, fft.Active, fft.WinSize)
}
