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

	var (
		uin = NewInput("In", rate, 0, in.NumChannels, in.Bus)
		ins = make([]Input, in.NumChannels)
	)
	for i := range ins {
		ins[i] = uin
	}
	return Multi(ins...)
}

func defIn(params Params) Ugen {
	var (
		in  = params.Add("in", 0)
		out = params.Add("out", 0)
	)
	return Out{
		Bus: out,
		Channels: In{
			NumChannels: 1,
			Bus:         in,
		}.Rate(AR),
	}.Rate(AR)
}
