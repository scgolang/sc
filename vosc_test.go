package sc

import (
	"testing"
)

func TestVOsc(t *testing.T) {
	defName := "VOscTest"

	// arg bufnum = 0;
	// Out.ar(out, VOsc.ar(bufnum, XLine.kr(2000,200), 0, 0.5));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			bufnum = p.Add("bufnum", 0)
		)
		return Out{
			Bus: C(0),
			Channels: A(VOsc{
				BufNum: bufnum,
				Freq: K(XLine{
					Start: C(2000),
					End:   C(200),
				}),
			}).Mul(C(0.5)),
		}.Rate(AR)
	}))
}
