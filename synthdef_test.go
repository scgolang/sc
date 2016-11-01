package sc

import (
	"fmt"
	"os"
	"testing"
)

func TestReadSynthdef(t *testing.T) {
	// read a synthdef file created by sclang
	f, err := os.Open("fixtures/SineTone.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	synthDef, err := ReadSynthdef(f)
	if err != nil {
		t.Fatal(err)
	}
	// check the name
	if synthDef.Name != "SineTone" {
		t.Fatal(fmt.Errorf("wrong synthdef name"))
	}
}

func TestAllpassExample(t *testing.T) {
	def := NewSynthdef("AllpassExample", func(p Params) Ugen {
		noise := WhiteNoise{}.Rate(AR).Mul(C(0.1))

		line := XLine{
			Start: C(0.0001),
			End:   C(0.01),
			Dur:   C(20),
			Done:  0,
		}.Rate(KR)

		all := AllpassC{
			In:       noise,
			MaxDelay: C(0.01),
			Delay:    line,
			Decay:    C(0.2),
		}.Rate(AR)

		return Out{C(0), all}.Rate(AR)
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
	f, err := os.Create("AllpassExample.gosyndef")
	if err != nil {
		t.Fatal(err)
	}
	err = def.Write(f)
	if err != nil {
		t.Fatal(err)
	}
	same, err := def.CompareToFile("fixtures/AllpassExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang-generated version")
	}
}

func ExampleNewSynthdef() {
	NewSynthdef("SineTone", func(p Params) Ugen {
		bus := C(0)
		sine := SinOsc{}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SineTone","constants":[440,0],"initialParamValues":[],"paramNames":[],"ugens":[{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":1}],"outputs":[2]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":1},{"ugenIndex":0,"outputIndex":0}],"outputs":[]}],"variants":[]}
}
