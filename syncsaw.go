package sc

// SyncSaw is a sawtooth wave that is hard synched to a fundamental pitch.
// This produces an effect similar to moving formants or pulse width modulation.
// The sawtooth oscillator has its phase reset when the sync oscillator completes a cycle.
// This is not a band limited waveform, so it may alias.
type SyncSaw struct {
	// SyncFreq is the frequency of the fundamental.
	SyncFreq Input

	// SawFreq is the frequency of the slave synched sawtooth wave.
	// It should always be greater than SyncFreq.
	SawFreq Input
}

func (s *SyncSaw) defaults() {
	if s.SyncFreq == nil {
		s.SyncFreq = C(440)
	}
	if s.SawFreq == nil {
		s.SawFreq = C(440)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (s SyncSaw) Rate(rate int8) Input {
	CheckRate(rate)
	(&s).defaults()
	return NewInput("SyncSaw", rate, 0, 1, s.SyncFreq, s.SawFreq)
}
