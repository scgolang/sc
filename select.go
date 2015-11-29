package sc

// Select selects an output from an array of inputs.
type Select struct {
	Which  Input
	Inputs []Input
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (sel Select) Rate(rate int8) Input {
	CheckRate(rate)
	inputs := append([]Input{sel.Which}, sel.Inputs...)
	return UgenInput("Select", rate, 0, 1, inputs...)
}
