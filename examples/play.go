package main

import . "github.com/scgolang/sc"
import . "github.com/scgolang/sc/types"
import . "github.com/scgolang/sc/ugens"

func main() {
	_ = Play(func() Ugen {
		noise := WhiteNoise{}.Rate(AR)
		dust := Dust{C(1)}.Rate(AR).Mul(C(0.5))
		decay := Decay{dust, C(0.2)}.Rate(AR).Mul(noise)
		sig := AllpassN{decay, C(0.2), C(0.2), C(3)}.Rate(AR)
		return Out{C(0), sig}.Rate(AR)
	})
}
