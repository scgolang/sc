package sc

// PMOsc is a phase modulation sine-wave oscillator pair.
type PMOsc struct {
	// CarFreq is the carrier frequency in Hz.
	CarFreq Input
	// ModFreq is the modulator frequency in Hz.
	ModFreq Input
	// PMIndex in the modulation index in radians.
	PMIndex Input
	// ModPhase is a modulation input for the modulator's phase in radians.
	ModPhase Input
}

func (pmosc *PMOsc) defaults() {
	if pmosc.CarFreq == nil {
		pmosc.CarFreq = C(440)
	}
	if pmosc.ModFreq == nil {
		pmosc.ModFreq = C(440)
	}
	if pmosc.PMIndex == nil {
		pmosc.PMIndex = C(0)
	}
	if pmosc.ModPhase == nil {
		pmosc.ModPhase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (pmosc PMOsc) Rate(rate int8) Input {
	CheckRate(rate)
	(&pmosc).defaults()
	modulator := SinOsc{Freq: pmosc.ModFreq, Phase: pmosc.ModPhase}.Rate(rate).Mul(pmosc.PMIndex)
	// return the carrier
	return SinOsc{Freq: pmosc.CarFreq, Phase: modulator}.Rate(rate)
}
