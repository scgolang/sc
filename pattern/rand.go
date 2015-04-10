package pattern

import "math/rand"

// Prand emits randomly selected values from an array a
// certain number of times
type Rand struct {
	Length int
	Values []interface{}
}

func (self Rand) Stream(ticks Ticks) Values {
	l := len(self.Values)
	c := make(chan interface{})
	go func() {
		for i := 0; i < self.Length; i++ {
			_ = <-ticks
			c <-self.Values[rand.Intn(l)]
		}
	}()
	return c
}
