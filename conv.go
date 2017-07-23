package sc

import (
	"math"
)

// Midicps converts midi note values to frequency in Hz
func Midicps(note float32) float32 {
	return float32(440) * float32(math.Pow(2, float64(note-69)/12.0))
}

// Cpsmidi converts frequency in Hz to midi note values.
func Cpsmidi(cps float32) float32 {
	return float32((12 * math.Log2(float64(cps)/440)) + 69)
}
