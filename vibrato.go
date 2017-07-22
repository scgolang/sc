package sc

// Vibrato is a slow frequency modulation.
// Consider the systematic deviation in pitch of a singer around a fundamental frequency,
// or a violinist whose finger wobbles in position on the fingerboard,
// slightly tightening and loosening the string to add shimmer to the pitch.
// There is often also a delay before vibrato is established on a note.
// This UGen models these processes; by setting more extreme settings,
// you can get back to the timbres of FM synthesis.
// You can also add in some noise to the vibrato rate and vibrato size (modulation depth)
// to make for a more realistic motor pattern.
// The vibrato output is a waveform based on a squared envelope shape with
// four stages marking out 0.0 to 1.0, 1.0 to 0.0, 0.0 to -1.0, and -1.0 back to 0.0.
// Vibrato rate determines how quickly you move through these stages.
type Vibrato struct {
	Freq           Input
	Speed          Input
	Depth          Input
	Delay          Input
	Onset          Input
	RateVariation  Input
	DepthVariation Input
	IPhase         Input
}

// freq: 440, rate: 6, depth: 0.02, delay: 0, onset: 0, rateVariation: 0.04, depthVariation: 0.1, iphase: 0
func (v *Vibrato) defaults() {
	if v.Freq == nil {
		v.Freq = C(440)
	}
	if v.Speed == nil {
		v.Speed = C(6)
	}
	if v.Depth == nil {
		v.Depth = C(0.02)
	}
	if v.Delay == nil {
		v.Delay = C(0)
	}
	if v.Onset == nil {
		v.Onset = C(0)
	}
	if v.RateVariation == nil {
		v.RateVariation = C(0.04)
	}
	if v.DepthVariation == nil {
		v.DepthVariation = C(0.1)
	}
	if v.IPhase == nil {
		v.IPhase = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (v Vibrato) Rate(rate int8) Input {
	CheckRate(rate)
	(&v).defaults()
	return NewInput("Vibrato", rate, 0, 1, v.Freq, v.Speed, v.Depth, v.Delay, v.Onset, v.RateVariation, v.DepthVariation, v.IPhase)
}
