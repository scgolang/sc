package gosc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	SYNTHDEF_START        = "SCgf"
	SYNTHDEF_FILE_VERSION = 2
)

var byteOrder = binary.BigEndian

type SynthDef struct {
	Name Pstring
	NumConstants int32
	Constants []float32
	NumParams int32
	InitialParamValues []float32
	NumParamNames int32
	ParamNames []ParamName
	NumUgens int32
	Ugens []Ugen
}

type ParamName struct {
	Name Pstring
	Index int32
}

func (p *ParamName) Write(w io.Writer) error {
	if we := p.Name.Write(w); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, p.Index)
}

// write a synthdef header
func (self *SynthDef) writeHead(w io.Writer) error {
	_, we := w.Write(bytes.NewBufferString("SCgf").Bytes())
	if we != nil {
		return we
	}
	we = binary.Write(w, byteOrder, int32(SYNTHDEF_FILE_VERSION))
	if we != nil {
		return we
	}
	return binary.Write(w, byteOrder, int16(1))
}

// write a synthdef body
func (self *SynthDef) writeBody(w io.Writer) error {
	// write constants
	we := binary.Write(w, byteOrder, self.NumConstants)
	if we != nil {
		return we
	}
	for _, c := range self.Constants {
		if we = binary.Write(w, byteOrder, c); we != nil {
			return we
		}
	}
	// write parameters
	we = binary.Write(w, byteOrder, self.NumParams)
	if we != nil {
		return we
	}
	for _, p := range self.InitialParamValues {
		we = binary.Write(w, byteOrder, p)
		if we != nil {
			return we
		}
	}
	we = binary.Write(w, byteOrder, self.NumParamNames)
	if we != nil {
		return we
	}
	for _, p := range self.ParamNames {
		if we = p.Write(w); we != nil {
			return we
		}
	}
	// number of ugens
	if binary.Write(w, byteOrder, int32(1));  we != nil {
		return we
	}

	// TODO: write ugens

	// number of variants
	if we = binary.Write(w, byteOrder, int16(0)); we != nil {
		return we
	}
	return nil
}

// write a synthdef
func (self *SynthDef) Write(w io.Writer) error {
	if he := self.writeHead(w); he != nil {
		return he
	}
	return self.writeBody(w)
}

// read a synthdef
func ReadSynthDef(r io.Reader) (*SynthDef, error) {
	startLen := len(SYNTHDEF_START)
	start := make([]byte, startLen)
	read, er := r.Read(start)
	if er != nil {
		return nil, er
	}
	if read != startLen {
		return nil, fmt.Errorf("bad synthdef")
	}
	if bytes.NewBuffer(start).String() != SYNTHDEF_START {
		return nil, fmt.Errorf("bad synthdef")
	}
	var version int32
	er = binary.Read(r, byteOrder, version)
	if er != nil {
		return nil, er
	}
	return nil, nil
}

func WriteSynthDef(w io.Writer) error {
	return nil
}
