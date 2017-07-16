package sc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
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
	Constants []float32 `json:"constants,omitempty" xml:"Constants>Constant"`

	// InitialParamValues is an array of initial values for synth params
	InitialParamValues []float32 `json:"initialParamValues,omitempty" xml:"InitialParamValues>initialParamValue"`

	// ParamNames contains the names of the synth parameters
	ParamNames []ParamName `json:"paramNames,omitempty" xml:"ParamNames>ParamName"`

	// Ugens is the list of ugens that appear in the synth def.
	// The root of the ugen graph will always be last.
	Ugens []*Ugen `json:"ugens,omitempty" xml:"Ugens>Ugen"`

	// Variants is the list of variants contained in the synth def
	Variants []*Variant `json:"variants,omitempty" xml:"Variants>Variant"`

	// inidx helps us track which outputs we have consumed from the In ugen.
	// In triggers multichannel expansion when it has > 1 outputs,
	// but usually the behavior for ugens with multiple outputs that act
	// as an input to another ugen is that the outputs get appended as
	// consecutive inputs.
	// In, on the other hand, puts one of its output channels on each channel of the
	// expression tree above it.
	inidx int

	// seen is an array of ugen nodes that have been added
	// to the synthdef
	seen []*Ugen

	// root is the root of the ugen tree that defines this synthdef
	// this is used, for example, when drawing an svg representation
	// of the synthdef
	root Ugen
}

// NewSynthdef creates a synthdef by traversing a ugen graph
func NewSynthdef(name string, graphFunc UgenFunc) *Synthdef {
	// It would be nice to parse synthdef params from function arguments
	// with the reflect package.
	// See https://groups.google.com/forum/#!topic/golang-nuts/nM_ZhL7fuGc
	// for discussion of the (im)possibility of getting function argument
	// names at runtime.
	// Since this is not possible, what we need to do is let users add
	// synthdef params anywhere in their UgenFunc using Params.
	// Then in order to correctly map the values passed when creating
	// a synth node they have to be passed in the same order
	// they were created in the UgenFunc.
	var (
		params = newParams()
		def    = &Synthdef{Name: name, root: graphFunc(params)}
	)
	return def.flatten(params)
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

// CompareToFile compares this synthdef to another one stored on disk.
func (def *Synthdef) CompareToFile(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	fromDisk, err := ioutil.ReadAll(f)
	_ = f.Close() // Best effort.
	if err != nil {
		return false, err
	}
	buf := &bytes.Buffer{}
	if err := def.Write(buf); err != nil {
		return false, err
	}
	return compareBytes(buf.Bytes(), fromDisk), nil
}

// CompareToDef compare this synthdef to another.
func (def *Synthdef) CompareToDef(other *Synthdef) (bool, error) {
	var (
		buf1 = &bytes.Buffer{}
		buf2 = &bytes.Buffer{}
	)
	if err := def.Write(buf1); err != nil {
		return false, err
	}
	if err := other.Write(buf2); err != nil {
		return false, err
	}
	return compareBytes(buf1.Bytes(), buf2.Bytes()), nil
}

// Diff returns a diff of one synthdef and another.
// A diff is represented as a slice of pairs of strings.
// The first string in each pair describes the synthdef on the left (the receiver),
// and the second string describes the other synthdef.
func (def *Synthdef) Diff(other *Synthdef) [][2]string {
	return differ{def, other}.do()
}

// Root returns the root node in the synthdef's ugen graph.
func (def *Synthdef) Root() int32 {
	parents := make([]int, len(def.Ugens)) // Number of parents per ugen.

	for _, u := range def.Ugens {
		for _, in := range u.Inputs {
			if in.IsConstant() {
				continue
			}
			parents[in.UgenIndex]++
		}
	}
	for i, count := range parents {
		if count == 0 {
			return int32(i)
		}
	}
	return 0
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
		control := []*Ugen{cloneUgen(ctl)}
		def.Ugens = append(control, def.Ugens...)
	}
}

