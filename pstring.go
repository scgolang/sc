package gosc

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Pstring is a pascal-format string, which is a byte containing
// the string length, followed by the bytes of the string
type Pstring struct {
	Length int8
	String string
}

func (self *Pstring) Write(w io.Writer, order binary.ByteOrder) error {
	binary.Write(w, order, self.Length)
	_, e := w.Write(bytes.NewBufferString(self.String).Bytes())
	return e
}

// NewPstring create a new Pstring
func NewPstring(s string) Pstring {
	length := len(s)
	return Pstring{int8(length), s}
}
