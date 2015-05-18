package sc

// import . "github.com/scgolang/sc/types"
// import . "github.com/scgolang/sc/ugens"
// import "os"
// import "testing"

// FIXME or document why this is broken
// func TestAllpassnExample(t *testing.T) {
// 	def := NewSynthdef("AllpassnExample", func(p *Params) Ugen {
// 		noise := WhiteNoise{}.Rate(AR)
// 		dust := Dust{C(1)}.Rate(AR).Mul(C(0.5))
// 		decay := Decay{dust, C(0.2)}.Rate(AR).Mul(noise)
// 		sig := AllpassN{decay, C(0.2), C(0.2), C(3)}.Rate(AR)
// 		return Out{C(0), sig}.Rate(AR)
// 	})
// 	f, err := os.Create("AllpassnExample.gosyndef")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	err = def.Write(f)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	same, err := def.CompareToFile("AllpassnExample.scsyndef")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !same {
// 		t.Fatalf("synthdef different from sclang-generated version")
// 	}
// }
