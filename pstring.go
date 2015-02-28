package sc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// pstring is a pascal-format string, which is a byte containing
// the string length, followed by the bytes of the string
type pstring struct {
	length int8
	string string
}

func (self *pstring) String() string {
	return self.string
}

// Equal determines if one pstring equals another
func (self *pstring) Equals(pstr pstring) bool {
	return self.string == pstr.string
}

// Write writes a pstring to an io.Writer
func (self *pstring) Write(w io.Writer) error {
	e := binary.Write(w, byteOrder, self.length)
	if e != nil {
		return e
	}
	_, e = w.Write(bytes.NewBufferString(self.string).Bytes())
	return e
}

// newPstring create a new pstring
func newPstring(s string) *pstring {
	length := len(s)
	p := pstring{int8(length), s}
	return &p
}

// readPstring reads a pstring from an io.Reader
func readPstring(r io.Reader) (*pstring, error) {
	var length int8
	e := binary.Read(r, byteOrder, &length)
	if e != nil {
		return nil, e
	}
	s := make([]byte, length)
	read, e := r.Read(s)
	if e != nil {
		return nil, e
	}
	if read != int(length) {
		return nil, fmt.Errorf("could not read %d bytes", length)
	}
	ps := pstring{
		length,
		bytes.NewBuffer(s).String(),
	}
	return &ps, nil
}
