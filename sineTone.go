package gosc

import (
	"bytes"
	"encoding/binary"
	"io"
)

const (
	Name = "SineTone"
)

type sineTone struct {
}

func (self *sineTone) Send(addr NetAddr) error {
	return nil
}

func (self *sineTone) Write(w io.Writer) error {
	he := self.writeHeader(w)
	if he != nil {
		return he
	}
	return nil
}

func (self *sineTone) writeHeader(w io.Writer) error {
	_, we := w.Write(bytes.NewBufferString("SCgf").Bytes())
	if we != nil {
		return we
	}
	we = binary.Write(w, byteOrder, int32(2))
	if we != nil {
		return we
	}
	we = binary.Write(w, byteOrder, int16(1))
	if we != nil {
		return we
	}
	return nil
}

func SineTone() SynthDef {
	tt := sineTone{}
	return &tt
}
