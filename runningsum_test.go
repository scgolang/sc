package sc

import (
	"testing"
)

func TestRunningSum(t *testing.T) {
	const defName = "RunningSumTest"

	// var x = 100;
	// Out.ar(0, RunningSum.ar(LFSaw.ar, x) * (x.reciprocal));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		x := C(100)
		sum := A(RunningSum{
			In:      A(LFSaw{}),
			NumSamp: x,
		})
		return Out{
			Bus:      C(0),
			Channels: sum.Mul(x.Reciprocal()),
		}.Rate(AR)
	}))
}
