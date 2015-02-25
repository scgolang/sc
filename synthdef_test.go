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
	def := NewSynthdef("SineTone", func(params Params) UgenNode {
		return Out.Ar(0, SinOsc.Ar(440))
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
}

func ExampleNewSynthdef() {
	NewSynthdef("SineTone", func(params Params) UgenNode {
		return Out.Ar(0, SinOsc.Ar(440))
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SineTone","constants":[440,0],"initialParamValues":[],"paramNames":[],"ugens":[{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":1}],"outputs":[{"rate":2}]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":1},{"ugenIndex":0,"outputIndex":0}],"outputs":[]}],"variants":[]}
}

func ExampleNewSynthdefSineTone2() {
	NewSynthdef("SineTone2", func(params Params) UgenNode {
		return Out.Ar(0, SinOsc.Ar(440, SinOsc.Ar(0.1, 0)).Mul(0.5))
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SineTone2","constants":[0.1,0,440,0.5],"initialParamValues":[],"paramNames":[],"ugens":[{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":1}],"outputs":[{"rate":2}]},{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":2},{"ugenIndex":0,"outputIndex":0}],"outputs":[{"rate":2}]},{"name":"BinaryOpUGen","rate":2,"specialIndex":2,"inputs":[{"ugenIndex":1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":3}],"outputs":[{"rate":2}]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":1},{"ugenIndex":2,"outputIndex":0}],"outputs":[]}],"variants":[]}
}

func ExampleNewSynthdefParams() {
	NewSynthdef("SineTone4", func(params Params) UgenNode {
		freq := params.Add("freq").SetDefault(440)
		return Out.Ar(0, SinOsc.Ar(freq))
	})
}
