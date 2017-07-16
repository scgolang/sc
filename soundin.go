package sc

// SoundIn is a convenience UGen to read audio from the input of your computer or soundcard.
// It is a wrapper UGen based on In, which offsets the index such that 0 will always
// correspond to the first input regardless of the number of inputs present.
type SoundIn struct {
	Bus Input
}

// Rate returns an Input at the specified rate.
// This method panics if rate is not a supported value.
func (s SoundIn) Rate(rate int8) Input {
	CheckRate(rate)

	var (
		nobs   = NumOutputBuses{}.Rate(IR)
		result Input
	)
	// TODO: simplify this.
	switch x := s.Bus.(type) {
	case C:
		first := NewUgen("In", rate, 0, 1, nobs)

		if x > 0 {
			result = NewUgen("In", rate, 0, 1, nobs.Add(x))
		} else if x == 0 {
			result = first
		} else {
			panic("SoundIn bus indices must not be negative")
		}
	case Inputs:
		var (
			curr   = soundInParams{numOutputs: 1}
			params = []soundInParams{}
			prev   C
		)
		for i, input := range x {
			switch y := input.(type) {
			case C:
				if i == 0 {
					prev = y
					continue
				}
				if prev+1 == y {
					curr.numOutputs++

					if i == len(x)-1 {
						params = append(params, curr)
					}
					continue
				}
				params = append(params, curr)

				if i == len(x)-1 {
					params = append(params, soundInParams{
						startBus:   y,
						numOutputs: 1,
					})
				} else {
					curr = soundInParams{startBus: y, numOutputs: 1}
				}
			default:
				panic("SoundIn busses must be constant or an array of constants")
			}
		}
		arr := make(Inputs, len(params))

		for i, params := range params {
			if params.startBus == 0 {
				arr[i] = NewUgen("In", rate, 0, params.numOutputs, nobs)
			} else {
				arr[i] = NewUgen("In", rate, 0, params.numOutputs, nobs.Add(params.startBus))
			}
		}
		result = arr
	default:
		panic("SoundIn busses must be constant or an array of constants")
	}
	return result
}

type soundInParams struct {
	startBus   C
	numOutputs int
}
