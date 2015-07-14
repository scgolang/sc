package ugens

// import "fmt"
import . "github.com/scgolang/sc/types"

// Out
type Out struct {
	Bus      C
	Channels Input
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause
// a runtime panic.
func (self Out) Rate(rate int8) *UgenNode {
	checkRate(rate)
	return NewUgenNode("Out", rate, 0, 1, self.Bus, self.Channels)
}
