package sc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	. "github.com/briansorahan/sc/types"
	"io"
	"os"
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

// AddParams will do nothing if there are no synthdef params.
// If there are synthdef params it will
// (1) Add their default values to initialParamValues
// (2) Add their names/indices to paramNames
// (3) Add a Control ugen as the first ugen
func (self *synthdef) AddParams(p Params) {
	paramList := p.List()
	numParams := len(paramList)
	self.InitialParamValues = make([]float32, numParams)
	self.ParamNames = make([]ParamName, numParams)
	for i, param := range paramList {
		self.InitialParamValues[i] = param.GetDefault()
		self.ParamNames[i] = ParamName{param.Name(),param.Index()}
	}
	if numParams > 0 {
		control := []*ugen{cloneUgen(p.Control())}
		self.Ugens = append(control, self.Ugens...)
	}
}

// Write writes a binary representation of a synthdef to an io.Writer.
// The binary representation written by this method is
// the data that scsynth expects at its /d_recv endpoint.
func (self *synthdef) Write(w io.Writer) error {
	return nil
}

// WriteJSON writes a json-formatted representation of a
// synthdef to an io.Writer
func (self *synthdef) WriteJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(self)
}

func (self *synthdef) CompareToFile(f string) (bool, error) {
	_, err := os.Open(f)
	if err != nil {
		return false, err
	}
	return false, nil
}

// ReadSynthdef reads a synthdef from an io.Reader
func ReadSynthdef(r io.Reader) (*synthdef, error) {
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
	defName, er := readPstring(r)
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
		pn, er := readParamName(r)
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
		v, er := readVariant(r, numParams)
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
func NewSynthdef(name string, graphFunc UgenGraphFunc) *synthdef {
	def := newsynthdef(name)
	// It would be nice to parse synthdef params from function arguments
	// with the reflect package, but see
	// https://groups.google.com/forum/#!topic/golang-nuts/nM_ZhL7fuGc
	// for discussion of the (im)possibility of getting function argument
	// names at runtime.
	// Since this is not possible, what we need to do is let users add
	// synthdef params anywhere in their UgenGraphFunc using the Params interface.
	// Then in order to correctly map the values passed when creating 
	// a synth node they have to be passed in the same order
	// they were created in the UgenGraphFunc.
	params := newParams()
	root := graphFunc(params)
	def.AddParams(params)
	flatten(root, params, def)
	return def
}

func flatten(node UgenNode, params Params, def *synthdef) *input {
	stack := NewStack()
	inputs := node.Inputs()
	// iterate through ugen inputs in reverse order
	for i := len(inputs)-1; i >= 0; i-- {
		in := inputs[i]
		if node, isNode := in.(UgenNode); isNode {
			stack.Push(flatten(node, params, def))
		} else {
			stack.Push(in)
		}
	}

	// add inputs to root
	var in *input
	u := cloneUgen(node)
	for val := stack.Pop(); val != nil; val = stack.Pop() {
		if intVal, isInt := val.(int); isInt {
			in = def.AddConstant(float32(intVal))
		} else if floatVal, isFloat32 := val.(float32); isFloat32 {
			in = def.AddConstant(float32(floatVal))
		} else if floatVal, isFloat64 := val.(float64); isFloat64 {
			in = def.AddConstant(float32(floatVal))
		} else if paramVal, isParam := val.(Param); isParam {
			in = &input{0, paramVal.Index()}
		} else if inputVal, isInput := val.(*input); isInput {
			in = inputVal
		} else {
			panic("ugen inputs must be constant, param, or ugens")
		}
		u.AppendInput(in)
	}

	return def.AddUgen(u)
}
