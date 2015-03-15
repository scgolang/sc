package ugens

import . "github.com/briansorahan/sc/types"

// Out
type Out struct {
	Bus      C
	Channels Input
}

func (self Out) Rate(rate int8) *BaseNode {
	// If self.Channels is an array, we need to expand it
	// to multiple individual inputs
	return NewNode("Out", rate, 0, self.Bus, self.Channels)
}
