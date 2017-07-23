package sc

import (
	"math/rand"
)

// Klang is a bank of fixed frequency sine oscillators.
// Klang is more efficient than creating individual oscillators but offers less flexibility.
type Klang struct {
	// Spec consists of 3 Input slices.
	//   1. frequencies: An Array of oscillator frequencies.
	//   2. amplitudes: A slice of oscillator amplitudes, or nil. If nil, then amplitudes default to 1.0.
	//   3. phases: A slice of initial phases, or nil. If nil, then phases default to 0.0.
	// The parameters in specificationsArrayRef can't be changed after it has been started.
	// For a modulatable but less efficient version, see DynKlang.
	Spec Input

	// A scale factor multiplied by all frequencies at initialization time.
	FreqScale Input

	// An offset added to all frequencies at initialization time.
	FreqOffset Input
}

func (k *Klang) defaults() {
	if k.FreqScale == nil {
		k.FreqScale = C(1)
	}
	if k.FreqOffset == nil {
		k.FreqOffset = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (k Klang) Rate(rate int8) Input {
	CheckRate(rate)
	(&k).defaults()

	specs := getArraySpecInputs(k.Spec)

	if len(specs) == 1 {
		ins := append(specs[0].inputs(false), k.FreqScale, k.FreqOffset)
		return NewInput("Klang", rate, 0, 1, ins...)
	}
	var klangs []Input
	for _, spec := range specs {
		ins := append(spec.inputs(false), k.FreqScale, k.FreqOffset)
		klangs = append(klangs, NewInput("Klang", rate, 0, 1, ins...))
	}
	return Multi(klangs...)
}

// Fill returns a slice of inputs that has length n and every element is set to input.
func Fill(n int, input Input) []Input {
	a := make([]Input, n)
	for i := range a {
		a[i] = input
	}
	return a
}

// RandC returns a random constant between min and max.
func RandC(min, max float64) Input {
	return C((rand.Float64() * (max - min)) + min)
}

// RandArray returns a slice (of length n) of random constants between min and max.
func RandArray(n int, min, max float64) []Input {
	ins := make([]Input, n)
	for i := range ins {
		ins[i] = C((rand.Float64() * (max - min)) + min)
	}
	return ins
}
