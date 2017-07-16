package sc

// Hasher returns a unique output value from -1 to +1 for each input value
// according to a hash function.
// The same input value will always produce the same output value.
// The input need not be in the range -1 to +1.
type Hasher struct {
	In Input
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If the input signal is nil then this method will cause a runtime panic.
func (hasher Hasher) Rate(rate int8) Input {
	CheckRate(rate)
	if hasher.In == nil {
		panic("Hasher requires an input signal")
	}
	return NewInput("Hasher", rate, 0, 1, hasher.In)
}
