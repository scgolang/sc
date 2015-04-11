package pattern

import "math/rand"

// Shuf emits a shuffled version of the Values list
// Repeats times
func Shuf(repeats int, values ...interface{}) chan interface{} {
	l := len(values)
	c := make(chan interface{})
	perm := rand.Perm(l)
	go func() {
		for r := 0; r < repeats; r++ {
			for i := range perm {
				c <-values[i]
			}
		}
		close(c)
	}()
	return c
}
