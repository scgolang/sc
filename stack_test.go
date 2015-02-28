package sc

import (
	"testing"
)

func TestStackPushPop(t *testing.T) {
	s := newStack()
	if s == nil {
		t.Fatalf("NewStack returned nil")
	}
	s.Push(float32(3.14))
	s.Push(float32(-1))
	val := s.Pop()
	if fv, isFloat32 := val.(float32); !isFloat32 {
		t.Fatalf("Could not cast val to float32")
		if fv != float32(-1) {
			t.Fatalf("Expected val to be 3.14")
		}
	}
	val = s.Pop()
	if fv, isFloat32 := val.(float32); !isFloat32 {
		t.Fatalf("Could not cast val to float32")
		if fv != float32(3.14) {
			t.Fatalf("Expected val to be 3.14")
		}
	}
	val = s.Pop()
	if val != nil {
		t.Fatalf("popping an empty stack should return nil")
	}
}
