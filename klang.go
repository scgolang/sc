package sc

import (
	"math/rand"
	"time"
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
	Spec [3][]Input

	// A scale factor multiplied by all frequencies at initialization time.
	FreqScale Input

	// An offset added to all frequencies at initialization time.
	FreqOffset Input
}

func (k *Klang) defaults() {
	if k.Spec[1] == nil {
		k.Spec[1] = Fill(len(k.Spec[0]), C(1))
	}
	if k.Spec[2] == nil {
		k.Spec[2] = Fill(len(k.Spec[0]), C(0))
	}
	if k.FreqScale == nil {
		k.FreqScale = C(1)
	}
	if k.FreqOffset == nil {
		k.FreqOffset = C(0)
	}
}

func (k Klang) inputs() []Input {
	var ins []Input

	for i, freq := range k.Spec[0] {
		if i >= len(k.Spec[1]) {
			ins = append(ins, C(1))
		} else {
			ins = append(ins, k.Spec[1][i])
		}
		if i >= len(k.Spec[2]) {
			ins = append(ins, C(0))
		} else {
			ins = append(ins, k.Spec[2][i])
		}
		ins = append(ins, freq)
	}
	return append(ins, k.FreqScale, k.FreqOffset)
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (k Klang) Rate(rate int8) Input {
	CheckRate(rate)
	(&k).defaults()
	return NewInput("Klang", rate, 0, 1, k.inputs()...)
}

func Fill(n int, input Input) []Input {
	a := make([]Input, n)
	for i := range a {
		a[i] = input
	}
	return a
}

func RandC(min, max float64) Input {
	return C((rand.Float64() * (max - min)) + min)
}

func RandArray(n int, min, max float64) []Input {
	ins := make([]Input, n)
	for i := range ins {
		ins[i] = C((rand.Float64() * (max - min)) + min)
	}
	return ins
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
