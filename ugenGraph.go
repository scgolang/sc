package sc

import (
	. "github.com/briansorahan/sc/types"
)

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
