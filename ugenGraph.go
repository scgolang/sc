package gosc

type Synth struct {
	Node
}

// Out
// |
// ---- Constant(0)
// |
// ---- SinOsc
//      |
//      ---- Constant(440)
type UgenGraph struct {
	root *Node
}

func (self *UgenGraph) Walk(visitor NodeVisitor) {
	// for _, n := range self.root.Group.nodes {
	// }
}

func (self *UgenGraph) walkFrom(visitor NodeVisitor, node Node) {
}

func NewUgenGraph(root *Node) *UgenGraph {
	ugg := UgenGraph{root}
	return &ugg
}

type UgenGraphFunc func(args ...interface{}) UgenGraph
