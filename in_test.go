package sc

import "testing"

func TestIn(t *testing.T) {
	const name = "InTest"

	compareAndWriteStructure(t, name, NewSynthdef(name, func(params Params) Ugen {
		var (
			out = params.Add("out", 0)
			in  = params.Add("in", 0)
		)
		return Out{
			Bus: out,
			Channels: SinOsc{
				Freq: In{
					Bus:         in,
					NumChannels: 2,
				}.Rate(KR),
			}.Rate(AR).Mul(C(0.1)),
		}.Rate(AR)
	}))
}
