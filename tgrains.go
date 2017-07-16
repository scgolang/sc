package sc

// TGrains is a buffer granulator
type TGrains struct {
	// Number of output channels.
	NumChannels int
	// At each trigger, the following arguments are sampled and used as the arguments of a new grain.
	// A trigger occurs when a signal changes from non-positive to positive value.
	// If the trigger is audio rate then the grains will start with sample accuracy.
	Trigger Input
	// The index of the buffer to use. It must be a one channel (mono) buffer.
	BufNum Input
	// 1.0 is normal, 2.0 is one octave up, 0.5 is one octave down -1.0 is backwards normal rate… etc.
	// (renamed to GRate, to avoid conflicts with Rate method)
	GRate Input
	// The position in the buffer in seconds at which the grain envelope will reach maximum amplitude.
	CenterPos Input
	// Duration of the grain in seconds.
	Dur Input
	// Determines where to pan the output.
	// If numChannels = 1, the pan argument is ignored.
	// If numChannels = 2, panning is similar to Pan2.
	// If numChannels > 2, panning is the same as PanAz.
	Pan Input
	// Amplitude of the grain.
	Amp Input
	// 1, 2, or 4. Determines whether the grain uses (1) no interpolation, (2) linear interpolation, or (4) cubic interpolation.
	Interp Input
}

func (tg *TGrains) defaults() {
	if tg.NumChannels == 0 {
		tg.NumChannels = 1
	}
	if tg.Trigger == nil {
		tg.Trigger = C(0)
	}
	if tg.BufNum == nil {
		tg.BufNum = C(0)
	}
	if tg.GRate == nil {
		tg.GRate = C(1)
	}
	if tg.CenterPos == nil {
		tg.CenterPos = C(0)
	}
	if tg.Dur == nil {
		tg.Dur = C(0.1)
	}
	if tg.Pan == nil {
		tg.Pan = C(0)
	}
	if tg.Amp == nil {
		tg.Amp = C(0.1)
	}
	if tg.Interp == nil {
		tg.Interp = C(4)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (tg TGrains) Rate(rate int8) Input {
	CheckRate(rate)
	(&tg).defaults()
	return NewInput("TGrains", rate, 0, tg.NumChannels, tg.Trigger, tg.BufNum, tg.GRate, tg.CenterPos, tg.Dur, tg.Pan, tg.Amp, tg.Interp)
}
