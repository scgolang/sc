package sc

import (
	"testing"
)

func TestSweep(t *testing.T) {
	defName := "SweepTest"

	def := NewSynthdef(defName, func(p Params) Ugen {
		return Out{
			Bus:      C(0),
			Channels: LFPulse{}.Rate(AR).Mul(Sweep{}.Rate(AR)),
		}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
