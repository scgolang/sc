package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"
import "testing"

func TestLFNoise1(t *testing.T) {
	def := NewSynthdef("LFNoise1Example", func(p *Params) UgenNode {
		start, end, dur, done := C(1000), C(10000), C(10), 0
		bus, gain := C(0), C(0.25)
		freq := XLine{start, end, dur, done}.Rate(KR)
		sig := LFNoise1{freq}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, "LFNoise1Example", def);
}
