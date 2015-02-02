package main

import (
	"github.com/briansorahan/gosc"
	. "github.com/briansorahan/gosc/synthdefs"
)

func main() {
	gosc.NewSynthDef("SineTone", function() {
		return Out.Ar(0, SinOsc.Ar(440))
	}).writeDefFile(os.Getcwd())
}
