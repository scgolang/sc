package sc

import (
	"encoding/binary"
	"io"
)

// Input is implemented by any value that can serve as a
// ugen input. This includes synthdef parameters,
// constants, and other ugens.
type Input interface {
	// Add adds one Input to another.
	Add(val Input) Input

	// Max returns the max of one signal and another.
	Max(other Input) Input

	// Midicps converts MIDI note number to cycles per second.
	Midicps() Input

	// Mul multiplies one Input by another.
	Mul(val Input) Input

	// MulAdd multiplies and adds an Input using two others.
	MulAdd(mul, add Input) Input

	// Neg is a convenience operator that multiplies a signal by -1.
	Neg() Input

	// SoftClip distorts a signal with a perfectly linear range from -0.5 to 0.5
	SoftClip() Input
}

func readInput(r io.Reader) (UgenInput, error) {
	var ui UgenInput
	if err := binary.Read(r, byteOrder, &ui.UgenIndex); err != nil {
		return ui, err
	}
	if err := binary.Read(r, byteOrder, &ui.OutputIndex); err != nil {
		return ui, err
	}
	return ui, nil
}
