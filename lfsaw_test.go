package sc

import "testing"

func TestLFSaw(t *testing.T) {
	def := NewSynthdef("LFSawExample", func(p Params) Ugen {
		lfoFreq, lfoPhase := C(4), C(0)
		bus, gain := C(0), C(0.1)
		freq := LFSaw{lfoFreq, lfoPhase}.Rate(KR).MulAdd(C(200), C(400))
		sig := LFSaw{freq, C(0)}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	same, err := def.Compare(`{
        var freq = LFSaw.kr(4, 0, 200, 400);
        Out.ar(0, LFSaw.ar(freq, 0, 0.1));
    }`)
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
