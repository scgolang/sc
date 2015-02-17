package ugens

import (
	"testing"
)

func TestOutNoArguments(t *testing.T) {
	defer func() {
		if val := recover(); val == nil {
			t.Fatalf("expected panic when Out.Ar is called with no arguments")
		}
	}()
	Out.Ar()
}

func TestOutOneArgument(t *testing.T) {
	defer func() {
		if val := recover(); val == nil {
			t.Fatalf("expected panic when Out.Ar is called with no arguments")
		}
	}()
	Out.Ar(0)
}

func TestOut(t *testing.T) {
	out := Out.Ar(0, SinOsc.Ar())
	if out == nil {
		t.Fatalf("out was nil")
	}
}
