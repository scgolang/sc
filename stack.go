package sc

type stack struct {
	l []interface{}
}

func (s *stack) Push(val interface{}) {
	s.l = append(s.l, val)
}

func (s *stack) Pop() interface{} {
	l := len(s.l)
	if l == 0 {
		return nil
	}
	el := s.l[l-1]
	s.l = s.l[0 : l-1]
	return el
}

func (s *stack) Size() int {
	return len(s.l)
}

func newStack() *stack {
	s := stack{make([]interface{}, 0)}
	return &s
}
