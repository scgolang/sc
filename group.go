package sc

import (
	"encoding/json"
	"fmt"
	"github.com/scgolang/osc"
	"io"
	"reflect"
)

const (
	gQueryTree      = "/g_queryTree"
	gQueryTreeReply = "/g_queryTree.reply"
)

type node struct {
	Id int32 `json:"id" xml:"id,attr"`
}

type group struct {
	node `json:"node"`
	Children []*node `json:"children" xml:children>child"`
}

func (self *group) WriteJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(self)
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
	g.Id, isint = msg.Arguments[1].(int32)
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
	if numChildren < 0 {
		return nil, fmt.Errorf("expected numChildren >= 0, got %d", numChildren)
	}
	g.Children = make([]*node, numChildren)
	// get the childrens' ids
	var nodeID, numControls, numSubChildren int32
	for i := 3; i < numArgs; {
		nodeID, isint = msg.Arguments[i].(int32)
		if !isint {
			v := msg.Arguments[i]
			t := reflect.TypeOf(v)
			return nil, fmt.Errorf("expected arg %d (nodeID) to be int32, got %s (%v)", i, t, v)
		}
		g.Children[i-3] = &node{nodeID}
		// get the number of children of this node
		// if -1 this is a synth, if >= 0 this is a group
		numSubChildren, isint = msg.Arguments[i+1].(int32)
		if !isint {
			v := msg.Arguments[i]
			t := reflect.TypeOf(v)
			return nil, fmt.Errorf("expected arg %d (numControls) to be int32, got %s (%v)", i, t, v)
		}
		if numSubChildren == -1 {
			// synth
			i += 3
			numControls, isint = msg.Arguments[i].(int32)
			if !isint {
				v := msg.Arguments[i]
				t := reflect.TypeOf(v)
				return nil, fmt.Errorf("expected arg %d (numControls) to be int32, got %s (%v)", i, t, v)
			}
			i += 1 + int(numControls*2)
		} else if numSubChildren >= 0 {
			// group
			i += 2
		}
	}
	return g, nil
}
