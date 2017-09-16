package sc

// import (
// 	"testing"
// )

// func TestBLowPass4(t *testing.T) {
// 	defName := "BLowPass4Test"

// 	// Out.ar(0, BLowPass4.ar(Blip.ar(400, 4), 300, 0.5));
// 	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
// 		return Out{
// 			Bus: C(0),
// 			Channels: BLowPass4{
// 				In: A(Blip{
// 					Freq: C(400),
// 					Harm: C(4),
// 				}),
// 				Freq: C(300),
// 				RQ:   C(0.5),
// 			}.Rate(AR),
// 		}.Rate(AR)
// 	}))
// }
