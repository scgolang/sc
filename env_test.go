package sc

import (
	"math"
	"testing"
)

// All of the expected values in these tests were generated with sclang using SCIDE.

func TestPerc(t *testing.T) {
	var (
		e      = EnvPerc{}
		expect = []float32{0, 2, -99, -99, 1, 0.01, 5, -4, 0, 1, 5, -4}
	)
	verifyInputs(t, expect, e.Inputs())
}

func TestLinen(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		var (
			e      = EnvLinen{}
			expect = []float32{0, 3, -99, -99, 1, 0.01, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0}
		)
		verifyInputs(t, expect, e.Inputs())
	})

	t.Run("welch", func(t *testing.T) {
		var (
			f           = EnvLinen{Curve: "welch"}
			expectWelch = []float32{0, 3, -99, -99, 1, 0.01, 4, 0, 1, 1, 4, 0, 0, 1, 4, 0}
		)
		verifyInputs(t, expectWelch, f.Inputs())
	})
}

func TestTriangle(t *testing.T) {
	var (
		e      = EnvTriangle{}
		expect = []float32{0, 2, -99, -99, 1, 0.5, 1, 0, 0, 0.5, 1, 0}
	)
	verifyInputs(t, expect, e.Inputs())
}

func TestSine(t *testing.T) {
	var (
		e      = EnvSine{}
		expect = []float32{0, 2, -99, -99, 1, 0.5, 3, 0, 0, 0.5, 3, 0}
	)
	verifyInputs(t, expect, e.Inputs())
}

func TestPairs(t *testing.T) {
	var (
		e = EnvPairs{
			Pairs: Pairs([][2]float32{
				{0, 1},
				{2.1, 0.5},
				{3, 1.4},
			}),
			Curve: "exponential",
		}
		expect = []float32{1, 2, -99, -99, 0.5, 2.1, 2, 0, 1.4, 0.9, 2, 0}
	)
	verifyInputs(t, expect, e.Inputs())
}

func TestTLC(t *testing.T) {
	var (
		e = EnvTLC([]TLC{
			{0, 1, "sine"},
			{2.1, 0.5, "lin"},
			{3, 1.4, "lin"},
		})
		expect = []float32{1, 2, -99, -99, 0.5, 2.1, 3, 0, 1.4, 0.9, 1, 0}
	)
	verifyInputs(t, expect, e.Inputs())
}

func TestTHX(t *testing.T) {
	var (
		// I wrote this test after encountering problems with Env
		// while implementing the THX deep note.
		env = Env{
			Levels: []Input{
				C(0),
				C(0.1),
				C(1),
			},
			Times: []Input{
				C(5),
				C(8),
			},
			Curve: []Input{
				CurveExp,
				CurveCustom,
			},
		}
		expect = []float32{0, 2, -99, -99, 0.1, 5, 5, 2, 1, 8, 5, 5}
	)
	verifyInputs(t, expect, env.Inputs())
}

const epsilon = 1e-6

func verifyInputs(t *testing.T, expect []float32, inputs []Input) {
	if len(expect) != len(inputs) {
		t.Fatalf("expected: %+v\n         got:      %+v", expect, inputs)
	}
	for i, in := range inputs {
		var val float32
		switch v := in.(type) {
		case C:
			val = float32(v)
		default:
			t.Fatalf("input %d was not a float or int (%v)", i, in)
		}
		if delta := float64(val - expect[i]); math.Abs(delta) > epsilon {
			t.Fatalf("expected: %+v\n         got:      %+v", expect, inputs)
		}
	}
}
