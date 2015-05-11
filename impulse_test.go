package sc

import . "github.com/briansorahan/sc/types"
import . "github.com/briansorahan/sc/ugens"
import "os"
import "testing"

func TestImpulse(t *testing.T) {
	def := NewSynthdef("ImpulseExample", func(p *Params) UgenNode {
		start, end, dur, done := C(800), C(100), C(5), 0
		freq := XLine{start, end, dur, done}.Rate(KR)
		bus, phase, gain := C(0), C(0), C(0.5)
		sig := Impulse{freq, phase}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	f, err := os.Create("ImpulseExample.gosyndef")
	if err != nil {
		t.Fatal(err)
	}
	err = def.Write(f)	
	if err != nil {
		t.Fatal(err)
	}
	same, err := def.CompareToFile("ImpulseExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef is not the same as sclang version")
	}
}
