package sc

type Stack interface {
	Push(val interface{})
	Pop() interface{}
}

type stack struct {
	l []interface{}
}

func (self *stack) Push(val interface{}) {
	self.l = append(self.l, val)
}

func (self *stack) Pop() interface{} {
	l := len(self.l)
	if l == 0 {
		return nil
	}
	el := self.l[l-1]
	self.l = self.l[0:l-1]
	return el
}

func NewStack() Stack {
	s := stack{make([]interface{}, 0)}
	return &s
}
