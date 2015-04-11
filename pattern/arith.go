package pattern

// Arith emits an arithmetic series.
// If length is not Inf or a positive integer, then
// a runtime panic will occur.
func Arith(start, step float64, length int) chan float64 {
	if length != Inf && length < 1 {
		panic("length must be positive or Inf")
	}
	c := make(chan float64)
	cur := start
	go func() {
		for i := 0; length == Inf || i < length; i++ {
			c <-cur
			cur = cur + step
		}
		close(c)
	}()
	return c
}
