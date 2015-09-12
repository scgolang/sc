package sc

import "testing"

func TestBuffer(t *testing.T) {
	c := NewClient("127.0.0.1:57112")
	buf := newReadBuffer("foo", c)
	n := buf.Num()
	if n != 0 {
		t.Fatalf("expected 0, but got %d", n)
	}
	buf = newReadBuffer("bar", c)
	n = buf.Num()
	if n != 1 {
		t.Fatalf("expected 1, but got %d", n)
	}
	// should return the first buffer
	newBuf := newReadBuffer("foo", c)
	n = newBuf.Num()
	if n != 0 {
		t.Fatalf("expected 0, but got %d", n)
	}
}
