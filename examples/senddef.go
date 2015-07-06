package main

import (
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"log"
	"time"
)

func main() {
	client, err := NewClient("127.0.0.1", ScsynthDefaultPort)
	if err != nil {
		log.Fatal(err)
	}
	def := NewSynthdef("SineTone", func(p Params) Ugen {
		bus, chaos := C(0), Line{C(1.0), C(2.0), C(10), DoNothing}.Rate(KR)
		sig := Crackle{chaos}.Rate(AR).MulAdd(C(0.5), C(0.5))
		return Out{bus, sig}.Rate(AR)
	})
	err = client.SendDef(def)
	if err != nil {
		log.Fatal(err)
	}
	sid := client.NextSynthID()
	err = client.NewSynth("SineTone", sid, AddToHead, DefaultGroupID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added synth %d\n", sid)
	time.Sleep(1000 * time.Millisecond)
}
