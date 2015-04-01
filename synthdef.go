package sc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	. "github.com/briansorahan/sc/types"
	. "github.com/briansorahan/sc/ugens"
	"io"
	"io/ioutil"
	"os"
)

const (
	synthdefStart   = "SCgf"
	synthdefVersion = 2
)

var byteOrder = binary.BigEndian

// ugenInput is a helper type used when parsing
// a ugen graph. When we are parsing a ugen graph
// we maintain a stack of ugenInput's that will
// all get added to the synthdef after the whole
// graph has been traversed.
type ugenInput struct {
	// ugen is the ugen that will get added to the synthdef
	ugen  *ugen
	// node is the UgenNode visited while walking the ugen graph.
	// To ensure that we don't add a given ugen twice, we compare
	// a UgenNode to the node members of all the ugenInput's on
	// the ugens stack
	node UgenNode
	// input is the input that ugens that have this ugen as an input
	// should add to their list of inputs
	input *input
}

// constantInput is a helper type used when parsing
// ugen graphs.
type constantInput struct {
	constant C
	input    *input
}

// synthdef defines the structure of synth def data as defined
// in http://doc.sccode.org/Reference/Synth-Definition-File-Format.html
type Synthdef struct {
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
	Variants []*Variant `json:"variants" xml:"Variants>Variant"`

	// ugens is a stack we use when parsing a ugen graph
	ugenStack []*ugenInput

	// constants is a stack we use when parsing a ugen graph
	constantStack []*constantInput
}

