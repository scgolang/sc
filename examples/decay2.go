package main

import (
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
)

func main() {
	const synthName = "Decay2Example"
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
		bus := C(0)
		line := XLine{Start: C(1), End: C(50), Dur: C(20)}.Rate(KR)
		pulse := Impulse{Freq: line, Phase: C(0.25)}.Rate(AR)
		sig := Decay2{In: pulse, Attack: C(0.01), Decay: C(0.2)}.Rate(AR)
		gain := FSinOsc{Freq: C(600)}.Rate(AR)
		return Out{bus, sig.Mul(gain)}.Rate(AR)
	})
	err = client.SendDef(def)
	if err != nil {
		panic(err)
	}
	synthID := client.NextSynthID()
	_, err = defaultGroup.Synth(synthName, synthID, AddToTail, nil)
}
