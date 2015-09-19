package sc

import (
	"fmt"
	"os"
	"testing"
)

func TestReadSynthdef(t *testing.T) {
	// read a synthdef file created by sclang
	f, err := os.Open("SineTone.scsyndef")
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
	same, err := def.CompareToFile("AllpassExample.scsyndef")
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

func ExampleNewSynthdefSineTone2() {
	NewSynthdef("SineTone2", func(p Params) Ugen {
		bus := C(0)
		freq := C(440)
		phase := SinOsc{Freq: C(0.1)}.Rate(AR)
		out := SinOsc{freq, phase}.Rate(AR).Mul(C(0.5))
		return Out{bus, out}.Rate(AR)
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SineTone2","constants":[0.1,0,440,0.5],"initialParamValues":[],"paramNames":[],"ugens":[{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":1}],"outputs":[2]},{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":2},{"ugenIndex":0,"outputIndex":0}],"outputs":[2]},{"name":"BinaryOpUGen","rate":2,"specialIndex":2,"inputs":[{"ugenIndex":1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":3}],"outputs":[2]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":1},{"ugenIndex":2,"outputIndex":0}],"outputs":[]}],"variants":[]}
}

func ExampleNewSynthdefParams() {
	NewSynthdef("SineTone4", func(p Params) Ugen {
		freq := p.Add("freq", 440)
		bus, sine := C(0), SinOsc{freq, C(0)}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SineTone4","constants":[0],"initialParamValues":[440],"paramNames":[{"Name":"freq","Index":0}],"ugens":[{"name":"Control","rate":1,"specialIndex":0,"inputs":[],"outputs":[1]},{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":0,"outputIndex":0},{"ugenIndex":-1,"outputIndex":0}],"outputs":[2]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":1,"outputIndex":0}],"outputs":[]}],"variants":[]}
}

func ExampleSynthdefParams2() {
	NewSynthdef("SawTone1", func(p Params) Ugen {
		freq := p.Add("freq", 440)
		cutoff, q := p.Add("cutoff", 1200), p.Add("q", 0.5)
		bus := C(0)
		saw := Saw{freq}.Rate(AR)
		out := RLPF{saw, cutoff, q}.Rate(AR)
		return Out{bus, out}.Rate(AR)
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SawTone1","constants":[0],"initialParamValues":[440,1200,0.5],"paramNames":[{"Name":"freq","Index":0},{"Name":"cutoff","Index":1},{"Name":"q","Index":2}],"ugens":[{"name":"Control","rate":1,"specialIndex":0,"inputs":[],"outputs":[1,1,1]},{"name":"Saw","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":0,"outputIndex":0}],"outputs":[2]},{"name":"RLPF","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":1,"outputIndex":0},{"ugenIndex":0,"outputIndex":1},{"ugenIndex":0,"outputIndex":2}],"outputs":[2]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":2,"outputIndex":0}],"outputs":[]}],"variants":[]}
}
