package sc

import . "github.com/briansorahan/sc/types"
import . "github.com/briansorahan/sc/ugens"
import "os"
import "testing"

func TestFSinOsc(t *testing.T) {
	def := NewSynthdef("FSinOscExample", func(p *Params) UgenNode {
		line := XLine{C(4), C(401), C(8), 0}.Rate(KR)
		sin1 := FSinOsc{line, C(0)}.Rate(AR).MulAdd(C(200), C(800))
		sin2 := FSinOsc{Freq:sin1}.Rate(AR).Mul(C(0.2))
		bus := C(0)
		return Out{bus, sin2}.Rate(AR)
	})
	f, err := os.Create("FSinOscExample.gosyndef")
	if err != nil {
		t.Fatal(err)
	}
	err = def.Write(f)
	if err != nil {
		t.Fatal(err)
	}
	same, err := def.CompareToFile("FSinOscExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}
