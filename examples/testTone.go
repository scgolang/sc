package main

import (
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"log"
)

func main() {
	// create a client and connect to the server
	client := NewClient("127.0.0.1", 57121)
	err := client.Connect("127.0.0.1", 57120)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.AddDefaultGroup()
	if err != nil {
		log.Fatal(err)
	}
	// create a new synthdef and send it to the server
	def := NewSynthdef("testTone", func(p Params) Ugen {
		bus, env := C(0), EnvGen{Env: EnvPerc{}, Done: FreeEnclosing}.Rate(KR)
		return Out{bus, SinOsc{}.Rate(AR).Mul(env)}.Rate(AR)
	})
	err = client.SendDef(def)
	if err != nil {
		log.Fatal(err)
	}
	// get the next available synth ID
	id := client.NextSynthID()
	// add a synth instance to the tail of the default group
	_, err = client.Synth("testTone", id, AddToTail, DefaultGroupID)
	if err != nil {
		log.Fatal(err)
	}
}
