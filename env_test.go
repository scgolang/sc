package sc

import (
	"math"
	"testing"
)

const (
	ACCEPTABLE_FLOAT_ERROR = 0.0001
)

// All of the expected values in these tests were generated
// with sclang using SCIDE.

func TestPerc(t *testing.T) {
	e := EnvPerc{}
	expect := []float32{0, 2, -99, -99, 1, 0.01, 5, -4, 0, 1, 5, -4}
	verifyInputs(t, expect, e.Inputs())
}

func TestLinen(t *testing.T) {
	e := EnvLinen{}
	expect := []float32{0, 3, -99, -99, 1, 0.01, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0}
	verifyInputs(t, expect, e.Inputs())
	f := EnvLinen{CurveType: CurveWelch}
	expectWelch := []float32{0, 3, -99, -99, 1, 0.01, 4, 0, 1, 1, 4, 0, 0, 1, 4, 0}
	verifyInputs(t, expectWelch, f.Inputs())
}

func TestTriangle(t *testing.T) {
	e := EnvTriangle{}
	expect := []float32{0, 2, -99, -99, 1, 0.5, 1, 0, 0, 0.5, 1, 0}
	verifyInputs(t, expect, e.Inputs())
}

func TestSine(t *testing.T) {
	e := EnvSine{}
	expect := []float32{0, 2, -99, -99, 1, 0.5, 3, 0, 0, 0.5, 3, 0}
	verifyInputs(t, expect, e.Inputs())
}

func TestPairs(t *testing.T) {
	e := EnvPairs{
		Pairs([][2]float32{
			[2]float32{0, 1},
			[2]float32{2.1, 0.5},
			[2]float32{3, 1.4},
		}),
		CurveExp,
	}
	expect := []float32{1, 2, -99, -99, 0.5, 2.1, 2, 0, 1.4, 0.9, 2, 0}
	verifyInputs(t, expect, e.Inputs())
}

func TestTLC(t *testing.T) {
	e := EnvTLC([]TLC{
		TLC{0, 1, CurveSine},
		TLC{2.1, 0.5, CurveLinear},
		TLC{3, 1.4, CurveLinear},
	})
	// 1 2 -99 -99 0.5 2.1 1 0 1.4 0.9000001 1 0
	expect := []float32{1, 2, -99, -99, 0.5, 2.1, 3, 0, 1.4, 0.9, 1, 0}
	verifyInputs(t, expect, e.Inputs())
}

func verifyInputs(t *testing.T, expect []float32, inputs []Input) {
	t.Logf("%v\n", inputs)
	errmsg := "expected %f for input %d but got %f"
	for i, in := range inputs {
		var val float32
		switch v := in.(type) {
		case C:
			val = float32(v)
		default:
			t.Fatalf("input %d was not a float or int (%v)", i, in)
		}
		if err := math.Abs(float64(val - expect[i])); err > ACCEPTABLE_FLOAT_ERROR {
			t.Fatalf(errmsg, expect[i], i, val)
		}
	}
}
