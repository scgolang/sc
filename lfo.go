package sc

func defLFO(params Params) Ugen {
	var (
		freq   = params.Add("freq", 1)
		out    = params.Add("out", 0)
		saw    = LFSaw{Freq: freq}.Rate(KR)
		sine   = SinOsc{Freq: freq}.Rate(KR)
		square = LFPulse{Freq: freq}.Rate(KR)
	)
	return Out{
		Bus:      out,
		Channels: Multi(saw, sine, square),
	}.Rate(KR)
}
