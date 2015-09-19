package sc

import (
	"testing"
)

func TestAddConstantInput(t *testing.T) {
	n := NewUgenNode("foo", 2, 0, 1, C(3.14))
	if inputs := n.Inputs(); len(inputs) != 1 {
		t.Fatalf("len(inputs) was %d", len(inputs))
	}
}

func TestIsOutput(t *testing.T) {
	n := NewUgenNode("foo", 2, 0, 1)
	n.IsOutput()
	outputs := n.Outputs()
	if numOutputs := len(outputs); numOutputs != 1 {
		t.Fatalf("number of outputs was %d", numOutputs)
	}
}

func TestAddUgenInput(t *testing.T) {
	s := SinOsc{}.Rate(AR)
	if s == nil {
		t.Fatalf("SinOsc.Rate returned nil")
	}
	Out{C(0), s}.Rate(AR)
	if sn, isNode := s.(*UgenNode); isNode {
		outputs := sn.Outputs()
		if numOutputs := len(outputs); numOutputs != 1 {
			t.Fatalf("number of SinOsc outputs was %d", numOutputs)
		}
	} else {
		t.Fatalf("SinOsc with no multi inputs should return *Node")
	}
}
