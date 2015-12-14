package sc

// func TestSums(t *testing.T) {
// 	defName := "Sum3Test"
// 	def := NewSynthdef(defName, func(params Params) Ugen {
// 		bus := C(0)
// 		sig := Sum3(AR,
// 			PinkNoise{}.Rate(AR).Mul(C(0.1)),
// 			FSinOsc{Freq: C(801), Phase: C(0.1)}.Rate(AR),
// 			LFSaw{Freq: C(40), Iphase: C(0.1)}.Rate(AR),
// 		)
// 		return Out{bus, sig}.Rate(AR)
// 	})
// 	compareAndWrite(t, defName, def)
// }
