package sc

import (
	"math"
	"testing"
)

func TestSinOscFB(t *testing.T) {
	const defName = "SinOscFBTest"

	// Out.ar(0, SinOscFB.ar(100*SinOscFB.ar(MouseY.kr(1,1000,'exponential'))+200,MouseX.kr(0.5pi,pi))*0.1);
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		modulator := A(SinOscFB{
			Freq: K(MouseY{
				Min:  C(1),
				Max:  C(1000),
				Warp: WarpExp,
			}),
		})
		feedback := K(MouseX{
			Min: C(0.5 * math.Pi),
			Max: C(math.Pi),
		})
		return Out{
			Bus: C(0),
			Channels: A(SinOscFB{
				Freq:     modulator.MulAdd(C(100), C(200)),
				Feedback: feedback,
			}).Mul(C(0.1)),
		}.Rate(AR)
	}))
}
