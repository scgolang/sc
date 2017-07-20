package sc

import (
	"testing"
)

func TestMedian(t *testing.T) {
	const defName = "MedianTest"

	// Out.ar(0, LeakDC.ar(Median.ar(31, WhiteNoise.ar(0.1) + SinOsc.ar(800,0,0.1)), 0.9));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		sine := A(SinOsc{Freq: C(800)}).Mul(C(0.1))
		return Out{
			Bus: C(0),
			Channels: A(LeakDC{
				In: A(Median{
					Length: C(31),
					In:     A(WhiteNoise{}).MulAdd(C(0.1), sine),
				}),
				Coeff: C(0.9),
			}),
		}.Rate(AR)
	}))
}
