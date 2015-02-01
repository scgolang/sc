package gosc

import (
	"os"
	"testing"
)

func TestReadSynthDef(t *testing.T) {
	f, err := os.Open("SineTone.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ReadSynthDef(f)
	if err != nil {
		t.Fatal(err)
	}
}
