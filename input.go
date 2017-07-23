package sc

import (
	"encoding/binary"
	"io"
)

// Input is implemented by any value that can serve as a
// ugen input. This includes synthdef parameters,
// constants, and other ugens.
type Input interface {
	// Abs computes the absolute value of a signal.
	Abs() Input

	// Add adds one Input to another.
	Add(val Input) Input

	// AmpDb converts linear amplitude to decibels.
	AmpDb() Input

	// Ceil computes the ceiling (next lowest integer) of a signal.
	Ceil() Input

	// Cpsmidi converts frequency in Hz to midi note values.
	Cpsmidi() Input

	// Cubed raises a signal to the power of 3.
	Cubed() Input

	// DbAmp converts decibels to linear amplitude.
	DbAmp() Input

	// Exp computes the exponential of a signal.
	Exp() Input

	// Floor computes the floor (next lowest integer) of a signal.
	Floor() Input

	// Frac returns the fractional part of a signal.
	Frac() Input

	// Max returns the max of one signal and another.
	Max(other Input) Input

	// Midicps converts MIDI note number to cycles per second.
	Midicps() Input

	// Midiratio converts an interval in MIDI notes into a frequency ratio.
	Midiratio() Input

	// Mul multiplies one Input by another.
	Mul(val Input) Input

	// MulAdd multiplies and adds an Input using two others.
	MulAdd(mul, add Input) Input

	// Neg is a convenience operator that multiplies a signal by -1.
	Neg() Input

	// Ratiomidi converts a frequency ratio to an interval in MIDI notes.
	Ratiomidi() Input

	// Reciprocal computes the reciprocal of a signal.
	Reciprocal() Input

	// Sign computes the sign of a signal.
	// This returns -1 when a < 0, +1 when a > 0, 0 when a is 0.
	Sign() Input

	// SoftClip distorts a signal with a perfectly linear range from -0.5 to 0.5
	SoftClip() Input

	// Sqrt computes the square root of a signal.
	// The definition of square root is extended for signals so that sqrt(a)
	// when a < 0 returns -sqrt(-a).
	Sqrt() Input

	// Squared raises a signal to the power of 2.
	Squared() Input
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
