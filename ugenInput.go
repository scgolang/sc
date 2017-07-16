package sc

import (
	"encoding/binary"
	"io"
)

// NewInput creates a ugen suitable for use as an input to other ugens.
// It will return either a single-channel ugen or a multi-channel ugen.
func NewInput(name string, rate int8, specialIndex int16, numOutputs int, inputs ...Input) Input {
	var (
		expanded = expandInputs(inputs...)
		l        = len(expanded)
	)
	if l == 1 {
		return NewUgen(name, rate, specialIndex, numOutputs, inputs...)
	}
	a := make([]Input, l)
	for i := range a {
		a[i] = NewUgen(name, rate, specialIndex, numOutputs, expanded[i]...)
	}
	return Multi(a...)
}

// expandInputs turns an array of inputs into a 2-dimensional array
// of inputs where the 1st dimension is the channel and
// the second is the array of inputs for each channel.
func expandInputs(inputs ...Input) [][]Input {
	// first pass to determine how large each array needs to be
	// this could probably be more efficient
	sz := 0
	for _, in := range inputs {
		if multi, isMulti := in.(MultiInput); isMulti {
			ins := multi.InputArray()
			l := len(ins)
			if l > sz {
				sz = l
			}
		}
	}
	if sz == 0 {
		// none were multi-channel inputs
		return [][]Input{inputs}
	}
	var (
		arr = make([][]Input, sz)
		n   = len(inputs)
	)
	for i := range arr {
		arr[i] = make([]Input, n)

		for j := range arr[i] {
			in := inputs[j]

			if multi, isMulti := in.(MultiInput); isMulti {
				ins := multi.InputArray()
				arr[i][j] = ins[i%len(ins)]
			} else {
				arr[i][j] = in
			}
		}
	}

	return arr
}

// UgenInput defines a ugen input as it appears in the synthdef file format.
type UgenInput struct {
	UgenIndex   int32 `json:"ugenIndex"   xml:"ugenIndex,attr"`
	OutputIndex int32 `json:"outputIndex" xml:"outputIndex,attr"`
}

// IsConstant returns true if a UgenInput has a UgenIndex of -1
func (ui UgenInput) IsConstant() bool {
	return ui.UgenIndex == -1
}

// Write writes an input to an io.Writer
func (ui UgenInput) Write(w io.Writer) error {
	if we := binary.Write(w, byteOrder, ui.UgenIndex); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, ui.OutputIndex)
}

func readUgenInput(r io.Reader) (UgenInput, error) {
	var ui UgenInput
	if err := binary.Read(r, byteOrder, &ui.UgenIndex); err != nil {
		return ui, err
	}
	if err := binary.Read(r, byteOrder, &ui.OutputIndex); err != nil {
		return ui, err
	}
	return ui, nil
}
