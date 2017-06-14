package sc

import (
	"fmt"
	"os"
	"testing"
)

func TestReadSynthdef(t *testing.T) {
	// read a synthdef file created by sclang
	f, err := os.Open("testdata/SineTone.scsyndef")
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

func ExampleNewSynthdef() {
	NewSynthdef("SineTone", func(p Params) Ugen {
		bus := C(0)
		sine := SinOsc{}.Rate(AR)
		return Out{bus, sine}.Rate(AR)
	}).WriteJSON(os.Stdout)
	// Output:
	// {"name":"SineTone","constants":[440,0],"initialParamValues":[],"paramNames":[],"ugens":[{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":1}],"outputs":[2]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":1},{"ugenIndex":0,"outputIndex":0}],"outputs":[]}],"variants":[]}
}
