package ugens

import (
	"testing"
)

func TestAddConstantInput(t *testing.T) {
	n := newNode("foo", 2, 0)
	n.addInput(3.14)
	if inputs := n.Inputs(); len(inputs) != 1 {
		t.Fatalf("len(inputs) was %d", len(inputs))
	}
}

func TestIsOutput(t *testing.T) {
	n := newNode("foo", 2, 0)
	n.IsOutput()
	outputs := n.Outputs()
	if numOutputs := len(outputs); numOutputs != 1 {
		t.Fatalf("number of outputs was %d", numOutputs)
	}
}

func TestAddUgenInput(t *testing.T) {
	s := SinOsc.Ar(440)
	if _, isBase := s.(*BaseNode); !isBase {
		t.Fatalf("SinOsc.Ar did not return *BaseNode")
	}
	Out.Ar(0, s)
	outputs := s.Outputs()
	if numOutputs := len(outputs); numOutputs != 1 {
		t.Fatalf("number of SinOsc outputs was %d", numOutputs)
	}
}
