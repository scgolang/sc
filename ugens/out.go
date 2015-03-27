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
	return NewNode("Out", rate, 0, self.Bus, self.Channels)
}
