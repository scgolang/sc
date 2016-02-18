package sc

// Interpolation is a smoothing strategy for delay lines.
// Possible values are
//   - InterpolationNone
//   - InterpolationLinear
//   - InterpolationCubic
type Interpolation int

// Interpolation types.
const (
	InterpolationNone Interpolation = iota
	InterpolationLinear
	InterpolationCubic
)
