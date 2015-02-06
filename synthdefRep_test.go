package sc

import (
	"fmt"
	"os"
	"testing"
)

func TestReadSynthDef(t *testing.T) {
	// read a synthdef file created by sclang
	f, err := os.Open("SineTone.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	synthDef, err := readSynthdefRep(f)
	if err != nil {
		t.Fatal(err)
	}
	// check the name
	if synthDef.Name != "SineTone" {
		t.Fatal(fmt.Errorf("wrong synthdef name"))
	}
	synthDef.Dump(os.Stdout)
}

// FIXME
//
// func ExampleNewSynthDef() {
// 	NewSynthdef("SineTone", func() (*Ugen, error) {
// 		sinOsc, err := Ar("SinOsc", float32(440))
// 		if err != nil {
// 			return nil, err
// 		}
// 		root, err := Ar("Out", float32(0), sinOsc)
// 		return root, err
// 	}).Dump(os.Stdout)
// 	// Output:
// 	// {"name":"SineTone","constants":[440,0],"initialParamValues":[],"paramNames":[],"ugens":[{"name":"SinOsc","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":0},{"ugenIndex":-1,"outputIndex":1}],"outputs":[{"rate":2}]},{"name":"Out","rate":2,"specialIndex":0,"inputs":[{"ugenIndex":-1,"outputIndex":1},{"ugenIndex":0,"outputIndex":0}],"outputs":[]}],"variants":[]}
// }
