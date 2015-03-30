package sc

type cstack []float32

func (self *cstack) Push(val float32) {
	for _, f := range *self {
		if f == val {
			return
		}
	}
	*self = append(*self, val)
}

func (self *cstack) Pop() (float32, bool) {
	l := len(*self)
	if l == 0 {
		return 0, false
	}
	f := (*self)[l-1]
	*self = (*self)[0:l-1]
	return f, true
}

func newCstack() *cstack {
	stack := cstack(make([]float32, 0))
	return &stack
}
