package sc

import (
	"testing"
)

func TestBall(t *testing.T) {
	const defName = "BallTest"

	// var sf = LFNoise0.ar(MouseX.kr(1, 100, 1));
	// var g  = MouseY.kr(0.1, 10, 1);
	// var f  = Ball.ar(sf, g, 0.01, 0.01);
	// f = f * 140 + 500;
	// SinOsc.ar(f, 0, 0.2)
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		sf := A(LFNoise{
			Interpolation: NoiseStep,
			Freq: K(MouseX{
				Min:  C(1),
				Max:  C(100),
				Warp: WarpExp,
			}),
		})
		g := K(MouseY{
			Min:  C(0.1),
			Max:  C(10),
			Warp: WarpExp,
		})
		f := A(Ball{
			In:       sf,
			Gravity:  g,
			Damp:     C(0.01),
			Friction: C(0.01),
		})
		f = f.MulAdd(C(140), C(500))

		return Out{
			Bus:      C(0),
			Channels: SinOsc{Freq: f}.Rate(AR).Mul(C(0.2)),
		}.Rate(AR)
	}))
}
