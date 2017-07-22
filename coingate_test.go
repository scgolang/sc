package sc

import (
	"testing"
)

func TestCoinGate(t *testing.T) {
	const defName = "CoinGateTest"

	// arg out=0, prob=0.5;
	// var trig;
	// trig = CoinGate.kr(prob, Impulse.kr(10));
	// Out.ar(out,
	//     SinOsc.ar(
	//         TRand.kr(300.0, 400.0, trig), 0, 0.2
	//     )
	// )
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			out  = p.Add("out", 0)
			prob = p.Add("prob", 0.5)
		)
		trig := K(CoinGate{
			Prob: prob,
			In: K(Impulse{
				Freq: C(10),
			}),
		})
		sine := A(SinOsc{
			Freq: K(TRand{
				Lo:   C(300),
				Hi:   C(400),
				Trig: trig,
			}),
		})
		return Out{
			Bus:      out,
			Channels: sine.Mul(C(0.2)),
		}.Rate(AR)
	}))
}