// addUgen adds a Ugen to a synthdef and returns
// the ugen that was added, the position in the ugens array, and
// a flag indicating whether this is a new ugen or one that
// has already been visited.
func (def *Synthdef) addUgen(u *Ugen) (*Ugen, int, bool) {
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

// flatten converts a ugen graph into a format more suitable for sending /d_recv
func (def *Synthdef) flatten(params Params) *Synthdef {
	def.addParams(params)

	// Get a topologically sorted ugens list.
	ugenNodes := def.topsort(&def.root)

	for _, ugenNode := range ugenNodes {
		// add ugen to synthdef
		ugen, _, seen := def.addUgen(ugenNode)
		if seen {
			continue
		}
		// Add inputs to synthdef and to ugen.
		for _, input := range ugenNode.inputs {
			def.flattenInput(params, ugen, input)
		}
	}
	return def
}

// flattenInput flattens a ugen graph starting from a particular ugen's input.
func (def *Synthdef) flattenInput(params Params, ugen *Ugen, input Input) {
	switch v := input.(type) {
	case *Ugen:
		_, idx, _ := def.addUgen(v)

		// In has different behavior than other ugens when it is multichannel expanded.
		// Each channel of its output gets mapped to a single channel of the expression
		// tree above it. [briansorahan]
		if v.Name == "In" {
			ugen.Inputs = append(ugen.Inputs, UgenInput{
				UgenIndex:   int32(idx),
				OutputIndex: int32(def.inidx),
			})
			def.inidx++
			break
		}
		for outputIndex := range v.Outputs {
			ugen.Inputs = append(ugen.Inputs, UgenInput{
				UgenIndex:   int32(idx),
				OutputIndex: int32(outputIndex),
			})
		}
	case C:
		idx := def.addConstant(v)
		ugen.Inputs = append(ugen.Inputs, UgenInput{
			UgenIndex:   -1,
			OutputIndex: int32(idx),
		})
	case *param:
		idx := v.Index()
		ugen.Inputs = append(ugen.Inputs, UgenInput{
			UgenIndex:   0,
			OutputIndex: idx,
		})
	case MultiInput:
		for _, min := range v.InputArray() {
			switch x := min.(type) {
			case *Ugen:
				_, idx, _ := def.addUgen(x)
				// will we ever need to use a different output index? [bps]
				for outputIndex := range x.Outputs {
					ugen.Inputs = append(ugen.Inputs, UgenInput{
						UgenIndex:   int32(idx),
						OutputIndex: int32(outputIndex),
					})
				}
			case C:
				ugen.Inputs = append(ugen.Inputs, UgenInput{
					UgenIndex:   -1,
					OutputIndex: int32(def.addConstant(x)),
				})
			case *param:
				ugen.Inputs = append(ugen.Inputs, UgenInput{
					UgenIndex:   0,
					OutputIndex: x.Index(),
				})
			}
		}
	}
}

// processUgenInput processes a single ugen input
func (def *Synthdef) processUgenInput(input Input, stack *stack, depth int) {
	switch v := input.(type) {
	case *Ugen:
		def.topsortr(v, stack, depth+1)
		break
	case MultiInput:
		// multi input
		mins := v.InputArray()
		for j := len(mins) - 1; j >= 0; j-- {
			switch w := mins[j].(type) {
			case *Ugen:
				def.topsortr(w, stack, depth+1)
				break
			}
		}
		break
	}
}

// topsort performs a depth-first-search of a ugen tree
func (def *Synthdef) topsort(root *Ugen) []*Ugen {
	stack := newStack()

	def.topsortr(root, stack, 0)

	var (
		i     = 0
		n     = stack.Size()
		ugens = make([]*Ugen, n)
	)
	for v := stack.Pop(); v != nil; v = stack.Pop() {
		ugens[i] = v.(*Ugen)
		i = i + 1
	}
	return ugens
}

// topsortr performs a depth-first-search of a ugen tree starting at a given depth.
func (def *Synthdef) topsortr(root *Ugen, stack *stack, depth int) {
	stack.Push(root)

	var (
		inputs    = root.inputs
		numInputs = len(inputs)
	)
	for i := numInputs - 1; i >= 0; i-- {
		def.processUgenInput(inputs[i], stack, depth)
	}
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

// differ creates a diff of two synthdefs.
type differ [2]*Synthdef

// crawl crawls the ugen graph, starting at the specified ugens in each synthdef,
// looking for structural differences.
func (d differ) crawl(diffs [][2]string, idx1, idx2 int32) [][2]string {
	var (
		u1 = d[0].Ugens[idx1]
		u2 = d[1].Ugens[idx2]
	)
	if u1.Name != u2.Name {
		return append(diffs, [2]string{
			fmt.Sprintf("ugen %d is a(n) %s", idx1, u1.Name),
			fmt.Sprintf("ugen %d is a(n) %s", idx2, u2.Name),
		})
	}
	if l1, l2 := len(u1.Inputs), len(u2.Inputs); l1 != l2 {
		return append(diffs, [2]string{
			fmt.Sprintf("%s has %d inputs", u1.Name, l1),
			fmt.Sprintf("%s has %d inputs", u2.Name, l2),
		})
	}
	// They have the same number of inputs.
	for i, in1 := range u1.Inputs {
		var (
			in2 = u2.Inputs[i]
			oi1 = in1.OutputIndex
			oi2 = in2.OutputIndex
			ui1 = in1.UgenIndex
			ui2 = in2.UgenIndex
		)
		if in1.IsConstant() && in2.IsConstant() {
			if v1, v2 := d[0].Constants[oi1], d[1].Constants[oi2]; v1 != v2 {
				diffs = append(diffs, [2]string{
					fmt.Sprintf("%s (ugen %d), input %d has constant value %f", u1.Name, idx1, i, v1),
					fmt.Sprintf("%s (ugen %d), input %d has constant value %f", u2.Name, idx2, i, v2),
				})
			}
			continue
		}
		if in1.IsConstant() && !in2.IsConstant() {
			diffs = append(diffs, [2]string{
				fmt.Sprintf("%s(%d), input %d is constant (%f)", u1.Name, idx1, i, d[0].Constants[oi1]),
				fmt.Sprintf("%s(%d), input %d points to %s(%d)", u2.Name, idx2, i, d[1].Ugens[ui2].Name, ui2),
			})
			continue
		}
		if !in1.IsConstant() && in2.IsConstant() {
			diffs = append(diffs, [2]string{
				fmt.Sprintf("%s(%d), input %d points to %s(%d)", u1.Name, idx1, i, d[0].Ugens[ui1].Name, ui1),
				fmt.Sprintf("%s(%d), input %d is constant (%f)", u2.Name, idx2, i, d[1].Constants[oi2]),
			})
			continue
		}
		// They are both not constant.
		// TODO: detect cycles
		diffs = append(diffs, d.crawl(diffs, ui1, ui2)...)
	}
	return diffs
}

// do performs a diff.
// The diff shows whether one ugen graph differs structurally from another.
// If the returned slice is empty it means the synthdefs are structurally identical.
func (d differ) do() [][2]string {
	// Early out if they have different numbers of ugens or constants.
	if l1, l2 := len(d[0].Ugens), len(d[1].Ugens); l1 != l2 {
		return [][2]string{
			{
				fmt.Sprintf("%d ugens", l1),
				fmt.Sprintf("%d ugens", l2),
			},
		}
	}
	if l1, l2 := len(d[0].Constants), len(d[1].Constants); l1 != l2 {
		return [][2]string{
			{
				fmt.Sprintf("%d constants", l1),
				fmt.Sprintf("%d constants", l2),
			},
		}
	}
	return d.crawl([][2]string{}, d[0].Root(), d[1].Root())
}

var commutative = map[string]struct{}{
	"BinaryOpUgen": struct{}{},
}
