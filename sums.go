package sc

// Sum3 sums three signals. Should be used via Mix.
func Sum3(rate int8, in1, in2, in3 Input) Input {
	CheckRate(rate)
	return UgenInput("Sum3", rate, 0, 1, in1, in2, in3)
}

// Sum4 sums four signals. Should be used via Mix.
func Sum4(rate int8, in1, in2, in3, in4 Input) Input {
	CheckRate(rate)
	return UgenInput("Sum4", rate, 0, 1, in1, in2, in3, in4)
}
