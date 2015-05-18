package main

import (
	. "github.com/scgolang/sc"
	. "github.com/scgolang/sc/types"
	. "github.com/scgolang/sc/ugens"
	"log"
	"time"
)

func main() {
	options := ServerOptions{
		EchoScsynthStdout: true,
	}
	server, err := NewServer("127.0.0.1", 51670, options)
	if err != nil {
		log.Fatal(err)
	}
	// HACK convert Params to an interface type
	def := NewSynthdef("SineTone", func(params *Params) Ugen {
		return Out{C(0), SinOsc{}.Rate(AR)}.Rate(AR)
	})
	done := server.Run()
	err = server.SendDef(def)
	if err != nil {
		log.Fatal(err)
	}
	sid := server.NextSynthID()
	err = server.NewSynth("SineTone", sid, AddToHead, DefaultGroupID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("added synth %d\n", sid)
	time.Sleep(1000 * time.Millisecond)
	server.Quit()
	err = <-done
	if err != nil {
		log.Fatal(err)
	}
}
