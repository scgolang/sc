package main

import (
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"log"
	"time"
)

func main() {
	client, err := NewClient("127.0.0.1", 51670)
	if err != nil {
		log.Fatal(err)
	}
	def := NewSynthdef("Envgen1", func(p Params) Ugen {
		bus := C(0)
		attack, release := C(0.01), C(1)
		level, curveature := C(1), C(-4)
		perc := EnvPerc{attack, release, level, curveature}
		gate, levelScale, levelBias, timeScale := C(1), C(1), C(0), C(1)
		ampEnv := EnvGen{perc, gate, levelScale, levelBias, timeScale, FreeEnclosing}.Rate(KR)
		noise := PinkNoise{}.Rate(AR).Mul(ampEnv)
		return Out{bus, noise}.Rate(AR)
	})
	err = client.SendDef(def)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1000 * time.Millisecond)
	err = client.NewSynth("Envgen1", client.NextSynthID(), AddToHead, DefaultGroupID)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5000 * time.Millisecond)
}
