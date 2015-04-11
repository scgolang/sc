package pattern

import "math/rand"

// White emits length random numbers between lo and hi.
// If length is not Inf or a positive integer, then
// a runtime panic will occur.
// If hi is not greater than lo a runtime panic will occur.
func White(lo, hi float64, length int) chan float64 {
	if length != Inf && length < 1 {
		panic("length must be positive or Inf")
	}
	if hi <= lo {
		panic("hi must be greater than lo")
	}
	c := make(chan float64)
	m := hi - lo
	b := lo
	go func() {
		for i := 0; length == Inf || i < length; i++ {
			c <- (m * rand.Float64()) + b
		}
		close(c)
	}()
	return c
}
