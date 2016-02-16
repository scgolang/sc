package sc

import (
	"testing"
)

func TestImpulse(t *testing.T) {
	def := NewSynthdef("ImpulseExample", func(p Params) Ugen {
		start, end, dur, done := C(800), C(100), C(5), 0
		freq := XLine{start, end, dur, done}.Rate(KR)
		bus, phase, gain := C(0), C(0), C(0.5)
		sig := Impulse{freq, phase}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/ImpulseExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
