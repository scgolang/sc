package sc

import "fmt"

// Ugen rates.
//   - IR is initialization rate.
//   - KR is control rate.
//   - AR is audio rate.
// See http://doc.sccode.org/Tutorials/Mark_Polishook_tutorial/08_Rates.html.
const (
	IR = 0
	KR = 1
	AR = 2
)

// CheckRate causes a panic if rate is not IR, KR, or AR.
func CheckRate(rate int8) {
	if rate != IR && rate != KR && rate != AR {
		panic(fmt.Errorf("Unsupported rate %d", rate))
	}
}

// Rater is the interface of a Ugen that can compute its output at different rates.
type Rater interface {
	Rate(int8) Input
}

// A returns a ugen input at audio rate.
func A(r Rater) Input {
	return r.Rate(AR)
}

// I returns a ugen input at initialization rate.
func I(r Rater) Input {
	return r.Rate(IR)
}

// K returns a ugen input at control rate.
func K(r Rater) Input {
	return r.Rate(KR)
}
