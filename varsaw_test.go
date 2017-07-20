package sc

import (
	"testing"
)

func TestVarSaw(t *testing.T) {
	const defname = "VarSawTest"

	// Out.ar(0, VarSaw.ar(LFPulse.kr(3, 0, 0.3, 200, 200), 0, 0.2, 0.1));
	compareAndWriteStructure(t, defname, NewSynthdef(defname, func(p Params) Ugen {
		pulse := K(LFPulse{
			Freq:   C(3),
			IPhase: C(0),
			Width:  C(0.3),
		})
		saw := A(VarSaw{
			Freq:   pulse.MulAdd(C(200), C(200)),
			IPhase: C(0),
			Width:  C(0.2),
		})
		return Out{
			Bus:      C(0),
			Channels: saw.Mul(C(0.1)),
		}.Rate(AR)
	}))
}
