package sc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/briansorahan/sc/types"
	"io"
)

const (
	SYNTHDEF_START   = "SCgf"
	SYNTHDEF_VERSION = 2
)

var byteOrder = binary.BigEndian

// synthdef defines the structure of synth def data as defined
// in http://doc.sccode.org/Reference/Synth-Definition-File-Format.html
type synthdef struct {
	// Name is the name of the synthdef
	Name string `json:"name" xml:"Name,attr"`

	// Constants is a list of constants that appear in the synth def
	Constants []float32 `json:"constants" xml:"Constants>Constant"`

	// InitialParamValues is an array of initial values for synth params
	InitialParamValues []float32 `json:"initialParamValues" xml:"InitialParamValues>initialParamValue"`

	// ParamNames contains the names of the synth parameters
	ParamNames []ParamName `json:"paramNames" xml:"ParamNames>ParamName"`

	// Ugens is the list of ugens that appear in the synth def
	Ugens []*ugen `json:"ugens" xml:"Ugens>Ugen"`

	// Variants is the list of variants contained in the synth def
	Variants []variant `json:"variants" xml:"Variants>Variant"`
}

// AddUgen returns an input pointing to either the (newly created)
// last position in the ugens array if this ugen has never been
// added before or the ugens existing position in the Ugens array
func (self *synthdef) AddUgen(u *ugen) *input {
	for i, v := range self.Ugens {
		if u == v {
			return &input{int32(i), 0}
		}
	}
	idx := len(self.Ugens)
	self.Ugens = append(self.Ugens, u)
	return &input{int32(idx), 0}
}

// AddConstant returns an input pointing to either the (newly created)
// last position in the constants array if this constant has never been
// added before or the constants existing position in the Constants array
func (self *synthdef) AddConstant(c float32) *input {
	for i, d := range self.Constants {
		if c == d {
			return &input{-1, int32(i)}
		}
	}
	idx := len(self.Constants)
	self.Constants = append(self.Constants, c)
	return &input{-1, int32(idx)}
}

func (self *synthdef) WriteJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(self)
}

// readsynthdef reads a synthdef from an io.Reader
func readsynthdef(r io.Reader) (*synthdef, error) {
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
	ugens := make([]*ugen, numUgens)
	for i := 0; int32(i) < numUgens; i++ {
		ugen, er := readugen(r)
		if er != nil {
			return nil, er
		}
		ugens[i] = ugen
	}
	// read number of variants
	var numVariants int16
	er = binary.Read(r, byteOrder, &numVariants)
	if er != nil {
		return nil, er
	}
	// read variants
	variants := make([]variant, numVariants)
	for i := 0; int16(i) < numVariants; i++ {
		v, er := ReadVariant(r, numParams)
		if er != nil {
			return nil, er
		}
		variants[i] = *v
	}
	// return a new synthdef
	synthDef := synthdef{
		defName.String(),
		constants,
		initialValues,
		paramNames,
		ugens,
		variants,
	}
	return &synthDef, nil
}

func newsynthdef(name string) *synthdef {
	def := synthdef{
		name,
		make([]float32, 0),
		make([]float32, 0),
		make([]ParamName, 0),
		make([]*ugen, 0),
		make([]variant, 0),
	}
	return &def
}

// NewSynthdef creates a synthdef by traversing a ugen graph
//
// Out.Ar(0, SinOsc.Ar(440))
//
// Example 1
// =========
// + Out (ugen 1)
// |
// +---+ 0 (constant 1)
// |
// +---+ SinOsc (ugen 0)
//     |
//     +---+ 440 (constant 0)
//
//
// Example 2
// =========
// + Out (ugen 1)
// |
// +---+ 0 (constant 1)
// |
// +---+ SinOsc (ugen 0)
//     |
//     +---+ 440 (constant 0)
//
func NewSynthdef(name string, graphFunc types.UgenGraphFunc) *synthdef {
	def := newsynthdef(name)
	params := newParams()
	root := graphFunc(params)
	flatten(root, def)
	return def
}
