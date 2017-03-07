package sc

// LFDNoise generates random values at a rate given by the freq argument, with two
// differences: no time quantization, and fast recovery from low freq values.
type LFDNoise struct {
	Interpolation Interpolation
	Freq          Input
}

func (lfdn *LFDNoise) defaults() {
	if lfdn.Freq == nil {
		lfdn.Freq = C(500)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (lfdn LFDNoise) Rate(rate int8) Input {
	CheckRate(rate)
	(&lfdn).defaults()
	switch lfdn.Interpolation {
	case InterpolationNone:
		return UgenInput("LFDNoise0", rate, 0, 1, lfdn.Freq)
	case InterpolationLinear:
		return UgenInput("LFDNoise1", rate, 0, 1, lfdn.Freq)
	case InterpolationCubic:
		return UgenInput("LFDNoise3", rate, 0, 1, lfdn.Freq)
	default:
		panic("unknown interpolation value")
	}
}
