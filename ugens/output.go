package ugens

type output int8

func (self output) Rate() int8 {
	return int8(self)
}
