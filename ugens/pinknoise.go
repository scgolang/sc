package ugens

// PinkNoise
type PinkNoise struct {
}

func (self PinkNoise) Rate(rate int8) *BaseNode {
	checkRate(rate)
	n := newNode("PinkNoise", rate, 0)
	return n
}
