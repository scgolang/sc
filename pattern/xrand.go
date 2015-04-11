package pattern

import "math/rand"

type Xrand struct {
	Length int
	Values []interface{}
}

func (self Xrand) Stream() chan interface{} {
	l := len(self.Values)
	c := make(chan interface{})
	last := -1
	var n int
	go func() {
		for i := 0; i < self.Length; i++ {
			n = rand.Intn(l)
			// ensure that n is not the last one
			for n == last {
				n = rand.Intn(l)
			}
			c <-self.Values[n]
			last = n
		}
		close(c)
	}()
	return c
}
