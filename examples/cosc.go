package main

import (
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
)

func main() {
	const synthName = "COscExample"
	client := NewClient("127.0.0.1", 57120)
	err := client.Connect("127.0.0.1", 57110)
	if err != nil {
		panic(err)
	}
	defaultGroup, err := client.AddDefaultGroup()
	if err != nil {
		panic(err)
	}
	buf, err := client.AllocBuffer(512, 1)
	if err != nil {
		panic(err)
	}
	bufRoutine := BufferRoutineSine1
	bufFlags := BufferFlagNormalize | BufferFlagWavetable | BufferFlagClear
	partials := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, p := range partials {
		partials[i] = 1 / p
	}
	err = buf.Gen(bufRoutine, bufFlags, partials...)
	if err != nil {
		panic(err)
	}
	def := NewSynthdef(synthName, func(p Params) Ugen {
		bus, gain := C(0), C(0.25)
		freq, beats := C(200), C(0.7)
		sig := COsc{
			BufNum: C(float32(buf.Num())),
			Freq:   freq,
			Beats:  beats,
		}.Rate(AR)
		return Out{bus, sig.Mul(gain)}.Rate(AR)
	})
	err = client.SendDef(def)
	if err != nil {
		panic(err)
	}
	synthID := client.NextSynthID()
	_, err = defaultGroup.Synth(synthName, synthID, AddToTail, nil)
	if err != nil {
		panic(err)
	}
}
