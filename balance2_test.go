package sc

import (
	"testing"
)

func TestBalance2(t *testing.T) {
	defName := "Balance2Test"
	def := NewSynthdef(defName, func(p Params) Ugen {
		bus, gain := C(0), C(0.1)
		l, r := LFSaw{Freq: C(44)}.Rate(AR), Pulse{Freq: C(33)}.Rate(AR)
		pos := FSinOsc{Freq: C(0.5)}.Rate(KR)
		sig := Balance2{L: l, R: r, Pos: pos, Level: gain}.Rate(AR)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
