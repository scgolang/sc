package sc

import (
	"testing"
)

func TestLFPar(t *testing.T) {
	const defName = "LFParTest"

	// Out.ar(0, LFPar.ar(XLine.kr(100,8000,30),0,0.1));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(LFPar{
				Freq: K(XLine{
					Start: C(100),
					End:   C(8000),
					Dur:   C(30),
				}),
			}).Mul(C(0.1)),
		}.Rate(AR)
	}))
}
