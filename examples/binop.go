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
	def := NewSynthdef("Envgen1", func(params *Params) UgenNode {
		bus := C(0)
		attack, release := C(0.01), C(1)
		level, curveature := C(1), C(-4)
		perc := EnvPerc{attack, release, level, curveature}
		gate, levelScale, levelBias, timeScale := C(1), C(1), C(0), C(1)
		ampEnv := EnvGen{perc, gate, levelScale, levelBias, timeScale, FreeEnclosing}.Rate(KR)
		noise := PinkNoise{}.Rate(AR).Mul(ampEnv)
		return Out{bus, noise}.Rate(AR)
	})
	done := server.Run()
	err = server.SendDef(def)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1000 * time.Millisecond)
	err = server.NewSynth("Envgen1", server.NextSynthID(), AddToHead, DefaultGroupID)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5000 * time.Millisecond)
	server.Quit()
	err = <-done
	if err != nil {
		log.Fatal(err)
	}
}
