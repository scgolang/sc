package sc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/awalterschulze/gographviz"
)

const (
	synthdefStart     = "SCgf"
	synthdefVersion   = 2
	constantUgenIndex = -1
)

var byteOrder = binary.BigEndian

// Synthdef defines the structure of synthdef data as defined
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

// Bytes writes a synthdef to a byte array
func (def *Synthdef) Bytes() ([]byte, error) {
	arr := []byte{}
	buf := bytes.NewBuffer(arr)
	err := def.Write(buf)
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
func (def *Synthdef) CompareToFile(path string) (bool, error) {
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
	err = def.Write(buf)
	if err != nil {
		return false, err
	}
	bufBytes := buf.Bytes()
	return compareBytes(bufBytes, fromDisk), nil
}

// CompareToDef compare this synthdef to another.
func (def *Synthdef) CompareToDef(other *Synthdef) (bool, error) {
	var err error
	buf1 := bytes.NewBuffer(make([]byte, 0))
	buf2 := bytes.NewBuffer(make([]byte, 0))
	err = def.Write(buf1)
	if err != nil {
		return false, err
	}
	err = other.Write(buf2)
	if err != nil {
		return false, err
	}
	return compareBytes(buf1.Bytes(), buf2.Bytes()), nil
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
func (def *Synthdef) WriteGraph(w io.Writer) error {
	graph := newGraph(def.Name)
	for i, ugen := range def.Ugens {
		ustr := fmt.Sprintf("%s_%d", ugen.Name, i)
		graph.AddNode(def.Name, ustr, nil)
		for j := len(ugen.Inputs) - 1; j >= 0; j-- {
			input := ugen.Inputs[j]
			if input.UgenIndex == -1 {
				c := def.Constants[input.OutputIndex]
				cstr := fmt.Sprintf("%f", c)
				graph.AddNode(ustr, cstr, constAttrs)
				graph.AddEdge(cstr, ustr, true, nil)
			} else {
				subgraph := def.addsub(input.UgenIndex, def.Ugens[input.UgenIndex])
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
func (def *Synthdef) addsub(idx int32, ugen *ugen) *gographviz.Graph {
	ustr := fmt.Sprintf("%s_%d", ugen.Name, idx)
	graph := newGraph(ustr)
	for j := len(ugen.Inputs) - 1; j >= 0; j-- {
		input := ugen.Inputs[j]
		if input.UgenIndex == -1 {
			c := def.Constants[input.OutputIndex]
			cstr := fmt.Sprintf("%f", c)
			graph.AddNode(ustr, cstr, constAttrs)
			graph.AddEdge(cstr, ustr, true, nil)
		} else {
			subgraph := def.addsub(input.UgenIndex, def.Ugens[input.UgenIndex])
			graph.AddSubGraph(ustr, subgraph.Name, nil)
			graph.AddEdge(subgraph.Name, ustr, true, nil)
		}
	}
	return graph
}

// flatten converts a ugen graph into a format more
// suitable for sending /d_recv
func (def *Synthdef) flatten(params Params) {
	def.addParams(params)
	// get a topologically sorted ugens list
	ugenNodes := def.topsort(def.root)

	for _, ugenNode := range ugenNodes {
		// add ugen to synthdef
		ugen, _, seen := def.addUgen(ugenNode)
		if seen {
			continue
		}
		// add inputs to synthdef and to ugen
		for _, input := range ugenNode.Inputs() {
			def.flattenInput(params, ugen, input)
		}
	}
}

// flattenInput flattens a ugen graph starting from
// a particular ugen's input
func (def *Synthdef) flattenInput(params Params, ugen *ugen, input Input) {
	switch v := input.(type) {
	case Ugen:
		_, idx, _ := def.addUgen(v)
		for outputIndex := range v.Outputs() {
			ugen.AppendInput(newInput(int32(idx), int32(outputIndex)))
		}
	case C:
		idx := def.addConstant(v)
		ugen.AppendInput(newInput(-1, int32(idx)))
	case *param:
		idx := v.Index()
		ugen.AppendInput(newInput(0, idx))
	case MultiInput:
		mins := v.InputArray()
		for _, min := range mins {
			switch x := min.(type) {
			case Ugen:
				_, idx, _ := def.addUgen(x)
				// will we ever need to use a different output index? [bps]
				for outputIndex := range x.Outputs() {
					ugen.AppendInput(newInput(int32(idx), int32(outputIndex)))
				}
			case C:
				idx := def.addConstant(x)
				ugen.AppendInput(newInput(-1, int32(idx)))
			case *param:
				idx := x.Index()
				ugen.AppendInput(newInput(0, idx))
			}
		}
	}
}

// topsort performs a depth-first-search of a ugen tree
func (def *Synthdef) topsort(root Ugen) []Ugen {
	stack := newStack()
	def.topsortr(root, stack, 0)
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
func (def *Synthdef) topsortr(root Ugen, stack *stack, depth int) {
	stack.Push(root)
	inputs := root.Inputs()
	numInputs := len(inputs)
	for i := numInputs - 1; i >= 0; i-- {
		def.processUgenInput(inputs[i], stack, depth)
	}
}

// processUgenInput processes a single ugen input
func (def *Synthdef) processUgenInput(input Input, stack *stack, depth int) {
	switch v := input.(type) {
	case Ugen:
		def.topsortr(v, stack, depth+1)
		break
	case MultiInput:
		// multi input
		mins := v.InputArray()
		for j := len(mins) - 1; j >= 0; j-- {
			min := mins[j]
			switch w := min.(type) {
			case Ugen:
				def.topsortr(w, stack, depth+1)
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
func (def *Synthdef) addParams(p Params) {
	paramList := p.List()
	numParams := len(paramList)
	def.InitialParamValues = make([]float32, numParams)
	def.ParamNames = make([]ParamName, numParams)
	for i, param := range paramList {
		def.InitialParamValues[i] = param.InitialValue()
		def.ParamNames[i] = ParamName{param.Name(), param.Index()}
	}
	if numParams > 0 {
		ctl := p.Control()
		def.seen = append(def.seen, ctl)
		// create a list with the single Control ugen,
		// then append any existing ugens in the synthdef
		// to that list
		control := []*ugen{cloneUgen(ctl)}
		def.Ugens = append(control, def.Ugens...)
	}
}

// addUgen adds a Ugen to a synthdef and returns
// the ugen that was added, the position in the ugens array, and
// a flag indicating whether this is a new ugen or one that
// has already been visited
func (def *Synthdef) addUgen(u Ugen) (*ugen, int, bool) {
	for i, un := range def.seen {
		if un == u {
			return def.Ugens[i], i, true
		}
	}
	def.seen = append(def.seen, u)
	idx := len(def.Ugens)
	ugen := cloneUgen(u)
	def.Ugens = append(def.Ugens, ugen)
	return ugen, idx, false
}

// addConstant adds a constant to a synthdef and returns
// the index in the constants array where that constant is
// located.
// It ensures that constants are not added twice by returning the
// position in the constants array of the existing constant if
// you try to add a duplicate.
func (def *Synthdef) addConstant(c C) int {
	for i, f := range def.Constants {
		if f == float32(c) {
			return i
		}
	}
	l := len(def.Constants)
	def.Constants = append(def.Constants, float32(c))
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
