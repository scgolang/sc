package sc

// LFSaw a non-band-limited sawtooth oscillator
// output ranges from -1 to +1
type LFSaw struct {
	// Freq frequency in Hz
	Freq Input
	// Iphase initial phase offset in cycles:
	// for efficiency this is in the rage [0, 2]
	Iphase Input
}

func (lfsaw *LFSaw) defaults() {
	if lfsaw.Freq == nil {
		lfsaw.Freq = C(440)
	}
	if lfsaw.Iphase == nil {
		lfsaw.Iphase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (lfsaw LFSaw) Rate(rate int8) Input {
	CheckRate(rate)
	(&lfsaw).defaults()
	return NewInput("LFSaw", rate, 0, 1, lfsaw.Freq, lfsaw.Iphase)
}

func defLFSaw(params Params) Ugen {
	var (
		add    = params.Add("add", 0)
		freq   = params.Add("freq", 440)
		iphase = params.Add("iphase", 0)
		mul    = params.Add("mul", 1)
		out    = params.Add("out", 0)
	)
	return Out{
		Bus: out,
		Channels: LFSaw{
			Freq:   freq,
			Iphase: iphase,
		}.Rate(AR).MulAdd(mul, add),
	}.Rate(AR)
}
