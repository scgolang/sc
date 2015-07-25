package sc

import (
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"testing"
)

func TestFreeVerb(t *testing.T) {
	defName := "FreeVerbTest"
	def := NewSynthdef(defName, func(p Params) Ugen {
		// arg mix=0.25, room=0.15, damp=0.5;
		// var decay = Decay.ar(Impulse.ar(1), 0.25, LFCub.ar(1200, 0, 0.1));
		// var sig = FreeVerb.ar(decay, mix, room, damp);
		// Out.ar(0, sig);
		mix := p.Add("mix", 0.25)
		room := p.Add("room", 0.15)
		damp := p.Add("damp", 0.5)
		bus := C(0)
		impulse := Impulse{Freq: C(1)}.Rate(AR)
		lfcub := LFCub{Freq: C(1200), Iphase: C(0)}.Rate(AR).Mul(C(0.1))
		decay := Decay{In: impulse, Decay: C(0.25)}.Rate(AR).Mul(lfcub)
		sig := FreeVerb{In: decay, Mix: mix, Room: room, Damp: damp}.Rate(AR)
		return Out{bus, sig}.Rate(AR)
	})
	compareAndWrite(t, defName, def)
}
