package ugens

import (
	"testing"
)

func TestAr(t *testing.T) {
	node := SinOsc.Ar()
	if node == nil {
		t.Fatalf("Ar returned nil")
	}
	node = SinOsc.Ar(440)
	if node == nil {
		t.Fatalf("Ar returned nil")
	}
	node = SinOsc.Ar(220.1, 0.5)
	if node == nil {
		t.Fatalf("Ar returned nil")
	}
}

func TestKr(t *testing.T) {
	node := SinOsc.Kr()
	if node == nil {
		t.Fatalf("Kr returned nil")
	}
	node = SinOsc.Kr(440)
	if node == nil {
		t.Fatalf("Kr returned nil")
	}
	node = SinOsc.Kr(220.1, 0.5)
	if node == nil {
		t.Fatalf("Kr returned nil")
	}
}

func TestIr(t *testing.T) {
	node := SinOsc.Ir()
	if node == nil {
		t.Fatalf("Ir returned nil")
	}
	node = SinOsc.Ir(440)
	if node == nil {
		t.Fatalf("Ir returned nil")
	}
	node = SinOsc.Ir(220.1, 0.5)
	if node == nil {
		t.Fatalf("Ir returned nil")
	}
}
