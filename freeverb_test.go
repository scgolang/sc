package sc

import "testing"

func TestFreeVerb(t *testing.T) {
	name := "FreeVerbExample"
	def := NewSynthdef(name, func(params Params) Ugen {
		bus, In := C(0), SinOsc{Freq: C(220)}.Rate(AR)
		wrp := FreeVerb{In: In}.Rate(AR)
		return Out{bus, wrp}.Rate(AR)
	})
	same, err := def.CompareToFile("testdata/FreeVerbExample.scsyndef")
	if err != nil {
		t.Fatal(err)
	}
	if !same {
		t.Fatalf("synthdef different from sclang version")
	}
}

// FIXME: Order of constants is different

// func TestFreeVerb(t *testing.T) {
// 	defName := "FreeVerbTest"
// 	def := NewSynthdef(defName, func(p Params) Ugen {
// 		mix := p.Add("mix", 0.25)
// 		room := p.Add("room", 0.15)
// 		damp := p.Add("damp", 0.5)
// 		bus := C(0)
// 		impulse := Impulse{Freq: C(1)}.Rate(AR)
// 		lfcub := LFCub{Freq: C(1200), Iphase: C(0)}.Rate(AR).Mul(C(0.1))
// 		decay := Decay{In: impulse, Decay: C(0.25)}.Rate(AR).Mul(lfcub)
// 		sig := FreeVerb{In: decay, Mix: mix, Room: room, Damp: damp}.Rate(AR)
// 		return Out{bus, sig}.Rate(AR)
// 	})
// 	compareAndWrite(t, defName, def)
// }
