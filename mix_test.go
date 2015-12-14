package sc

// func TestMix(t *testing.T) {
// 	defName := "MixTest"
// 	def := NewSynthdef(defName, func(params Params) Ugen {
// 		bus := C(0)
// 		sig := Mix(AR, []Input{
// 			PinkNoise{}.Rate(AR).Mul(C(0.1)),
// 			FSinOsc{Freq: C(801), Phase: C(0.1)}.Rate(AR),
// 			LFSaw{Freq: C(40), Iphase: C(0.1)}.Rate(AR),
// 			Pulse{Freq: C(436)}.Rate(AR),
// 			Dust{Density: C(4.0)}.Rate(AR),
// 		})
// 		return Out{bus, sig}.Rate(AR)
// 	})
// 	compareAndWrite(t, defName, def)
// }
