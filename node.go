package gosc

type NodeVisitor func(node Node)

type Node struct {
	value interface{}
	children []*Node
	onFree func()
}

func (self *Node) Add(child *Node) {
	self.children = append(self.children, child)
}

// Constant returns true if this node is a constant
func (self *Node) Constant() bool {
	if _, ok := self.GetInt32(); !ok {
		return false
	}
	if _, ok := self.GetFloat32(); !ok {
		return false
	}
	return true
}

func (self *Node) GetInt32() (int32, bool) {
	val, ok := self.value.(int32)
	return val, ok
}

func (self *Node) GetFloat32() (float32, bool) {
	val, ok := self.value.(float32)
	return val, ok
}

func (self *Node) OnFree(f func()) {
	self.onFree = f
}

func IsNode(val interface{}) bool {
	_, ok := val.(Node)
	return ok
}

func NewNode(val interface{}) *Node {
	n := Node{
		val,
		make([]*Node, 0),
		func(){},
	}
	return &n
}
