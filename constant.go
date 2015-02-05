package sc

// Constant constant node in a UgenGraph
type Constant interface {
	// Constant index
	Index() int32
	// Constant value
	Value() float32
}

type constant struct {
	index int32
	value float32
}

func (self *constant) Index() int32 {
	return self.index
}

func (self *constant) Value() float32 {
	return self.value
}

func NewConstant(index int, value float32) Constant {
	return &constant{int32(index), value}
}
