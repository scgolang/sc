package gosc

type NodeVisitor func(node Node)

type Node struct {
	value interface{}
}

// Constant returns true if this node is a constant
func (n *Node) Constant() bool {
	if _, ok := n.GetInt32(); !ok {
		return false
	}
	if _, ok := n.GetFloat32(); !ok {
		return false
	}
	return true
}

func (n *Node) GetInt32() (int32, bool) {
	val, ok := n.value.(int32)
	return val, ok
}

func (n *Node) GetFloat32() (float32, bool) {
	val, ok := n.value.(float32)
	return val, ok
}

func (n *Node) OnFree(f func()) *Node {
	return n
}

func NewNode(val interface{}) *Node {
	n := Node{val}
	return &n
}

type Synth struct {
	Node
}

type Group struct {
	Node
	nodes []Node
}

type RootNode struct {
	Group
}

// Out
// |
// ---- Constant(0)
// |
// ---- SinOsc
//      |
//      ---- Constant(440)
type UgenGraph struct {
	root RootNode
}

func (self *UgenGraph) Walk(visitor NodeVisitor) {
	// for _, n := range self.root.Group.nodes {
	// }
}

func (self *UgenGraph) walkFrom(visitor NodeVisitor, node Node) {
}

func NewUgenGraph() *UgenGraph {
	nodes := make([]Node, 0)
	ugg := UgenGraph{
		RootNode{
			Group{Node{nil}, nodes},
		},
	}
	return &ugg
}

type UgenGraphFunc func(args ...interface{}) UgenGraph
