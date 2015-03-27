package sc

import (
	"fmt"
	. "github.com/briansorahan/sc/types"
	. "github.com/briansorahan/sc/ugens"
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

func TestNewSynthdef(t *testing.T) {
	def := NewSynthdef("SineTone", func(params *Params) UgenNode {
		bus, freq, phase := C(0), C(440), C(0)
		sine := SinOsc{freq, phase}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
}

func TestCompareToFile(t *testing.T) {
	def := NewSynthdef("SineTone", func(params *Params) UgenNode {
		bus, freq := C(0), C(440)
		sine := SinOsc{Freq:freq}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}

	f, err := os.Create("SineTone.gosyndef")
	if err != nil {
		t.Fatal(err)
	}
	err = def.Write(f)
	if err != nil {
		t.Fatal(err)
	}

	same, err := def.CompareToFile("SineTone.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang-generated version")
	}
}

func TestSynthdefEnvgen(t *testing.T) {
	def := NewSynthdef("Envgen1", func(params *Params) UgenNode {
		bus := C(0)
		attack, release := C(0.01), C(1)
		level, curveature := C(1), C(-4)
		perc := EnvPerc{attack, release, level, curveature}
		gate, levelScale, levelBias, timeScale := C(1), C(1), C(0), C(1)
		ampEnv := EnvGen{perc, gate, levelScale, levelBias, timeScale, FreeEnclosing}.Rate(KR)
		noise := PinkNoise{}.Rate(AR).Mul(ampEnv)
		return Out{bus, noise}.Rate(AR)
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
	same, err := def.CompareToFile("Envgen1.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang-generated version")
	}
}

func TestSimpleMulti(t *testing.T) {
	def := NewSynthdef("SimpleMulti", func(p *Params) UgenNode {
		bus, freq := C(0), Multi(C(440), C(441))
		sine := SinOsc{Freq:freq}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
	same, err := def.CompareToFile("SimpleMulti.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang-generated version")
	}
}

func TestCascade(t *testing.T) {
	// var mod1 = SinOsc.ar([440, 441]);
	// var mod2 = SinOsc.ar(mod1);
	// Out.ar(0, SinOsc.ar(mod2));
	def := NewSynthdef("Cascade", func(p *Params) UgenNode {
		bus := C(0)
		freq := Multi(C(440), C(441))
		mod1 := SinOsc{Freq: freq}.Rate(AR)
		mod2 := SinOsc{Freq: mod1}.Rate(AR)
		return Out{bus, SinOsc{Freq: mod2}.Rate(AR)}.Rate(AR)
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
	same, err := def.CompareToFile("Cascade.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang-generated version")
	}
}

func ExampleNewSynthdef() {
	NewSynthdef("SineTone", func(params *Params) UgenNode {
		bus := C(0)
		sine := SinOsc{}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SineTone","constants":[440,0],"initialParamValues":[],"paramNames":[],"ugens":[{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":1}],"outputs":[{"rate":2}]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":1},{"ugenIndex":0,"outputIndex":0}],"outputs":[]}],"variants":[]}
}

func ExampleNewSynthdefSineTone2() {
	NewSynthdef("SineTone2", func(params *Params) UgenNode {
		bus := C(0)
		freq:= C(440)
		phase := SinOsc{Freq: C(0.1)}.Rate(AR)
		out := SinOsc{freq, phase}.Rate(AR).Mul(C(0.5))
		return Out{bus, out}.Rate(AR)
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SineTone2","constants":[0.1,0,440,0.5],"initialParamValues":[],"paramNames":[],"ugens":[{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":1}],"outputs":[{"rate":2}]},{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":2},{"ugenIndex":0,"outputIndex":0}],"outputs":[{"rate":2}]},{"name":"BinaryOpUGen","rate":2,"specialIndex":2,"inputs":[{"ugenIndex":1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":3}],"outputs":[{"rate":2}]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":1},{"ugenIndex":2,"outputIndex":0}],"outputs":[]}],"variants":[]}
}

func ExampleNewSynthdefParams() {
	NewSynthdef("SineTone4", func(params *Params) UgenNode {
		freq := params.Add("freq", 440)
		bus, sine := C(0), SinOsc{freq, C(0)}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SineTone4","constants":[0],"initialParamValues":[440],"paramNames":[{"Name":"freq","Index":0}],"ugens":[{"name":"Control","rate":1,"specialIndex":0,"inputs":[],"outputs":[{"rate":1}]},{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":0,"outputIndex":0},{"ugenIndex":-1,"outputIndex":0}],"outputs":[{"rate":2}]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":1,"outputIndex":0}],"outputs":[]}],"variants":[]}
}

func ExampleSynthdefParams2() {
	NewSynthdef("SawTone1", func(params *Params) UgenNode {
		freq := params.Add("freq", 440)
		cutoff, q := params.Add("cutoff", 1200), params.Add("q", 0.5)
		bus := C(0)
		saw := Saw{freq}.Rate(AR)
		out := RLPF{saw, cutoff, q}.Rate(AR)
		return Out{bus, out}.Rate(AR)
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SawTone1","constants":[0],"initialParamValues":[440,1200,0.5],"paramNames":[{"Name":"freq","Index":0},{"Name":"cutoff","Index":1},{"Name":"q","Index":2}],"ugens":[{"name":"Control","rate":1,"specialIndex":0,"inputs":[],"outputs":[{"rate":1},{"rate":1},{"rate":1}]},{"name":"Saw","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":0,"outputIndex":0}],"outputs":[{"rate":2}]},{"name":"RLPF","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":1,"outputIndex":0},{"ugenIndex":0,"outputIndex":1},{"ugenIndex":0,"outputIndex":2}],"outputs":[{"rate":2}]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":2,"outputIndex":0}],"outputs":[]}],"variants":[]}
}
