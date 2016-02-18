package sc

// FIXME
// import (
// 	"testing"
// )

// func TestCombC(t *testing.T) {
// 	defName := "CombCTest"
// 	def := NewSynthdef(defName, func(p Params) Ugen {
// 		bus := C(0)
// 		line := XLine{
// 			Start: C(0.0001),
// 			End:   C(0.01),
// 			Dur:   C(20),
// 		}.Rate(KR)
// 		sig := CombC{
// 			In:           WhiteNoise{}.Rate(AR).Mul(C(0.01)),
// 			MaxDelayTime: C(0.01),
// 			DelayTime:    line,
// 			DecayTime:    C(0.2),
// 		}.Rate(AR)
// 		return Out{bus, sig}.Rate(AR)
// 	})
// 	compareAndWrite(t, defName, def)
// }

// func TestCombL(t *testing.T) {
// 	defName := "CombLTest"
// 	def := NewSynthdef(defName, func(p Params) Ugen {
// 		bus := C(0)
// 		line := XLine{
// 			Start: C(0.0001),
// 			End:   C(0.01),
// 			Dur:   C(20),
// 		}.Rate(KR)
// 		sig := Comb{
// 			Interpolation: InterpolationLinear,
// 			In:            WhiteNoise{}.Rate(AR).Mul(C(0.01)),
// 			MaxDelayTime:  C(0.01),
// 			DelayTime:     line,
// 			DecayTime:     C(0.2),
// 		}.Rate(AR)

// 		return Out{Bus: bus, Channels: sig}.Rate(AR)
// 	})
// 	compareAndWrite(t, defName, def)
// }

// func TestCombN(t *testing.T) {
// 	defName := "CombNTest"
// 	def := NewSynthdef(defName, func(p Params) Ugen {
// 		bus := C(0)
// 		line := XLine{
// 			Start: C(0.0001),
// 			End:   C(0.01),
// 			Dur:   C(20),
// 		}.Rate(KR)
// 		sig := CombN{
// 			In:           WhiteNoise{}.Rate(AR).Mul(C(0.01)),
// 			MaxDelayTime: C(0.01),
// 			DelayTime:    line,
// 			DecayTime:    C(0.2),
// 		}.Rate(AR)
// 		return Out{bus, sig}.Rate(AR)
// 	})
// 	compareAndWrite(t, defName, def)
// }
