package pattern

// Seq cycles over a list of values. The repeats variable gives
// the number of times to repeat the entire list.
type Seq struct {
	Repeats int
	Values  []interface{}
}

func (self Seq) Stream(ticks Ticks) Values {
	l := len(self.Values)
	vc := make(chan interface{})
	go func() {
		i, repeats := 0, 0
		for _ = range ticks {
			vc <- self.Values[i]
			i = i + 1
			if i == l {
				repeats = repeats + 1
				if repeats == self.Repeats {
					// done
					close(vc)
					break
				}
				i = 0
			}
		}
	}()
	return vc
}
