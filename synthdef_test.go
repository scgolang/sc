package sc

import (
	"encoding/json"
	"fmt"
	. "github.com/briansorahan/sc/types"
	. "github.com/briansorahan/sc/ugens"
	"os"
	"testing"
)

func TestReadSynthDef(t *testing.T) {
	// read a synthdef file created by sclang
	f, err := os.Open("SineTone.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	synthDef, err := readsynthdef(f)
	if err != nil {
		t.Fatal(err)
	}
	// check the name
	if synthDef.Name != "SineTone" {
		t.Fatal(fmt.Errorf("wrong synthdef name"))
	}
	enc := json.NewEncoder(os.Stdout)
	if err = enc.Encode(synthDef); err != nil {
		t.Fatal(err)
	}
}

func TestNewSynthDef(t *testing.T) {
	def := NewSynthdef("SineTone", func(params Params) UgenNode {
		//sc-> Out.ar(0, SinOsc.ar(440, SinOsc.ar(0.1), 0.5));
		return Out.Ar(0, SinOsc.Ar(440, SinOsc.Ar(0.1), 0.5))
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(def); err != nil {
		t.Fatal(err)
	}
}
