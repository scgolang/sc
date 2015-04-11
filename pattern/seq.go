package pattern

// Seq cycles over a list of values. The repeats variable gives
// the number of times to repeat the entire list.
type Seq struct {
	Repeats int
	Values  []interface{}
}

func (self Seq) Stream() chan interface{} {
	l := len(self.Values)
	c := make(chan interface{})
	go func() {
		for r := 0; r < self.Repeats; r++ {
			for i := 0; i < l; i++ {
				c <- self.Values[i]
			}
		}
		close(c)
	}()
	return c
}
