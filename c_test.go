package sc_test

import (
	"testing"

	"github.com/scgolang/sc"
)

func TestTruncf(t *testing.T) {
	for i, c := range []struct {
		in  [2]float32
		out float32
	}{
		{
			in:  [2]float32{4, 2},
			out: 4,
		},
	} {
		if expected, got := c.out, sc.Truncf(c.in[0], c.in[1]); expected != got {
			t.Fatalf("[%d] expected %f, got %f", i, expected, got)
		}
	}
}
