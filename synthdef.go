package gosc

import (
	"encoding/binary"
	"io"
)

var byteOrder = binary.BigEndian

type SynthDef interface {
	Send(addr NetAddr) error
	Write(w io.Writer) error
}
