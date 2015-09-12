package main

import (
	"fmt"
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
)

func main() {
	const synthName = "DelayCExample"
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
		bus, dust := C(0), Dust{Density: C(1)}.Rate(AR).Mul(C(0.5))
		noise := WhiteNoise{}.Rate(AR)
		decay := Decay{In: dust, Decay: C(0.3)}.Rate(AR).Mul(noise)
		sig := DelayC{
			In:           decay,
			MaxDelayTime: C(0.2),
			DelayTime:    C(0.2),
		}.Rate(AR).Add(decay)
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
