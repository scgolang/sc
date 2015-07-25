package main

import (
	"fmt"
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
)

func main() {
	const synthName = "FormletExample"
	client := NewClient("127.0.0.1", 57112)
	err := client.Connect("127.0.0.1", 57110)
	if err != nil {
		panic(err)
	}
	defaultGroup, err := client.AddDefaultGroup()
	if err != nil {
		panic(err)
	}
	def := NewSynthdef(synthName, func(p Params) Ugen {
		bus, sine := C(0), SinOsc{Freq: C(5)}.Rate(KR).MulAdd(C(20), C(300))
		blip := Blip{Freq: sine, Harm: C(1000)}.Rate(AR).Mul(C(0.1))
		line := XLine{Start: C(1500), End: C(700), Dur: C(8)}.Rate(KR)
		sig := Formlet{
			In:         blip,
			Freq:       line,
			AttackTime: C(0.005),
			DecayTime:  C(0.4),
		}.Rate(AR)
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
