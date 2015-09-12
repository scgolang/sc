package main

import (
	"fmt"
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
)

func main() {
	const synthName = "FreeVerbExample"
	client := NewClient("127.0.0.1:57112")
	err := client.Connect("127.0.0.1:57110")
	if err != nil {
		panic(err)
	}
	defaultGroup, err := client.AddDefaultGroup()
	if err != nil {
		panic(err)
	}
	def := NewSynthdef(synthName, func(p Params) Ugen {
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
	err = client.SendDef(def)
	if err != nil {
		panic(err)
	}
	synthID := client.NextSynthID()
	_, err = defaultGroup.Synth(synthName, synthID, AddToTail, nil)
	fmt.Printf("created synth %d\n", synthID)
}
