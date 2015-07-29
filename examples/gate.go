package main

import (
	"fmt"
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
)

func main() {
	const synthName = "GateExample"
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
		bus, noise := C(0), WhiteNoise{}.Rate(KR)
		pulse := LFPulse{Freq: C(1.333), Iphase: C(0.5)}.Rate(KR)
		sig := Gate{In: noise, Trig: pulse}.Rate(AR)
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
