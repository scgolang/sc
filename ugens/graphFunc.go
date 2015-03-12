package ugens

import (
	. "github.com/briansorahan/sc/types"
)

// HACK convert Params to an interface type
type UgenGraphFunc func(params *Params) UgenNode

type PlayFunc func() UgenNode

// Play corresponds to http://doc.sccode.org/Classes/Function.html#-play.
// This is syntactic sugar for
//     temp := DefaultServer.TempDefName()
//     def := NewSynthdef(temp, func(params Params) UgenNode {
//     })
//     DefaultServer.SendDef(def)
//     sid := DefaultServer.NextSynthID()
//     DefaultServer.NewSynth(temp, sid, AddToHead, DefaultGroupID)
func Play(node UgenNode) error {
	// To implement this we need a DefaultServer and
	// a way to generate the names of the temp synthdefs.
	// If the ugen node returned by f is not Out,
	// wrap the node in an Out ugen with bus 0.
	return nil
}

// Play(SinOsc.Ar(440))

// Once we support multichannel expansion...
// Play([]UgenNode{SinOsc.Ar(440), SinOsc.Ar(441)})
