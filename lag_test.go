package sc

import (
	"testing"
)

func TestLag(t *testing.T) {
	defName := "LagTest"

	// Out.ar(0, SinOsc.ar(
	//         Lag.kr(
	//                 LFPulse.kr(4, 0, 0.5, 50, 400),
	//                 Line.kr(0, 1, 15)
	//         ),
	//         0,
	//         0.3
	// ));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		pulse := K(LFPulse{
			Freq:   C(4),
			IPhase: C(0),
			Width:  C(0.5),
		})
		line := K(Line{
			Start: C(0),
			End:   C(1),
			Dur:   C(15),
		})
		return Out{
			Bus: C(0),
			Channels: A(SinOsc{
				Freq: K(Lag{
					In:      pulse.MulAdd(C(50), C(400)),
					LagTime: line,
				}),
			}).Mul(C(0.3)),
		}.Rate(AR)
	}))
}
