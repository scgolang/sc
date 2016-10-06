package sc

import (
	"math"
)

// Midicps converts midi note values to frequency in Hz
func Midicps(note float32) float32 {
	return float32(440) * float32(math.Pow(2, float64(note-69)/12.0))
}
