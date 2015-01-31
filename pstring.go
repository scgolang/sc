package gosc

// Pstring is a pascal-format string, which is a byte containing
// the string length, followed by the bytes of the string
type Pstring struct {
	Length int8
	String string
}

// NewPstring create a new Pstring
func NewPstring(s string) Pstring {
	length := len(s)
	return Pstring{int8(length), s}
}
