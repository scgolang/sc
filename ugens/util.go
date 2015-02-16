package ugens

import (
	"github.com/briansorahan/sc"
)

// getConstant attempts to get a float32 value from
// an interface{}
// it first tries to type assert with int, then float64
// if both of these fail it will return false as the second
// value
func getConstant(arg interface{}) (float32, bool) {
	if iv, isInt := arg.(int); isInt {
		return float32(iv), true
	}
	if fv, isFloat := arg.(float64); isFloat {
		return float32(fv), true
	}
	return 0, false
}

// getInput either returns a constant input or a ugen input
// by running some type assertions on the provided arg
// if the arg is neither of these, then it returns nil
func getInput(arg interface{}) sc.Input {
	if cv, isConstant := getConstant(arg); isConstant {
		return constantInput(cv)
	}
	if nv, isNode := arg.(*baseNode); isNode {
		return nv
	}
	return nil
}

func parseArgs(node *baseNode, args ...interface{}) {
}
