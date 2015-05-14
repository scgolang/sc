package sc

import "testing"

func TestBuffer(t *testing.T) {
	buf := newBuffer("foo")
	n := buf.Num()
	if n != 0 {
		t.Fatalf("expected 0, but got %d", n)
	}
	buf = newBuffer("bar")
	n = buf.Num()
	if n != 1 {
		t.Fatalf("expected 1, but got %d", n)
	}
	// should return the first buffer
	newBuf := newBuffer("foo")
	n = newBuf.Num()
	if n != 0 {
		t.Fatalf("expected 0, but got %d", n)
	}
}
