package gosc

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
	synthDef, err := ReadSynthDef(f)
	if err != nil {
		t.Fatal(err)
	}
	// get the name
	if synthDef.Name() != "SineTone" {
		t.Fatal(fmt.Errorf("wrong synthdef name"))
	}
	synthDef.Dump(os.Stdout)
}
