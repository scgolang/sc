package sc

// JPverb is a reverb effect from sc3-plugins.
// See https://github.com/supercollider/sc3-plugins
// and http://doc.sccode.org/Classes/JPverb.html
type JPverb struct {
	// In is the input signal.
	In Input

	// T60 is time for the reverb to decay 60db.
	// Does not effect early reflections. (0.1..60)
	T60 Input

	// Damp controls damping of high-frequencies as the reverb decays.
	// 0 is no damping, 1 is very strong damping (0..1)
	Damp Input

	// Size scales size of delay-lines within the reverberator,
	// producing the impression of a larger or smaller space.
	// Values below 1 can sound metallic. (0.5..5)
	Size Input

	// EarlyDiff controls shape of early reflections.
	// Values of 0.707 or more produce smooth exponential decay.
	// Lower values produce a slower build-up of echoes. (0..1)
	EarlyDiff Input

	// ModDepth is the depth of delay-line modulation.
	// Use in combination with ModFreq to set amount of chorusing
	// within the structure. (0..1)
	ModDepth Input

	// ModFreq is the frequency of delay-line modulation.
	// Use in combination with modDepth to set amount of chorusing
	// within the structure. (0..10)
	ModFreq Input

	// Low is the multiplier for the reverberation time within the low band. (0..1)
	Low Input

	// Mid is the multiplier for the reverberation time within the mid band. (0..1)
	Mid Input

	// High is the multiplier for the reverberation time within the high band. (0..1)
	High Input

	// LowCut is the frequency at which the crossover between
	// the low and mid bands of the reverb occurs. (100..6000)
	LowCut Input

	// HighCut is the frequency at which the crossover between
	// the mid and high bands of the reverb occurs. (1000..10000)

	HighCut Input
}

func (jpv *JPverb) defaults() {
	if jpv.T60 == nil {
		jpv.T60 = C(1)
	}
	if jpv.Damp == nil {
		jpv.Damp = C(0)
	}
	if jpv.Size == nil {
		jpv.Size = C(1)
	}
	if jpv.EarlyDiff == nil {
		jpv.EarlyDiff = C(0.707)
	}
	if jpv.ModDepth == nil {
		jpv.ModDepth = C(0.1)
	}
	if jpv.ModFreq == nil {
		jpv.ModFreq = C(2)
	}
	if jpv.Low == nil {
		jpv.Low = C(1)
	}
	if jpv.Mid == nil {
		jpv.Mid = C(1)
	}
	if jpv.High == nil {
		jpv.High = C(1)
	}
	if jpv.LowCut == nil {
		jpv.LowCut = C(500)
	}
	if jpv.HighCut == nil {
		jpv.HighCut = C(2000)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil this method will panic.
func (jpv JPverb) Rate(rate int8) Input {
	CheckRate(rate)
	if jpv.In == nil {
		panic("JPverb requires an input signal")
	}
	(&jpv).defaults()

	var in1, in2 Input

	switch x := jpv.In.(type) {
	case Inputs:
		if len(x) == 0 {
			panic("JPverb requires an input signal")
		}
		in1, in2 = x[0], x[len(x)-1]
	default:
		in1, in2 = jpv.In, jpv.In
	}
	return NewInput("JPverbRaw", rate, 0, 2, in1, in2, jpv.T60, jpv.Damp, jpv.Size, jpv.EarlyDiff, jpv.ModDepth, jpv.ModFreq, jpv.Low, jpv.Mid, jpv.High, jpv.LowCut, jpv.HighCut)
}

// defJPverb is a synthdef that exposes the fields of the JPverb ugen.
func defJPverb(params Params) Ugen {
	var (
		in        = params.Add("in", 0)
		out       = params.Add("out", 0)
		t60       = params.Add("t60", 1)
		damp      = params.Add("damp", 0)
		size      = params.Add("size", 1)
		earlyDiff = params.Add("earlyDiff", 0.707)
		modDepth  = params.Add("modDepth", 0.1)
		modFreq   = params.Add("modFreq", 2)
		low       = params.Add("low", 1)
		mid       = params.Add("mid", 1)
		high      = params.Add("high", 1)
		lowcut    = params.Add("lowcut", 500)
		highcut   = params.Add("highcut", 2000)
	)
	return Out{
		Bus: out,
		Channels: JPverb{
			In: In{
				Bus:         in,
				NumChannels: 2,
			}.Rate(AR),
			T60:       t60,
			Damp:      damp,
			Size:      size,
			EarlyDiff: earlyDiff,
			ModDepth:  modDepth,
			ModFreq:   modFreq,
			Low:       low,
			Mid:       mid,
			High:      high,
			LowCut:    lowcut,
			HighCut:   highcut,
		}.Rate(AR),
	}.Rate(AR)
}
