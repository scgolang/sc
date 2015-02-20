package ugens

import (
	"testing"
)

func TestAddConstantInput(t *testing.T) {
	n := newNode("foo", 2)
	n.addConstantInput(3.14)
	if inputs := n.Inputs(); len(inputs) != 1 {
		t.Fatalf("len(inputs) was %d", len(inputs))
	}
}
