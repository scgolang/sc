package sc

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

// SynthDef
type SynthDef interface {
	// Dump write info about a synthdef to an io.Writer
	Dump(w io.Writer) error
	// Name returns the name of the synthdef
	Name() string
	// Load writes a synthdef file to disk and tells a server to load it
	Load(s Server) error
}

type synthDef struct {
	name               Pstring
	NumConstants       int32
	Constants          []float32
	NumParams          int32
	InitialParamValues []float32
	NumParamNames      int32
	ParamNames         []ParamName
	NumUgens           int32
	Ugens              []Ugen
	NumVariants        int16
	Variants           []Variant
}

func (self *synthDef) Name() string {
	return self.name.String
}

// Write writes a synthdef to an io.Writer
func (self *synthDef) Write(w io.Writer) error {
	if he := self.writeHead(w); he != nil {
		return he
	}
	return self.writeBody(w)
}

// Dump writes information about a SynthDef to an io.Writer
func (self *synthDef) Dump(w io.Writer) error {
	var e error
	
	fmt.Fprintf(w, "%-30s %s\n", "Name", self.name.String)
	// write constants
	fmt.Fprintf(w, "%-30s %d\n", "NumConstants", self.NumConstants)
	fmt.Fprintf(w, "%s\n", "Constants")
	for i := 0; int32(i) < self.NumConstants; i++ {
		fmt.Fprintf(w, "    %-26d %g\n", i, self.Constants[i])
	}
	// write params
	fmt.Fprintf(w, "%-30s %d\n", "NumParams", self.NumParams)
	if self.NumParams > 0 {
		fmt.Fprintf(w, "%-30s\n", "Params:")
		for i := 0; int32(i) < self.NumParams; i++ {
			fmt.Fprintf(w, "    Initial Value %-12d %g\n", i, self.InitialParamValues[i])
		}
	}
	// write param names
	fmt.Fprintf(w, "%-30s %d\n", "NumParamNames", self.NumParamNames)
	if self.NumParamNames > 0 {
		fmt.Fprintf(w, "%-30s\n", "Param Names:")
		for i := 0; int32(i) < self.NumParamNames; i++ {
			fmt.Fprintf(w, "    %-26d %g\n", i, self.ParamNames[i])
		}
	}
	// write ugens and variants
	fmt.Fprintf(w, "%-30s %d\n", "NumUgens", self.NumUgens)
	fmt.Fprintf(w, "%-30s %d\n", "NumVariants", self.NumVariants)
	if self.NumUgens > 0 {
		for i := 0; int32(i) < self.NumUgens; i++ {
			fmt.Fprintf(w, "\nUgen %d:\n", i)
			e = self.Ugens[i].Dump(w)
			if e != nil {
				return e
			}
		}
	}
	if self.NumVariants > 0 {
		fmt.Fprintf(w, "%-30s\n", "Variants:")
		for i := 0; int16(i) < self.NumVariants; i++ {
			e = self.Ugens[i].Dump(w)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

// write a synthdef header
func (self *synthDef) writeHead(w io.Writer) error {
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
func (self *synthDef) writeBody(w io.Writer) error {
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

func (self *synthDef) Load(s Server) error {
	return nil
}

// ReadSynthDef reads a synthdef from an io.Reader
func ReadSynthDef(r io.Reader) (SynthDef, error) {
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
	ugens := make([]Ugen, numUgens)
	for i := 0; int32(i) < numUgens; i++ {
		ugen, er := ReadUgen(r)
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
	synthDef := synthDef{
		*defName,
		numConstants,
		constants,
		numParams,
		initialValues,
		numParamNames,
		paramNames,
		numUgens,
		ugens,
		numVariants,
		variants,
	}
	return &synthDef, nil
}

// NewSynthDef creates a new SynthDef from a UgenGraphFunc
func NewSynthDef(name string, f UgenGraphFunc) SynthDef {
	// this function has to be able to traverse a ugen
	// graph and turn it into a synth def
	return nil
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

type Variant struct {
	Name Pstring
	InitialParamValues []float32
}

// ReadVariant read a Variant from an io.Reader
func ReadVariant(r io.Reader, numParams int32) (*Variant, error) {
	name, err := ReadPstring(r)
	if err != nil {
		return nil, err
	}
	paramValues := make([]float32, numParams)
	for i := 0; int32(i) < numParams; i++ {
		err = binary.Read(r, byteOrder, &paramValues[i])
		if err != nil {
			return nil, err
		}
	}
	v := Variant{*name, paramValues}
	return &v, nil
}
