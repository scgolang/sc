package sc

// Mix will mix an array of channels down to a single channel.
func Mix(rate int8, inputs []Input) Input {
	CheckRate(rate)
	switch len(inputs) {
	case 0:
		panic("can not mix empty inputs slice")
	case 1:
		return inputs[0]
	case 2:
		return inputs[0].Add(inputs[1])
	case 3:
		return Sum3(rate, inputs[0], inputs[1], inputs[2])
	case 4:
		return Sum4(rate, inputs[0], inputs[1], inputs[2], inputs[3])
	default:
		head := []Input{Sum4(rate, inputs[0], inputs[1], inputs[2], inputs[3])}
		return Mix(rate, append(head, inputs[4:]...))
	}
}
