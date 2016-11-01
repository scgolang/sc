package sc

import (
	"testing"
	"time"
)

func TestBuffer(t *testing.T) {
	c, err := NewClient("udp", "127.0.0.1:57112", "127.0.0.1:57120", 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = c.Close() }() // Best effort.

	buf := newReadBuffer("foo", 0, c)
	if buf.Num != 0 {
		t.Fatalf("expected 0, but got %d", buf.Num)
	}
	buf = newReadBuffer("bar", 1, c)
	if buf.Num != 1 {
		t.Fatalf("expected 1, but got %d", buf.Num)
	}
	// should return the first buffer
	newBuf := newReadBuffer("foo", 0, c)
	if newBuf.Num != 0 {
		t.Fatalf("expected 0, but got %d", newBuf.Num)
	}
}
