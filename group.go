package sc

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/scgolang/osc"
	"io"
	"reflect"
)

const (
	gQueryTree      = "/g_queryTree"
	gQueryTreeReply = "/g_queryTree.reply"
)

// Node is a node in a graph
type Node struct {
	id int32 `json:"id" xml:"id,attr"`
}

// Group is a group of synth nodes
type Group struct {
	Node     `json:"node" xml:"node"`
	children []*Node `json:"children" xml:"children>child"`
	client   *Client
}

// Synth adds a synth to a group
func (self *Group) Synth(defName string, id, action int32, ctls map[string]float32) (*Synth, error) {
	return self.client.Synth(defName, id, action, self.Node.id, ctls)
}

// Free frees all the nodes in a group
func (self *Group) Free() error {
	return nil
}

// FreeAll frees all the nodes in a group recursively
func (self *Group) FreeAll() error {
	return nil
}

// WriteJSON writes a JSON representation of a group to an io.Writer
func (self *Group) WriteJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(self)
}

// WriteXML writes an XML representation of a group to an io.Writer
func (self *Group) WriteXML(w io.Writer) error {
	enc := xml.NewEncoder(w)
	return enc.Encode(self)
}

// newGroup creates a new Group structure
func newGroup(client *Client, id int32) *Group {
	return &Group{
		Node:     Node{id: id},
		children: make([]*Node, 0),
		client:   client,
	}
}

// parseGroup parses information about a group from a message
// received at /g_queryTree
// it *does not* recursively query for child groups
func parseGroup(msg *osc.Message) (*Group, error) {
	// return an error if msg.Address is not right
	if msg.Address != gQueryTreeReply {
		return nil, fmt.Errorf("msg.Address should be %s, got %s", gQueryTreeReply, msg.Address)
	}
	// g_queryTree replies should have at least 3 arguments
	g, numArgs := new(Group), msg.CountArguments()
	if numArgs < 3 {
		return nil, fmt.Errorf("expected 3 arguments for message, got %d", numArgs)
	}
	// get the id of the group this reply is for
	var isint bool
	g.Node.id, isint = msg.Arguments[1].(int32)
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
	g.children = make([]*Node, numChildren)
	// get the childrens' ids
	var nodeID, numControls, numSubChildren int32
	for i := 3; i < numArgs; {
		nodeID, isint = msg.Arguments[i].(int32)
		if !isint {
			v := msg.Arguments[i]
			t := reflect.TypeOf(v)
			return nil, fmt.Errorf("expected arg %d (nodeID) to be int32, got %s (%v)", i, t, v)
		}
		g.children[i-3] = &Node{nodeID}
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
