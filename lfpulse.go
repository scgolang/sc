package sc

// LFPulse a non-band-limited pulse oscillator
type LFPulse struct {
	// Freq in Hz
	Freq Input
	// Iphase initial phase offset in cycles (0..1)
	Iphase Input
	// Width pulse width duty cycle from 0 to 1
	Width Input
}

func (lfpulse *LFPulse) defaults() {
	if lfpulse.Freq == nil {
		lfpulse.Freq = C(440)
	}
	if lfpulse.Iphase == nil {
		lfpulse.Iphase = C(0)
	}
	if lfpulse.Width == nil {
		lfpulse.Width = C(0.5)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (lfpulse LFPulse) Rate(rate int8) Input {
	CheckRate(rate)
	(&lfpulse).defaults()
	return UgenInput("LFPulse", rate, 0, 1, lfpulse.Freq, lfpulse.Iphase, lfpulse.Width)
}

func init() {
	if err := RegisterSynthdef("lfpulse", func(params Params) Ugen {
		var (
			add    = params.Add("add", 0)
			freq   = params.Add("freq", 440)
			iphase = params.Add("iphase", 0)
			mul    = params.Add("mul", 1)
			out    = params.Add("out", 0)
			width  = params.Add("width", 0)
		)
		return Out{
			Bus: out,
			Channels: LFPulse{
				Freq:   freq,
				Iphase: iphase,
				Width:  width,
			}.Rate(AR).MulAdd(mul, add),
		}.Rate(AR)
	}); err != nil {
		panic(err)
	}
}
