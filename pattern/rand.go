package pattern

import "math/rand"

// Rand emits randomly selected values from an array a
// certain number of times
func Rand(length int, values ...interface{}) chan interface{} {
	l := len(values)
	c := make(chan interface{})
	go func() {
		for i := 0; i < length; i++ {
			c <-values[rand.Intn(l)]
		}
		close(c)
	}()
	return c
}
