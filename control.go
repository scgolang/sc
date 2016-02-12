package sc

// Control represents a synth control.
// See http://doc.sccode.org/Classes/Control.html.
type Control struct {
	inputs  []Input
	outputs []Output
}

// Name returns the name of the Control.
func (ctl *Control) Name() string {
	return "Control"
}

// Rate returns the rate of the Control.
func (ctl *Control) Rate() int8 {
	return int8(1)
}

// SpecialIndex returns the special index of the Control.
func (ctl *Control) SpecialIndex() int16 {
	return 0
}

// Inputs returns the inputs for the Control.
func (ctl *Control) Inputs() []Input {
	return ctl.inputs
}

// Outputs returns the outputs for the Control.
func (ctl *Control) Outputs() []Output {
	return ctl.outputs
}

// Add adds another input to the Control.
func (ctl *Control) Add(val Input) Input {
	return ctl
}

// Mul multiplies another input by the Control.
func (ctl *Control) Mul(val Input) Input {
	return ctl
}

// MulAdd multiplies and adds inputs to the Control.
func (ctl *Control) MulAdd(mul, add Input) Input {
	return ctl
}

// NewControl creates a new Control.
func NewControl(numOutputs int) Ugen {
	outputs := make([]Output, numOutputs)
	o := Output(1)
	for i := 0; i < numOutputs; i++ {
		outputs[i] = o
	}
	c := Control{make([]Input, 0), outputs}
	return &c
}
