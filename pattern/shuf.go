package pattern

import "math/rand"

// Shuf emits a shuffled version of the Values list
// Repeats times
type Shuf struct {
	Repeats int
	Values  []interface{}
}

func (self Shuf) Stream() chan interface{} {
	l := len(self.Values)
	c := make(chan interface{})
	perm := rand.Perm(l)
	go func() {
		for r := 0; r < self.Repeats; r++ {
			for i := range perm {
				c <-self.Values[i]
			}
		}
		close(c)
	}()
	return c
}
