package ugens

import . "github.com/briansorahan/sc/types"

// Out
type Out struct {
	Bus      C
	Channels Input
}

func (self Out) Rate(rate int8) *BaseNode {
	n := NewNode("Out", rate, 0)
	n.addInputs(self.Bus, self.Channels)
	return n
}
