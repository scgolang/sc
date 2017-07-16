package sc

// NewControl creates a new Control.
func NewControl(numOutputs int) *Ugen {
	outputs := make([]Output, numOutputs)
	o := Output(1)
	for i := 0; i < numOutputs; i++ {
		outputs[i] = o
	}
	return &Ugen{
		Name:         "Control",
		Rate:         KR,
		SpecialIndex: 0,
		Outputs:      outputs,
	}
}
