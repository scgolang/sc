package types

// Input is implemented by any value that can serve as a
// ugen input. This includes synthdef parameters,
// constants, and other ugens.
type Input interface {
	Mul(val Input) Input
	Add(val Input) Input
	IsMulti() bool
}
