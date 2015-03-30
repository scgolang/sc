package sc

import (
	"testing"
)

func TestCstackPushPop(t *testing.T) {
	s := newCstack()
	if s == nil {
		t.Fatalf("newCstack returned nil")
	}
	s.Push(float32(3.14))
	s.Push(float32(-1))
	s.Push(float32(3.14))
	val, more := s.Pop()
	if !more {
		t.Fatalf("expected more values on the stack")
	}
	if val != float32(-1) {
		t.Fatalf("Expected val to be 3.14")
	}
	val, more = s.Pop()
	if !more {
		t.Fatalf("expected more values on the stack")
	}
	if val != float32(3.14) {
		t.Fatalf("Expected val to be 3.14")
	}
	_, more = s.Pop()
	if more {
		t.Fatalf("popping an empty stack should return false for more")
	}
}
