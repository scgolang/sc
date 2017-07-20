package sc

import (
	"testing"
)

func TestRLPF(t *testing.T) {
	const defName = "RLPFTest"

	// Out.ar(0, RLPF.ar(Saw.ar(200, 0.1), FSinOsc.kr(XLine.kr(0.7, 300, 20), 0, 3600, 4000), 0.2));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		saw := A(Saw{
			Freq: C(200),
		})
		sine := K(FSinOsc{
			Freq: K(XLine{
				Start: C(0.7),
				End:   C(300),
				Dur:   C(20),
			}),
		})
		return Out{
			Bus: C(0),
			Channels: A(RLPF{
				In:   saw.Mul(C(0.1)),
				Freq: sine.MulAdd(C(3600), C(4000)),
				RQ:   C(0.2),
			}),
		}.Rate(AR)
	}))
}
