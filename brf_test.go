package sc

import . "github.com/briansorahan/sc/types"
import . "github.com/briansorahan/sc/ugens"
import "os"
import "testing"

func TestBRF(t *testing.T) {
	def := NewSynthdef("BRFExample", func(p *Params) UgenNode {
		line := XLine{C(0.7), C(300), C(20), 0}.Rate(KR)
		saw := Saw{C(200)}.Rate(AR).Mul(C(0.5))
		sine := FSinOsc{line, C(0)}.Rate(KR).MulAdd(C(3800), C(4000))
		bpf := BRF{saw, sine, C(0.3)}.Rate(AR)
		return Out{C(0), bpf}.Rate(AR)
	})
	f, err := os.Create("BRFExample.gosyndef")
	if err != nil {
		t.Fatal(err)
	}
	err = def.Write(f)	
	if err != nil {
		t.Fatal(err)
	}
	same, err := def.CompareToFile("BRFExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef is not the same as sclang version")
	}
}
