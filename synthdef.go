package sc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

const (
	synthdefStart     = "SCgf"
	synthdefVersion   = 2
	constantUgenIndex = -1
)

var byteOrder = binary.BigEndian

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

	// Ugens is the list of ugens that appear in the synth def.
	// The root of the ugen graph will always be last.
	Ugens []*ugen `json:"ugens" xml:"Ugens>Ugen"`

	// Variants is the list of variants contained in the synth def
	Variants []*Variant `json:"variants" xml:"Variants>Variant"`

	// seen is an array of ugen nodes that have been added
	// to the synthdef
	seen []Ugen

	// root is the root of the ugen tree that defines this synthdef
	// this is used, for example, when drawing an svg representation
	// of the synthdef
	root Ugen
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

func (self *Synthdef) WriteXML(w io.Writer) error {
	enc := xml.NewEncoder(w)
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

// compareBytes returns true if two byte arrays
// are identical, false if they are not
func compareBytes(a, b []byte) bool {
	la, lb := len(a), len(b)
	if la != lb {
		return false
	}
	for i, octet := range a {
		if octet != b[i] {
			return false
		}
	}
	return true
}

// CompareToFile compares this synthdef to another one stored on disk.
func (self *Synthdef) CompareToFile(path string) (bool, error) {
	f, err := os.Open(path)
	defer f.Close()
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
	bufBytes := buf.Bytes()
	return compareBytes(bufBytes, fromDisk), nil
}

// CompareToDef compare this synthdef to another.
func (self *Synthdef) CompareToDef(def *Synthdef) (bool, error) {
	var err error
	buf1 := bytes.NewBuffer(make([]byte, 0))
	buf2 := bytes.NewBuffer(make([]byte, 0))
	err = self.Write(buf1)
	if err != nil {
		return false, err
	}
	err = def.Write(buf2)
	if err != nil {
		return false, err
	}
	return compareBytes(buf1.Bytes(), buf2.Bytes()), nil
}

// WriteDiff provides a diagnostic report indicating any differences
// between this synthdef and another.
func (self *Synthdef) WriteDiff(w io.Writer, def *Synthdef) (bool, error) {
	bytewise, err := self.CompareToDef(def)
	if err != nil {
		return false, err
	}
	// early out if they are bytewise identical
	if bytewise {
		return true, nil
	}

	same := true
	// compare number of ugens
	nuself := len(self.Ugens)
	nudef := len(def.Ugens)
	if nudef != nuself {
		_, err = fmt.Fprintf(w, "different number of ugens: self=%d def=%d\n", nuself, nudef)
		if err != nil {
			return false, err
		}
		same = false
	}

	// compare number of constants
	ncdef := len(def.Constants)
	ncself := len(self.Constants)
	if ncdef != ncself {
		_, err = fmt.Fprintf(w, "different number of constants: self=%d def=%d\n", ncself, ncdef)
		if err != nil {
			return false, err
		}
		same = false
	}

	// compare number of InitialParamValues and ParamNames
	nipvdef := len(def.InitialParamValues)
	nipvself := len(self.InitialParamValues)
	if nipvdef != nipvself {
		_, err = fmt.Fprintf(w, "different number of initial param values: self=%d def=%d\n", nipvself, nipvdef)
		if err != nil {
			return false, err
		}
		same = false
	}
	npndef := len(def.ParamNames)
	npnself := len(self.ParamNames)
	if npndef != npnself {
		_, err = fmt.Fprintf(w, "different number of param names self=%d def=%d\n", npnself, npndef)
		if err != nil {
			return false, err
		}
		same = false
	}

	if !same {
		return false, nil
	}
	return self.writeUgensDiff(w, def, def.Ugens[nudef-1], self.Ugens[nudef-1])
}

// writeUgensDiff writes any differences in a ugen graph to an io.Writer
// and returns a bool indicating if the ugen graphs are the same, and an error
// for any errors that occured
func (self *Synthdef) writeUgensDiff(w io.Writer, def *Synthdef, uself, udef *ugen) (bool, error) {
	var err error
	same := true
	// Name
	if udef.Name != uself.Name {
		_, err = fmt.Fprintf(w, "Name is different self=%s def=%s\n", uself.Name, udef.Name)
		if err != nil {
			return false, err
		}
		same = false
	}
	// Rate
	if udef.Rate != uself.Rate {
		_, err = fmt.Fprintf(w, "Rate is different self=%d def=%d)", uself.Rate, udef.Rate)
		if err != nil {
			return false, err
		}
		same = false
	}
	// SpecialIndex
	if udef.SpecialIndex != uself.SpecialIndex {
		_, err = fmt.Fprintf(w, "SpecialIndex is different self=%d def=%d\n", udef.SpecialIndex, uself.SpecialIndex)
		if err != nil {
			return false, err
		}
		same = false
	}
	// number of Inputs
	nidef := len(udef.Inputs)
	niself := len(uself.Inputs)
	if nidef != niself {
		_, err = fmt.Fprintf(w, "different number of inputs self=%d def=%d\n", niself, nidef)
		if err != nil {
			return false, err
		}
		same = false
	}
	// number of Outputs
	nodef := len(udef.Outputs)
	noself := len(uself.Outputs)
	if noself != nodef {
		_, err = fmt.Fprintf(w, "different number of outputs self=%d def=%d\n", noself, nodef)
		if err != nil {
			return false, err
		}
		same = false
	}
	// ugen inputs
	var iself input
	for i, idef := range udef.Inputs {
		iself = uself.Inputs[i]
		if idef.UgenIndex == -1 {
			if iself.UgenIndex != -1 {
				// def's input is a constant, but self's is not
				_, err = fmt.Fprintf(w, "input %d def is constant but self is not\n", i)
				if err != nil {
					return false, err
				}
				same = false
				continue
			}
			if int32(len(def.Constants)-1) < idef.OutputIndex {
				return false, fmt.Errorf("output index %d is too large", idef.OutputIndex)
			}
			if int32(len(self.Constants)-1) < iself.OutputIndex {
				return false, fmt.Errorf("output index %d is too large", iself.OutputIndex)
			}
			cdef := def.Constants[idef.OutputIndex]
			cself := self.Constants[iself.OutputIndex]
			if cdef != cself {
				_, err = fmt.Fprintf(w, "constants have different values self=%f def=%f\n", cself, cdef)
				if err != nil {
					return false, err
				}
				same = false
			}
		} else {
			// ugen input
			if idef.OutputIndex != iself.OutputIndex {
				_, err = fmt.Fprintf(w, "input %d different output indices self=%d def=%d\n", iself.OutputIndex, idef.OutputIndex)
				if err != nil {
					return false, err
				}
				same = false
			}
			same, err = self.writeUgensDiff(w, def, def.Ugens[idef.UgenIndex], self.Ugens[iself.UgenIndex])
			if err != nil {
				return false, err
			}
		}
	}
	return same, nil
}

// Compare compares this synthdef byte-for-byte with
// the synthdef sclang generates using the given string.
// Note that using this method requires you to have sclang installed.
// Returns true if the synthdefs are the same, false otherwise.
func (self *Synthdef) Compare(def string) (bool, error) {
	tmp := os.TempDir()
	scSyndef := path.Join(tmp, fmt.Sprintf("%s.scsyndef", self.Name))
	const wrap = `SynthDef(\%s, %s).writeDefFile("%s"); 0.exit;`
	contents := fmt.Sprintf(wrap, self.Name, def, tmp)
	f, err := ioutil.TempFile(tmp, "sclang_")
	if err != nil {
		return false, err
	}
	written, err := f.Write([]byte(contents))
	if err != nil {
		return false, err
	}
	if written != len(contents) {
		return false, fmt.Errorf("only wrote %d bytes", written)
	}
	// generate the .scsyndef file
	cmd := exec.Command("sclang", f.Name())
	err = cmd.Run()
	if err != nil {
		return false, err
	}
	// read it and compare to this synthdef
	g, err := os.Open(scSyndef)
	if err != nil {
		return false, err
	}
	fromDisk, err := ioutil.ReadAll(g)
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
	// TODO: use newsynthdef here
	synthDef := Synthdef{
		defName.String(),
		constants,
		initialValues,
		paramNames,
		ugens,
		variants,
		make([]Ugen, 0),
		nil,
	}
	return &synthDef, nil
}

func newGraph(name string) *gographviz.Graph {
	g := gographviz.NewGraph()
	g.SetName(name)
	g.SetDir(true)
	g.AddAttr(name, "rankdir", "BT")
	// g.AddAttr(name, "ordering", "in")
	return g
}

var constAttrs = map[string]string{"shape": "record"}

// WriteGraph writes a dot-formatted representation of
// a synthdef's ugen graph to an io.Writer. See
// http://www.graphviz.org/content/dot-language.
func (self *Synthdef) WriteGraph(w io.Writer) error {
	graph := newGraph(self.Name)
	for i, ugen := range self.Ugens {
		ustr := fmt.Sprintf("%s_%d", ugen.Name, i)
		graph.AddNode(self.Name, ustr, nil)
		for j := len(ugen.Inputs) - 1; j >= 0; j-- {
			input := ugen.Inputs[j]
			if input.UgenIndex == -1 {
				c := self.Constants[input.OutputIndex]
				cstr := fmt.Sprintf("%f", c)
				graph.AddNode(ustr, cstr, constAttrs)
				graph.AddEdge(cstr, ustr, true, nil)
			} else {
				subgraph := self.addsub(input.UgenIndex, self.Ugens[input.UgenIndex])
				graph.AddSubGraph(ustr, subgraph.Name, nil)
				graph.AddEdge(subgraph.Name, ustr, true, nil)
			}
		}
	}
	gstr := graph.String()
	_, writeErr := w.Write(bytes.NewBufferString(gstr).Bytes())
	return writeErr
}

// addsub creates a subgraph rooted at a particular ugen
func (self *Synthdef) addsub(idx int32, ugen *ugen) *gographviz.Graph {
	ustr := fmt.Sprintf("%s_%d", ugen.Name, idx)
	graph := newGraph(ustr)
	for j := len(ugen.Inputs) - 1; j >= 0; j-- {
		input := ugen.Inputs[j]
		if input.UgenIndex == -1 {
			c := self.Constants[input.OutputIndex]
			cstr := fmt.Sprintf("%f", c)
			graph.AddNode(ustr, cstr, constAttrs)
			graph.AddEdge(cstr, ustr, true, nil)
		} else {
			subgraph := self.addsub(input.UgenIndex, self.Ugens[input.UgenIndex])
			graph.AddSubGraph(ustr, subgraph.Name, nil)
			graph.AddEdge(subgraph.Name, ustr, true, nil)
		}
	}
	return graph
}

// flatten
func (self *Synthdef) flatten(params Params) {
	self.addParams(params)
	// get a topologically sorted ugens list
	ugenNodes := self.topsort(self.root)

	for _, ugenNode := range ugenNodes {
		// add ugen to synthdef
		ugen, _, seen := self.addUgen(ugenNode)
		if seen {
			continue
		}
		// add inputs to synthdef and to ugen
		inputs := ugenNode.Inputs()

		for _, input := range inputs {
			switch v := input.(type) {
			case Ugen:
				_, idx, _ := self.addUgen(v)
				for outputIndex := range v.Outputs() {
					ugen.AppendInput(newInput(int32(idx), int32(outputIndex)))
				}
			case C:
				idx := self.addConstant(v)
				ugen.AppendInput(newInput(-1, int32(idx)))
			case *param:
				idx := v.Index()
				ugen.AppendInput(newInput(0, idx))
			case MultiInput:
				mins := v.InputArray()
				for _, min := range mins {
					switch x := min.(type) {
					case Ugen:
						_, idx, _ := self.addUgen(x)
						// will we ever need to use a different output index? [bps]
						for outputIndex := range x.Outputs() {
							ugen.AppendInput(newInput(int32(idx), int32(outputIndex)))
						}
					case C:
						idx := self.addConstant(x)
						ugen.AppendInput(newInput(-1, int32(idx)))
					case *param:
						idx := x.Index()
						ugen.AppendInput(newInput(0, idx))
					}
				}
			}
		}
	}
}

// topsort performs a depth-first-search of a ugen tree
func (self *Synthdef) topsort(root Ugen) []Ugen {
	stack := newStack()
	self.topsortr(root, stack, 0)
	n := stack.Size()
	ugens := make([]Ugen, n)
	i := 0
	for v := stack.Pop(); v != nil; v = stack.Pop() {
		ugens[i] = v.(Ugen)
		i = i + 1
	}
	return ugens
}

// topsortr performs a depth-first-search of a ugen tree
// starting at a given depth
func (self *Synthdef) topsortr(root Ugen, stack *stack, depth int) {
	stack.Push(root)
	inputs := root.Inputs()
	numInputs := len(inputs)
	for i := numInputs - 1; i >= 0; i-- {
		self.processUgenInput(inputs[i], stack, depth)
	}
}

// processUgenInput processes a single ugen input
func (self *Synthdef) processUgenInput(input Input, stack *stack, depth int) {
	switch v := input.(type) {
	case Ugen:
		self.topsortr(v, stack, depth+1)
		break
	case MultiInput:
		// multi input
		mins := v.InputArray()
		for j := len(mins) - 1; j >= 0; j-- {
			min := mins[j]
			switch w := min.(type) {
			case Ugen:
				self.topsortr(w, stack, depth+1)
				break
			}
		}
		break
	}
}

// addParams will do nothing if there are no synthdef params.
// If there are synthdef params it will
// (1) Add their default values to InitialParamValues
// (2) Add their names/indices to ParamNames
// (3) Add a Control ugen as the first ugen
func (self *Synthdef) addParams(p Params) {
	paramList := p.List()
	numParams := len(paramList)
	self.InitialParamValues = make([]float32, numParams)
	self.ParamNames = make([]ParamName, numParams)
	for i, param := range paramList {
		self.InitialParamValues[i] = param.InitialValue()
		self.ParamNames[i] = ParamName{param.Name(), param.Index()}
	}
	if numParams > 0 {
		ctl := p.Control()
		self.seen = append(self.seen, ctl)
		// create a list with the single Control ugen,
		// then append any existing ugens in the synthdef
		// to that list
		control := []*ugen{cloneUgen(ctl)}
		self.Ugens = append(control, self.Ugens...)
	}
}

// addUgen adds a Ugen to a synthdef and returns
// the ugen that was added, the position in the ugens array, and
// a flag indicating whether this is a new ugen or one that
// has already been visited
func (self *Synthdef) addUgen(u Ugen) (*ugen, int, bool) {
	for i, un := range self.seen {
		if un == u {
			return self.Ugens[i], i, true
		}
	}
	self.seen = append(self.seen, u)
	idx := len(self.Ugens)
	ugen := cloneUgen(u)
	self.Ugens = append(self.Ugens, ugen)
	return ugen, idx, false
}

// addConstant adds a constant to a synthdef and returns
// the index in the constants array where that constant is
// located.
// It ensures that constants are not added twice by returning the
// position in the constants array of the existing constant if
// you try to add a duplicate.
func (self *Synthdef) addConstant(c C) int {
	for i, f := range self.Constants {
		if f == float32(c) {
			return i
		}
	}
	l := len(self.Constants)
	self.Constants = append(self.Constants, float32(c))
	return l
}

func newsynthdef(name string, root Ugen) *Synthdef {
	def := Synthdef{
		name,
		make([]float32, 0),
		make([]float32, 0),
		make([]ParamName, 0),
		make([]*ugen, 0),
		make([]*Variant, 0),
		make([]Ugen, 0), // seen
		root,
	}
	return &def
}

// NewSynthdef creates a synthdef by traversing a ugen graph
func NewSynthdef(name string, graphFunc UgenFunc) *Synthdef {
	// It would be nice to parse synthdef params from function arguments
	// with the reflect package, but see
	// https://groups.google.com/forum/#!topic/golang-nuts/nM_ZhL7fuGc
	// for discussion of the (im)possibility of getting function argument
	// names at runtime.
	// Since this is not possible, what we need to do is let users add
	// synthdef params anywhere in their UgenFunc using Params.
	// Then in order to correctly map the values passed when creating
	// a synth node they have to be passed in the same order
	// they were created in the UgenFunc.
	params := newParams()
	root := graphFunc(params)
	def := newsynthdef(name, root)
	def.flatten(params)
	return def
}
