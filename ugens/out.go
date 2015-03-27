package ugens

// import "fmt"
import . "github.com/briansorahan/sc/types"

// Out
type Out struct {
	Bus      C
	Channels Input
}

func (self Out) Rate(rate int8) *Node {
	checkRate(rate)

	// If self.Channels is an array, we need to expand it
	// to multiple individual inputs
	if multi, isMulti := self.Channels.(MultiInput); isMulti {
		channels := multi.InputArray()
		for _, c := range channels {
			if n, isNode := c.(*Node); isNode {
				n.IsOutput()
			}
		}
	}

	// 	return NewNode("Out", rate, 0, ins...)
	// } else {
	return NewNode("Out", rate, 0, self.Bus, self.Channels)
	// }
}
