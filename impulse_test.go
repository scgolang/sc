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
	same, err := def.Compare(`{
		var freq = XLine.kr(800, 100, 5);
		var gain = 0.5;
		var phase = 0.0;
		var sig = Impulse.ar(freq, phase, gain);
		Out.ar(0, sig);
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
