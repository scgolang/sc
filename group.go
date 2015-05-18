package sc

import (
	"fmt"
	"github.com/scgolang/osc"
	"reflect"
)

const (
	gQueryTree      = "/g_queryTree"
	gQueryTreeReply = "/g_queryTree.reply"
)

type node struct {
	id int32 `json:"id" xml:"id,attr"`
}

type group struct {
	node
	children []*node `json:"children" xml:children>child"`
}

// parseGroup parses information about a group from a message
// received at /g_queryTree
// it *does not* recursively query for child groups
func parseGroup(msg *osc.Message) (*group, error) {
	// return an error if msg.Address is not right
	if msg.Address != gQueryTreeReply {
		return nil, fmt.Errorf("msg.Address should be %s, got %s", gQueryTreeReply, msg.Address)
	}
	// g_queryTree replies should have at least 3 arguments
	g, numArgs := new(group), msg.CountArguments()
	if numArgs < 3 {
		return nil, fmt.Errorf("expected 3 arguments for message, got %d", numArgs)
	}
	// get the id of the group this reply is for
	var isint bool
	g.id, isint = msg.Arguments[1].(int32)
	if !isint {
		v := msg.Arguments[1]
		t := reflect.TypeOf(v)
		return nil, fmt.Errorf("expected arg 1 to be int32, got %s (%v)", t, v)
	}
	// initialize the children array
	var numChildren int32
	numChildren, isint = msg.Arguments[2].(int32)
	if !isint {
		v := msg.Arguments[1]
		t := reflect.TypeOf(v)
		return nil, fmt.Errorf("expected arg 2 to be int32, got %s (%v)", t, v)
	}
	if numChildren <= 0 {
		return nil, fmt.Errorf("expected numChildren >= 0, got %d", numChildren)
	}
	g.children = make([]*node, numChildren)
	// get the childrens' ids
	var nodeID, numControls int32
	for i := 3; i < numArgs; {
		nodeID, isint = msg.Arguments[i].(int32)
		if !isint {
			v := msg.Arguments[i]
			t := reflect.TypeOf(v)
			return nil, fmt.Errorf("expected arg %d (nodeID) to be int32, got %s (%v)", i, t, v)
		}
		g.children[i] = &node{nodeID}
		i += 3
		numControls, isint = msg.Arguments[i].(int32)
		if !isint {
			v := msg.Arguments[i]
			t := reflect.TypeOf(v)
			return nil, fmt.Errorf("expected arg %d (numControls) to be int32, got %s (%v)", i, t, v)
		}
		i += 1 + int(numControls*2)
	}
	return g, nil
}
