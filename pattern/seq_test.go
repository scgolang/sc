package pattern

import "reflect"
import "testing"
import "time"

func TestSeq(t *testing.T) {
	d := 20 * time.Millisecond
	testReps(1, d, t)
	testReps(3, d, t)
}

func testReps(repeats int, dur time.Duration, t *testing.T) {
	// one repeat
	list := []interface{}{"foo", "bar", "baz"}
	pat := Seq{repeats, list}
	tc := make(chan uint64)
	go func() {
		i := 0
		for _ = range time.NewTicker(20 * time.Millisecond).C {
			tc <-uint64(i)
			i = i + 1
		}
		
	}()
	str := pat.Stream(Ticks(tc))
	l := make([]interface{}, 0)
	for v := range str {
		l = append(l, v)
	}

	single := []interface{}{"foo", "bar", "baz"}
	expect := make([]interface{}, 0)
	for r := 0; r < repeats; r++ {
		expect = append(expect, single...)
	}

	if !reflect.DeepEqual(l, expect) {
		t.Fatal("stream did not generate expected list")
	}
}
