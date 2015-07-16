package sc

import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"
import "testing"

func TestLFNoise1(t *testing.T) {
	def := NewSynthdef("LFNoise1Example", func(p Params) Ugen {
		start, end, dur, done := C(1000), C(10000), C(10), 0
		bus, gain := C(0), C(0.25)
		freq := XLine{start, end, dur, done}.Rate(KR)
		sig := LFNoise1{freq}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	same, err := def.Compare(`{
        var freq = XLine.kr(1000, 10000, 10);
        Out.ar(0, LFNoise1.ar(freq, 0.25));
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
