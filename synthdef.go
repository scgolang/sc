package sc

import (
	"io"
)

// Synthdef
type Synthdef interface {
	// Name returns the name of the synthdef.
	Name() string

	// AppendConstant appends a float to the synthdef's
	// list of constants only if it isn't there already.
	AppendConstant(value float32)

	// AppendUgen appends a Ugen to the synthdef's
	// list of Ugen only if it isn't there already.
	AppendUgen(value *Ugen)

	// Rep returns a structure that can be serialized
	// to the synthdef representation that scsynth supports.
	Rep() *SynthdefRep

	// Dump writes human-readable information about a synthdef
	// to an io.Writer
	Dump(w io.Writer) error
}
