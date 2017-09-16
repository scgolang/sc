package sc

// BLowPass4 is a lowpass filter based on the Second Order Section biquad UGen.
// This is a composite pseudo Ugen. BLowPass4 is built by cascading 2 SOS sections.
type BLowPass4 struct {
	// In is the input signal.
	In Input
	// Freq is frequency in Hz.
	Freq Input
	// RQ is the reciprocal of Q, bandwidth / cutoff.
	RQ Input
}

func (blp *BLowPass4) defaults() {
	if blp.In == nil {
		panic("BLowPass4 needs an input")
	}
	if blp.Freq == nil {
		blp.Freq = C(1200)
	}
	if blp.RQ == nil {
		blp.RQ = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (blp BLowPass4) Rate(rate int8) Input {
	CheckRate(rate)
	(&blp).defaults()

	// coeffs := blowpassCoeffs(blp.Freq, blp.RQ)

	return NewUgenInput("BLowPass4", rate, 0, 1, blp.In, blp.Freq, blp.RQ)
}

// See https://github.com/supercollider/supercollider/blob/772bdd6946a7253390ba0b1b23eb3adc0acb97ea/SCClassLibrary/Common/Audio/BEQSuite.sc
func blowpassCoeffs(freq, rq Input) []Input {
	if freq == nil {
		freq = C(1200)
	}
	if rq == nil {
		rq = C(1)
	}
	var (
		sd    = I(SampleDur{})
		w0    = sd.Mul(Pi.Mul(C(2)).Mul(freq))
		cosw0 = w0.Cos()
		i     = C(1).Add(cosw0.Neg())
		alpha = w0.Sin().Mul(C(0.5)).Mul(rq)
		b0rz  = alpha.Add(C(1)).Reciprocal()
		a0    = i.Mul(C(0.5)).Mul(b0rz)
		a1    = i.Mul(b0rz)
		b1    = cosw0.Mul(C(2)).Mul(b0rz)
		b2    = alpha.Neg().Add(C(1)).Mul(b0rz.Neg())
	)
	return []Input{a0, a1, a0, b1, b2}
}
