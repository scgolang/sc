package sc

import (
	"testing"
)

func TestOut(t *testing.T) {
	var (
		bus       = C(0)
		sin       = SinOsc{}.Rate(AR)
		out       = Out{bus, sin}.Rate(AR)
		inputs    = out.inputs
		numInputs = len(inputs)
	)
	if numInputs != 2 {
		t.Fatalf("expected 2 inputs but got %d", numInputs)
	}
}
