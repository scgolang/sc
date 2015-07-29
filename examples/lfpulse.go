package main

import (
	"fmt"
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
)

func main() {
	const synthName = "LFPulseExample"
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
		lfoFreq, lfoPhase, lfoWidth := C(3), C(0), C(0.3)
		bus, gain := C(0), C(0.1)
		freq := LFPulse{lfoFreq, lfoPhase, lfoWidth}.Rate(KR).MulAdd(C(200), C(200))
		iphase, width := C(0), C(0.2)
		sig := LFPulse{freq, iphase, width}.Rate(AR).Mul(gain)
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
