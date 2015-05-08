package sc

import . "github.com/briansorahan/sc/types"
import . "github.com/briansorahan/sc/ugens"
import "os"
import "testing"

func TestBlip(t *testing.T) {
	def := NewSynthdef("BlipExample", func(p *Params) UgenNode {
		start, end, dur, done := C(20000), C(200), C(6), 0
		freq := XLine{start, end, dur, done}.Rate(KR)
		bus, harms, gain := C(0), C(100), C(0.2)
		sig := Blip{freq, harms}.Rate(AR).Mul(gain)
		return Out{bus, sig}.Rate(AR)
	})
	f, err := os.Create("BlipExample.gosyndef")
	if err != nil {
		t.Fatal(err)
	}
	err = def.Write(f)	
	if err != nil {
		t.Fatal(err)
	}
	same, err := def.CompareToFile("BlipExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef is not the same as sclang version")
	}
}
