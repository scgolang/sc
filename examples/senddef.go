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
	def := NewSynthdef("SineTone", func(p Params) Ugen {
		return Out{C(0), SinOsc{}.Rate(AR)}.Rate(AR)
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
