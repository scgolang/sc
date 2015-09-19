package sc

// import "testing"

// FIXME
// func TestAllpassnExample(t *testing.T) {
// 	def := NewSynthdef("AllpassnExample", func(p Params) Ugen {
// 		noise := WhiteNoise{}.Rate(AR)
// 		dust := Dust{C(1)}.Rate(AR).Mul(C(0.5))
// 		decay := Decay{dust, C(0.2)}.Rate(AR).Mul(noise)
// 		sig := AllpassN{decay, C(0.2), C(0.2), C(3)}.Rate(AR)
// 		return Out{C(0), sig}.Rate(AR)
// 	})
// 	same, err := def.Compare(`{
//         var noise = WhiteNoise.ar();
//         // var dust = Dust.ar(1, 0.5);
//         var dust = Dust.ar(1) * 0.5;
//         // var decay = Decay.ar(dust, 0.2, noise);
//         var decay = Decay.ar(dust, 0.2) * noise;
//         var sig = AllpassN.ar(decay, 0.2, 0.2, 3);
//         Out.ar(0, sig);
//     }`)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !same {
// 		t.Fatalf("synthdef different from sclang-generated version")
// 	}
// }
