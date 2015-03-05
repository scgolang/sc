package main

import (
	. "github.com/briansorahan/sc"
	. "github.com/briansorahan/sc/types"
	. "github.com/briansorahan/sc/ugens"
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
	def := NewSynthdef("SineTone", func(params Params) UgenNode {
		return Out.Ar(0, SinOsc.Ar(440))
	})
	done := server.Run()
	err = server.DumpOSC(DumpAll)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(500 * time.Millisecond)
	err = server.SendDef(def)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5000 * time.Millisecond)
	err = server.NewSynth("SineTone", server.NextSynthID(), AddToHead, DefaultGroupID)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1000 * time.Millisecond)
	server.Quit()
	err = <-done
	if err != nil {
		log.Fatal(err)
	}
}
