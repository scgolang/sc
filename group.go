package sc

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/scgolang/osc"
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
	if msg.Address() != gQueryTreeReply {
		return nil, fmt.Errorf("msg.Address should be %s, got %s", gQueryTreeReply, msg.Address())
	}
	// g_queryTree replies should have at least 3 arguments
	g, numArgs := new(Group), msg.CountArguments()
	if numArgs < 3 {
		return nil, fmt.Errorf("expected 3 arguments for message, got %d", numArgs)
	}
	// get the id of the group this reply is for
	nodeID, err := msg.ReadInt32()
	if err != nil {
		return nil, err
	}
	g.Node.id = nodeID

	// initialize the children array
	numChildren, err := msg.ReadInt32()
	if err != nil {
		return nil, err
	}
	if numChildren < 0 {
		return nil, fmt.Errorf("expected numChildren >= 0, got %d", numChildren)
	}
	g.children = make([]*Node, numChildren)
	// get the childrens' ids
	var numControls, numSubChildren int32
	for i := 3; i < numArgs; {
		nodeID, err = msg.ReadInt32()
		if err != nil {
			return nil, err
		}
		g.children[i-3] = &Node{nodeID}
		// get the number of children of this node
		// if -1 this is a synth, if >= 0 this is a group
		numSubChildren, err = msg.ReadInt32()
		if err != nil {
			return nil, err
		}
		if numSubChildren == -1 {
			// synth
			i += 3
			numControls, err = msg.ReadInt32()
			if err != nil {
				return nil, err
			}
			i += 1 + int(numControls*2)
		} else if numSubChildren >= 0 {
			// group
			i += 2
		}
	}
	return g, nil
}