// Write writes a binary representation of a synthdef to an io.Writer.
// The binary representation written by this method is
// the data that scsynth expects at its /d_recv endpoint.
func (self *Synthdef) Write(w io.Writer) error {
	written, err := w.Write(bytes.NewBufferString(synthdefStart).Bytes())
	if written != len(synthdefStart) {
		return fmt.Errorf("Could not write synthdef")
	}
	if err != nil {
		return err
	}
	// write synthdef version
	err = binary.Write(w, byteOrder, int32(synthdefVersion))
	if err != nil {
		return err
	}
	// write number of synthdefs
	err = binary.Write(w, byteOrder, int16(1))
	if err != nil {
		return err
	}
	// write synthdef name
	name := newPstring(self.Name)
	err = name.Write(w)
	if err != nil {
		return err
	}
	// write number of constants
	err = binary.Write(w, byteOrder, int32(len(self.Constants)))
	if err != nil {
		return err
	}
	// write constant values
	for _, constant := range self.Constants {
		err = binary.Write(w, byteOrder, constant)
		if err != nil {
			return err
		}
	}
	// write number of params
	err = binary.Write(w, byteOrder, int32(len(self.ParamNames)))
	if err != nil {
		return err
	}
	// write initial param values
	// BUG(briansorahan) what happens in sclang when a ugen graph func
	//                   does not provide initial param values? do they
	//                   not appear in the synthdef? default to 0?
	for _, val := range self.InitialParamValues {
		err = binary.Write(w, byteOrder, val)
		if err != nil {
			return err
		}
	}
	// write number of param names
	err = binary.Write(w, byteOrder, int32(len(self.ParamNames)))
	if err != nil {
		return err
	}
	// write param names
	for _, p := range self.ParamNames {
		err = p.Write(w)
		if err != nil {
			return err
		}
	}
	// write number of ugens
	err = binary.Write(w, byteOrder, int32(len(self.Ugens)))
	if err != nil {
		return err
	}
	// write ugens
	for _, u := range self.Ugens {
		err = u.Write(w)
		if err != nil {
			return err
		}
	}
	// write number of variants
	err = binary.Write(w, byteOrder, int16(len(self.Variants)))
	if err != nil {
		return err
	}
	// write variants
	for _, v := range self.Variants {
		err = v.Write(w)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteJSON writes a json-formatted representation of a
// synthdef to an io.Writer
func (self *Synthdef) WriteJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(self)
}

// Bytes writes a synthdef to a byte array
func (self *Synthdef) Bytes() ([]byte, error) {
	arr := make([]byte, 0)
	buf := bytes.NewBuffer(arr)
	err := self.Write(buf)
	if err != nil {
		return arr, err
	}
	return buf.Bytes(), nil
}

// CompareToFile compares this synthdef to a synthdef
// stored on disk
func (self *Synthdef) CompareToFile(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	fromDisk, err := ioutil.ReadAll(f)
	if err != nil {
		return false, err
	}
	buf := bytes.NewBuffer(make([]byte, 0))
	err = self.Write(buf)
	if err != nil {
		return false, err
	}
	return compareBytes(buf.Bytes(), fromDisk), nil
}

// compareBytes returns true if two byte arrays
// are identical, false if they are not
func compareBytes(a, b []byte) bool {
	la, lb := len(a), len(b)
	if la != lb {
		fmt.Printf("different lengths a=%d b=%d\n", la, lb)
		return false
	}
	for i, octet := range a {
		if octet != b[i] {
			return false
		}
	}
	return true
}

// ReadSynthdef reads a synthdef from an io.Reader
func ReadSynthdef(r io.Reader) (*Synthdef, error) {
	// read the type
	startLen := len(synthdefStart)
	start := make([]byte, startLen)
	read, er := r.Read(start)
	if er != nil {
		return nil, er
	}
	if read != startLen {
		return nil, fmt.Errorf("Only read %d bytes of synthdef file", read)
	}
	actual := bytes.NewBuffer(start).String()
	if actual != synthdefStart {
		return nil, fmt.Errorf("synthdef started with %s instead of %s", actual, synthdefStart)
	}
	// read version
	var version int32
	er = binary.Read(r, byteOrder, &version)
	if er != nil {
		return nil, er
	}
	if version != synthdefVersion {
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
	variants := make([]*Variant, numVariants)
	for i := 0; int16(i) < numVariants; i++ {
		v, er := readVariant(r, numParams)
		if er != nil {
			return nil, er
		}
		variants[i] = v
	}
	// return a new synthdef
	synthDef := Synthdef{
		defName.String(),
		constants,
		initialValues,
		paramNames,
		ugens,
		variants,
		make([]*ugenInput, 0),
		make([]*constantInput, 0),
	}
	return &synthDef, nil
}

// addUgens adds all the ugens from ugenStack to
// the synthdef
func (self *Synthdef) addUgens() {
	for i, uin := 0, self.popUgenInput(); uin != nil; uin = self.popUgenInput() {
		uin.input.UgenIndex = int32(i)
		self.Ugens = append(self.Ugens, uin.ugen)
		i = i + 1
	}
}

// addConstants adds all the constants
func (self *Synthdef) addConstants() {
	for i, cin := 0, self.popConstantInput(); cin != nil; cin = self.popConstantInput() {
		cin.input.OutputIndex = int32(i)
		self.Constants = append(self.Constants, float32(cin.constant))
		i = i + 1
	}
}

// AddParams will do nothing if there are no synthdef params.
// If there are synthdef params it will
// (1) Add their default values to initialParamValues
// (2) Add their names/indices to paramNames
// (3) Add a Control ugen as the first ugen
func (self *Synthdef) AddParams(p *Params) {
	// HACK convert Params to an interface type
	paramList := p.List()
	numParams := len(paramList)
	self.InitialParamValues = make([]float32, numParams)
	self.ParamNames = make([]ParamName, numParams)
	for i, param := range paramList {
		self.InitialParamValues[i] = param.GetInitialValue()
		self.ParamNames[i] = ParamName{param.Name(), param.Index()}
	}
	if numParams > 0 {
		control := []*ugen{cloneUgen(p.Control())}
		self.Ugens = append(control, self.Ugens...)
	}
}

// pushUgenInput pushes a ugenInput onto ugenStack
func (self *Synthdef) pushUgenInput(uin *ugenInput) *ugenInput {
	for _, euin := range self.ugenStack {
		if uin.node == euin.node {
			return euin
		}
	}
	self.ugenStack = append(self.ugenStack, uin)
	return uin
}

// popUgenInput pops a ugenInput from ugenStack
func (self *Synthdef) popUgenInput() *ugenInput {
	l := len(self.ugenStack)
	if l == 0 {
		return nil
	}
	uin := self.ugenStack[l-1]
	self.ugenStack = self.ugenStack[0:l-1]
	return uin
}

// pushConstantInput pushes a constantInput onto constantStack
// returns the existing constantInput if the one you were
// trying to push has the same C value as one that is already
// on the stack
func (self *Synthdef) pushConstantInput(cin *constantInput) *constantInput {
	for _, ecin := range self.constantStack {
		if float32(ecin.constant) == float32(cin.constant) {
			return ecin
		}
	}
	self.constantStack = append(self.constantStack, cin)
	return cin
}

// popConstantInput pops a constantInput from constantStack
func (self *Synthdef) popConstantInput() *constantInput {
	l := len(self.constantStack)
	if l == 0 {
		return nil
	}
	cin := self.constantStack[l-1]
	self.constantStack = self.constantStack[0:l-1]
	return cin
}

// visitUgen visits a ugen node, processing all its inputs
// from last to first. since this method is not meant to visit
// ugens that are inputs to other ugens (see visitUgenInput).
// this method should only be used for visiting the root of a
// ugen graph
func (self *Synthdef) visitUgen(node UgenNode) {
	inputs := node.Inputs()
	inputStack := newStack()
	u := cloneUgen(node)
	uin := ugenInput{u, node, placeholderInput()}

	self.pushUgenInput(&uin)

	// iterate through ugen inputs in reverse order
	for i := len(inputs) - 1; i >= 0; i-- {
		i1 := inputs[i]

		if node, isNode := i1.(UgenNode); isNode {
			// flatten this ugen, return an input
			inputStack.Push(self.visitUgenInput(node))
		} else if multi, isMulti := i1.(MultiInput); isMulti {
			ins := multi.InputArray()

			for _, i2 := range ins {
				if ugen, isUgen := i2.(UgenNode); isUgen {
					inputStack.Push(self.visitUgenInput(ugen))
				} else if c, isc := i1.(C); isc {
					inputStack.Push(self.visitConstantInput(c))
				} else {
					panic("unrecognized input type")
				}
			}
		} else if c, isc := i1.(C); isc {
			inputStack.Push(self.visitConstantInput(c))
		} else {
			panic("unrecognized input type")
		}
	}

	// add the inputs
	for val := inputStack.Pop(); val != nil; val = inputStack.Pop() {
		switch in := val.(type) {
		case *ugenInput:
			u.AppendInput(in.input)
			break
		case *constantInput:
			u.AppendInput(in.input)
			break
		}
	}
}

// visitUgenInput visits a ugen that is being used as
// an input to another ugen
func (self *Synthdef) visitUgenInput(node UgenNode) *ugenInput {
	inputs := node.Inputs()
	inputStack := newStack()
	u := cloneUgen(node)

	uin := ugenInput{u, node, placeholderInput()}
	self.pushUgenInput(&uin)

	// iterate through ugen inputs in reverse order
	for i := len(inputs) - 1; i >= 0; i-- {
		i1 := inputs[i]

		if node, isNode := i1.(UgenNode); isNode {
			// flatten this ugen, return an input
			inputStack.Push(self.visitUgenInput(node))
		} else if multi, isMulti := i1.(MultiInput); isMulti {
			ins := multi.InputArray()

			for _, i2 := range ins {
				if ugen, isUgen := i2.(UgenNode); isUgen {
					inputStack.Push(self.visitUgenInput(ugen))
				} else if c, isc := i1.(C); isc {
					inputStack.Push(self.visitConstantInput(c))
				} else {
					panic("unrecognized input type")
				}
			}
		} else if c, isc := i1.(C); isc {
			inputStack.Push(self.visitConstantInput(c))
		} else {
			panic("unrecognized input type")
		}
	}

	for val := inputStack.Pop(); val != nil; val = inputStack.Pop() {
		switch in := val.(type) {
		case *ugenInput:
			u.AppendInput(in.input)
			break
		case *constantInput:
			u.AppendInput(in.input)
			break
		}
	}

	return &uin
}

// visitConstantInput visits a constant ugen input
func (self *Synthdef) visitConstantInput(c C) *constantInput {
	cin := &constantInput{c, &input{-1, 0}}
	return self.pushConstantInput(cin)
}

func (self *Synthdef) flatten(root UgenNode, params *Params) {
	self.AddParams(params)
	self.visitUgen(root)
	self.addUgens()
	self.addConstants()
}

func newsynthdef(name string) *Synthdef {
	def := Synthdef{
		name,
		make([]float32, 0),
		make([]float32, 0),
		make([]ParamName, 0),
		make([]*ugen, 0),
		make([]*Variant, 0),
		make([]*ugenInput, 0),     // ugens stack
		make([]*constantInput, 0), // constants stack
	}
	return &def
}

// NewSynthdef creates a synthdef by traversing a ugen graph
func NewSynthdef(name string, graphFunc UgenGraphFunc) *Synthdef {
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
	params := NewParams()
	root := graphFunc(params)
	def.flatten(root, params)
	return def
}

func placeholderInput() *input {
	return &input{-2, 0}
}

// Here we can see why it is not sufficient to simply
// scan the inputs at a given depth.
//
// If we had L and R stacks, what would they look like?
//
// L: [[BinaryOpUgen][WhiteNoise, 0.1]]
//
//                                 Out
//                                  |
//                              +-------+
//                              |       |
//                              0       AllpassC
//                                         |
//                          +--------+--------+--------+
//                          |        |        |        |
//               BinaryOpUgen      0.01     XLine     0.2
//                  |                         |
//              +--------+          +------+-------+-------+
//              |        |          |      |       |       |
//      WhiteNoise     SinOsc    0.0001   0.01   SinOsc    0
//                       |                         |
//                   +-------+                 +-------+
//                   |       |                 |       |
//                  0.6      0               0.02      0
//
// constants: [0.6, 0, 0.02, 0.0001, 0.01, 0.2]
//
// ugens: [SinOsc(0.6), WhiteNoise, BinaryOpUgen, SinOsc(0.02), XLine, AllpassC, Out]

// WhiteNoise => BinaryOpUgen
// SinOsc(0.6) => BinaryOpUgen
// BinaryOpUgen => AllpassC
// SinOsc(0.02) => XLine
// XLine => AllpassC
// AllpassC => Out
//
// Why does SinOsc(0.6) come before WhiteNoise above!?!?!?



// We need to add constants and ugens to the synthdef in depth-first order.
// If we let the root node be depth=0, then in the graph above we have
// depth=1 => 0 AllpassC
// depth=2 => BinaryOpUgen 0.01 XLine 0.2
// depth=3 => WhiteNoise 0.1 0.0001 0.01 20 0
//
// So the data structure we need for this could be [][]Input, but how
// do we keep track of the inputs for each ugen?
// This makes me think I should use [][]interface{} to store
// *ugenInput and constantInput*



// Ugen Graph Traversal Pseudocode
//
// VisitAt(0, root)
//
// func VisitAt(depth, node) {
//     for input in Reverse(node.Inputs()) {
//         if constant {
//     }
// }
