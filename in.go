package sc

// In reads signals from busses.
type In struct {
	NumChannels int
	Bus         Input
}

func (in *In) defaults() {
	if in.Bus == nil {
		in.Bus = C(0)
	}
	if in.NumChannels == 0 {
		in.NumChannels = 1
	}
}

// Rate returns an Input at the specified rate.
// This method panics if rate is not a supported value.
func (in In) Rate(rate int8) Input {
	CheckRate(rate)
	(&in).defaults()

	ins := make([]Input, in.NumChannels)
	for i := range ins {
		ins[i] = UgenInput("In", rate, 0, 1, in.Bus)
	}
	return Multi(ins...)
}
