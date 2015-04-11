package pattern

import "math/rand"

// Xrand emits length values randomly selected from the provided
// array, where an element will not be emitted two times in a row
func Xrand(length int, values ...interface{}) chan interface{} {
	l := len(values)
	c := make(chan interface{})
	last := -1
	var n int
	go func() {
		for i := 0; i < length; i++ {
			n = rand.Intn(l)
			// ensure that n is not the last one
			for n == last {
				n = rand.Intn(l)
			}
			c <-values[n]
			last = n
		}
		close(c)
	}()
	return c
}
