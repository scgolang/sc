package sc

import (
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
	synthDef.Dump(os.Stdout)
}

func TestNewSynthDef(t *testing.T) {
	def := NewSynthdef("SineTone", func(params Params) UgenNode {
		return Out.Ar(0, SinOsc.Ar(440))
	})
	if def == nil {
		t.Fatalf("nil synthdef")
	}
}
