package main

import (
	"github.com/briansorahan/gosc"
	. "github.com/briansorahan/gosc/synthdefs"
)

func main() {
	gosc.NewSynthDef("SineTone", function(ugen Ugen) {
		return ugen.Ar("Out", 0, ugen.Ar("SinOsc", 440))
	}).writeDefFile(os.Getcwd())
}
