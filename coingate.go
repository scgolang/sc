package sc

// CoinGate tosses a coin every time it receives a trigger
// and either passes the trigger or doesn't.
type CoinGate struct {
	// Value between 0.0 and 1.0 determines probability of either possibilities.
	Prob Input

	// Trigger signal.
	In Input
}

func (c *CoinGate) defaults() {
	if c.Prob == nil {
		c.Prob = C(0.5)
	}
	if c.In == nil {
		c.In = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
func (c CoinGate) Rate(rate int8) Input {
	CheckRate(rate)
	(&c).defaults()
	return NewInput("CoinGate", rate, 0, 1, c.Prob, c.In)
}
