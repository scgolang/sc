package ugens

import (
	"testing"
)

func TestOut(t *testing.T) {
	out := Out.Ar(0, SinOsc.Ar())
	if out == nil {
		t.Fatalf("out was nil")
	}
	inputs := out.Inputs()
	numInputs := len(inputs)
	if numInputs != 2 {
		t.Fatalf("expected 2 inputs but got %d", numInputs)
	}
}
