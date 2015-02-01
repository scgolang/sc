package gosc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	SYNTHDEF_START   = "SCgf"
	SYNTHDEF_VERSION = 2
)

var byteOrder = binary.BigEndian

type SynthDef struct {
	Name               Pstring
	NumConstants       int32
	Constants          []float32
	NumParams          int32
	InitialParamValues []float32
	NumParamNames      int32
	ParamNames         []ParamName
	NumUgens           int32
	Ugens              []Ugen
}

type ParamName struct {
	Name  Pstring
	Index int32
}

func (p *ParamName) Write(w io.Writer) error {
	if we := p.Name.Write(w); we != nil {
		return we
	}
	return binary.Write(w, byteOrder, p.Index)
}

// ReadParamName reads a ParamName from an io.Reader
func ReadParamName(r io.Reader) (*ParamName, error) {
	name, err := ReadPstring(r)
	if err != nil {
		return nil, err
	}
	var idx int32
	err = binary.Read(r, byteOrder, &idx)
	if err != nil {
		return nil, err
	}
	pn := ParamName{*name, idx}
	return &pn, nil
}

// write a synthdef header
func (self *SynthDef) writeHead(w io.Writer) error {
	_, we := w.Write(bytes.NewBufferString("SCgf").Bytes())
	if we != nil {
		return we
	}
	we = binary.Write(w, byteOrder, int32(SYNTHDEF_VERSION))
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
	if binary.Write(w, byteOrder, int32(1)); we != nil {
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
	// read the type
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
	// read version
	var version int32
	er = binary.Read(r, byteOrder, &version)
	if er != nil {
		return nil, er
	}
	if version != SYNTHDEF_VERSION {
		return nil, fmt.Errorf("bad synthdef version %d", version)
	}
	// read number of synth defs
	var numDefs int16
	er = binary.Read(r, byteOrder, &numDefs)
	if er != nil {
		return nil, er
	}
	if numDefs != 1 {
		return nil, fmt.Errorf("multiple synthdefs not supported")
	}
	// read synthdef name
	defName, er := ReadPstring(r)
	if er != nil {
		return nil, er
	}
	// read number of constants
	var numConstants int32
	er = binary.Read(r, byteOrder, &numConstants)
	if er != nil {
		return nil, er
	}
	// read constants
	constants := make([]float32, numConstants)
	for i := 0; i < int(numConstants); i++ {
		er = binary.Read(r, byteOrder, &constants[i])
		if er != nil {
			return nil, er
		}
	}
	// read number of parameters
	var numParams int32
	er = binary.Read(r, byteOrder, &numParams)
	if er != nil {
		return nil, er
	}
	// read initial parameter values
	initialValues := make([]float32, numParams)
	for i := 0; i < int(numParams); i++ {
		er = binary.Read(r, byteOrder, &initialValues[i])
		if er != nil {
			return nil, er
		}
	}
	// read number of parameter names
	var numParamNames int32
	er = binary.Read(r, byteOrder, &numParamNames)
	if er != nil {
		return nil, er
	}
	// read param names
	paramNames := make([]ParamName, numParamNames)
	for i := 0; int32(i) < numParamNames; i++ {
		pn, er := ReadParamName(r)
		if er != nil {
			return nil, er
		}
		paramNames[i] = *pn
	}
	// read number of ugens
	var numUgens int32
	er = binary.Read(r, byteOrder, &numUgens)
	if er != nil {
		return nil, er
	}
	// read ugens
	
	// read number of variants
	// read variants
	fmt.Printf("read %s\n", defName)
	return nil, nil
}

func WriteSynthDef(w io.Writer) error {
	return nil
}
