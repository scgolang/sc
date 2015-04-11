package pattern

// Seq cycles over a list of values repeats times.
func Seq(repeats int, values ...interface{}) chan interface{} {
	l := len(values)
	c := make(chan interface{})
	go func() {
		for r := 0; r < repeats; r++ {
			for i := 0; i < l; i++ {
				c <- values[i]
			}
		}
		close(c)
	}()
	return c
}
