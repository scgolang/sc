package sc

import (
	"testing"
)

func TestPulseDivider(t *testing.T) {
	defName := "PulseDividerTest"

	compareAndWrite(t, defName, NewSynthdef(defName, func(params Params) Ugen {
		var (
			out    = params.Add("out", 0)
			p      = Impulse{Freq: C(8)}.Rate(AR)
			decay1 = Decay2{In: p, Attack: C(0.005), Decay: C(0.1)}.Rate(AR)
			decay2 = Decay2{In: PulseDivider{Trig: p, Div: C(4)}.Rate(AR), Attack: C(0.005), Decay: C(0.5)}.Rate(AR)
			b      = SinOsc{Freq: C(600)}.Rate(AR).Mul(decay2)
			a      = SinOsc{Freq: C(1200)}.Rate(AR).MulAdd(decay1, b)
		)
		return Out{
			Bus:      out,
			Channels: a.Mul(C(0.4)),
		}.Rate(AR)
	}))
}
