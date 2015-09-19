package sc

// FIXME
// import (
// 	"testing"
// )

// func TestDecay2(t *testing.T) {
// 	defName := "Decay2Test"
// 	def := NewSynthdef(defName, func(p Params) Ugen {
// 		bus := C(0)
// 		line := XLine{Start: C(1), End: C(50), Dur: C(20)}.Rate(KR)
// 		pulse := Impulse{Freq: line, Phase: C(0.25)}.Rate(AR)
// 		sig := Decay2{In: pulse, Attack: C(0.01), Decay: C(0.2)}.Rate(AR)
// 		gain := FSinOsc{Freq: C(600)}.Rate(AR)
// 		return Out{bus, sig.Mul(gain)}.Rate(AR)
// 	})
// 	compareAndWrite(t, defName, def)
// }
