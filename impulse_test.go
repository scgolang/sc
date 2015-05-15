package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"
import "testing"

func TestImpulse(t *testing.T) {
	def := NewSynthdef("ImpulseExample", func(p *Params) UgenNode {
		start, end, dur, done := C(800), C(100), C(5), 0
		freq := XLine{start, end, dur, done}.Rate(KR)
		bus, phase, gain := C(0), C(0), C(0.5)
		sig := Impulse{freq, phase}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, "ImpulseExample", def)
}
