package sc

// import (
// 	"testing"
// )

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
