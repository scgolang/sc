package sc

import (
	"testing"
)

func TestBlip(t *testing.T) {
	def := NewSynthdef("BlipExample", func(p Params) Ugen {
		start, end, dur, done := C(20000), C(200), C(6), 0
		freq := XLine{start, end, dur, done}.Rate(KR)
		bus, harms, gain := C(0), C(100), C(0.2)
		sig := Blip{freq, harms}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	same, err := def.Compare(`{
        var freq = XLine.kr(20000, 200, 6);
        var harms = 100;
        var mul = 0.2;
        Out.ar(0, Blip.ar(freq, harms, mul));
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef is different from sclang version")
	}
}
