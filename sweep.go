package sc

// Sweep starts a linear raise by rate/sec from zero when trig input crosses from
// non-positive to positive.
// When rate == 1, Sweep may be used to get a continually-updating measurement
// of the time (in seconds) since the last trigger.
type Sweep struct {
	// Trig triggers when trig input crosses from non-positive to positive.
	Trig Input

	// RaiseRate is the rate/sec raise rate.
	RaiseRate Input
}

func (sweep *Sweep) defaults() {
	if sweep.Trig == nil {
		sweep.Trig = C(0)
	}
	if sweep.RaiseRate == nil {
		sweep.RaiseRate = C(1)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (sweep Sweep) Rate(rate int8) Input {
	CheckRate(rate)
	(&sweep).defaults()
	return NewInput("Sweep", rate, 0, 1, sweep.Trig, sweep.RaiseRate)
}
