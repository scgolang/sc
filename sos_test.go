package sc

import (
	"math"
	"testing"
)

func TestSOS(t *testing.T) {
	defName := "SOSTest"

	// var rho, theta, b1, b2;
	// theta = MouseX.kr(0.2pi, pi);
	// rho = MouseY.kr(0.6, 0.99);
	// b1 = 2.0 * rho * cos(theta);
	// b2 = rho.squared.neg;
	// Out.ar(0, SOS.ar(LFSaw.ar(200, 0, 0.1), 1.0, 0.0, 0.0, b1, b2));
	compareAndWriteStructure(t, defName, NewSynthdef(defName, func(p Params) Ugen {
		var (
			theta = K(MouseX{
				Min: C(0.2 * math.Pi),
				Max: C(math.Pi),
			})
			rho = K(MouseY{
				Min: C(0.6),
				Max: C(0.99),
			})
			b1 = C(2.0).Mul(rho).Mul(theta.Cos())
			b2 = rho.Squared().Neg()
		)
		saw := A(LFSaw{
			Freq:   C(200),
			Iphase: C(0),
		}).Mul(C(0.1))

		return Out{
			Bus: C(0),
			Channels: A(SOS{
				In: saw,
				A0: C(1.0),
				A1: C(0.0),
				A2: C(0.0),
				B1: b1,
				B2: b2,
			}),
		}.Rate(AR)
	}))
}
