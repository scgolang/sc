package sc

// FIXME
// import (
// 	. "github.com/scgolang/sc/types"
// 	. "github.com/scgolang/sc/ugens"
// 	"testing"
// )

// func TestCombL(t *testing.T) {
// 	defName := "CombLTest"
// 	def := NewSynthdef(defName, func(p Params) Ugen {
// 		bus := C(0)
// 		line := XLine{
// 			Start: C(0.0001),
// 			End:   C(0.01),
// 			Dur:   C(20),
// 		}.Rate(KR)
// 		sig := CombL{
// 			In:           WhiteNoise{}.Rate(AR).Mul(C(0.01)),
// 			MaxDelayTime: C(0.01),
// 			DelayTime:    line,
// 			DecayTime:    C(0.2),
// 		}.Rate(AR)
// 		return Out{bus, sig}.Rate(AR)
// 	})
// 	compareAndWrite(t, defName, def)
// }
