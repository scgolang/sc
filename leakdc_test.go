package sc

import (
	"testing"
)

func TestLeakDC(t *testing.T) {
	defName := "LeakDCTest"

	def := NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus: C(0),
			Channels: LeakDC{
				In: LFPulse{
					Freq:   C(800),
					IPhase: C(0.5),
					Width:  C(0.5),
				}.Rate(AR).Mul(C(0.5)),
				Coeff: C(0.995),
			}.Rate(AR),
		}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
