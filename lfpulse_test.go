package sc

import (
	"testing"
)

func TestLFPulse(t *testing.T) {
	def := NewSynthdef("LFPulseTest", func(p Params) Ugen {
		lfoFreq, lfoPhase, lfoWidth := C(3), C(0), C(0.3)
		bus, gain := C(0), C(0.1)
		freq := LFPulse{lfoFreq, lfoPhase, lfoWidth}.Rate(KR).MulAdd(C(200), C(200))
		iphase, width := C(0), C(0.2)
		sig := LFPulse{freq, iphase, width}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	same, err := def.CompareToFile("fixtures/LFPulseTest.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
