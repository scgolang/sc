package sc

import (
	"testing"
)

func TestLFGauss(t *testing.T) {
	const defName = "LFGaussTest"

	// Out.ar(0, LFGauss.ar(0.01, SampleDur.ir * MouseX.kr(10, 3000, 1)) * 0.2);
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: A(LFGauss{
				Duration: C(0.01),
				Width: SampleDur{}.Rate(IR).Mul(K(MouseX{
					Min:  C(10),
					Max:  C(3000),
					Warp: WarpExp,
				})),
			}).Mul(C(0.2)),
		}.Rate(AR)
	}))
}
