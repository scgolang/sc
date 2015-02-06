package sc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

const (
	SYNTHDEF_START   = "SCgf"
	SYNTHDEF_VERSION = 2
)

var byteOrder = binary.BigEndian

// SynthDefRep defines the structure of synth def data as defined
// in http://doc.sccode.org/Reference/Synth-Definition-File-Format.html
type SynthDefRep struct {
	// Name is the name of the synthdef
	Name string `json:"name"`

	// Constants is a list of constants that appear in the synth def
	Constants []float32 `json:"constants"`

	// InitialParamValues is an array of initial values for synth params
	InitialParamValues []float32 `json:"initialParamValues"`

	// ParamNames contains the names of the synth parameters
	ParamNames []ParamName `json:"paramNames"`

	// Ugens is the list of ugens that appear in the synth def
	Ugens []UgenRep `json:"ugens"`

	// Variants is the list of variants contained in the synth def
	Variants []Variant `json:"variants"`
}

// Write writes a synthdef to an io.Writer
func (self *SynthDefRep) Write(w io.Writer) error {
	if he := self.writeHead(w); he != nil {
		return he
	}
	return self.writeBody(w)
}

// Dump writes json-formatted information about a SynthDefRep to an io.Writer
func (self *SynthDefRep) Dump(w io.Writer) error {
	dec := json.NewEncoder(w)
	return dec.Encode(self)
}

// write a synthdef header
func (self *SynthDefRep) writeHead(w io.Writer) error {
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
func (self *SynthDefRep) writeBody(w io.Writer) error {
	// write constants
	numConstants := int32(len(self.Constants))
	we := binary.Write(w, byteOrder, numConstants)
	if we != nil {
		return we
	}
	for _, c := range self.Constants {
		if we = binary.Write(w, byteOrder, c); we != nil {
			return we
		}
	}
	// write parameters
	numParams := int32(len(self.InitialParamValues))
	we = binary.Write(w, byteOrder, numParams)
	if we != nil {
		return we
	}
	for _, p := range self.InitialParamValues {
		we = binary.Write(w, byteOrder, p)
		if we != nil {
			return we
		}
	}
	numParamNames := int32(len(self.ParamNames))
	we = binary.Write(w, byteOrder, numParamNames)
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

func (self *SynthDefRep) Load(s Server) error {
	return nil
}

// readSynthDefRep reads a synthdef from an io.Reader
func readSynthDefRep(r io.Reader) (*SynthDefRep, error) {
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
	ugens := make([]UgenRep, numUgens)
	for i := 0; int32(i) < numUgens; i++ {
		ugen, er := readUgenRep(r)
		if er != nil {
			return nil, er
		}
		ugens[i] = *ugen
	}
	// read number of variants
	var numVariants int16
	er = binary.Read(r, byteOrder, &numVariants)
	if er != nil {
		return nil, er
	}
	// read variants
	variants := make([]Variant, numVariants)
	for i := 0; int16(i) < numVariants; i++ {
		v, er := ReadVariant(r, numParams)
		if er != nil {
			return nil, er
		}
		variants[i] = *v
	}
	// return a new synthdef
	synthDef := SynthDefRep{
		defName.String(),
		constants,
		initialValues,
		paramNames,
		ugens,
		variants,
	}
	return &synthDef, nil
}
