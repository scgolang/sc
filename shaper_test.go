package sc

import (
	"testing"
)

func TestShaper(t *testing.T) {
	const defName = "ShaperTest"

	// arg bufnum = 0;
	// Out.ar(0, Shaper.ar(bufnum, SinOsc.ar(440, 0.5, Line.kr(0,0.9,6)));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			bufnum = p.Add("bufnum", 0)
		)
		sine := A(SinOsc{
			Freq:  C(440),
			Phase: C(0.5),
		})
		line := K(Line{
			Start: C(0),
			End:   C(0.9),
			Dur:   C(6),
		})
		return Out{
			Bus: C(0),
			Channels: A(Shaper{
				BufNum: bufnum,
				In:     sine.Mul(line),
			}),
		}.Rate(AR)
	}))
}
