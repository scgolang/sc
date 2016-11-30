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
	ID int32 `json:"id" xml:"id,attr"`
}

// Group is a group of synth nodes
type Group struct {
	Node     `json:"node" xml:"node"`
	Children []*Node `json:"children" xml:"children>child"`
	client   *Client
}

// Synth adds a synth to a group
func (g *Group) Synth(defName string, id, action int32, ctls map[string]float32) (*Synth, error) {
	return g.client.Synth(defName, id, action, g.Node.ID, ctls)
}

// Free frees all the nodes in a group
func (g *Group) Free() error {
	return nil
}

// FreeAll frees all the nodes in a group recursively
func (g *Group) FreeAll() error {
	return nil
}

// WriteJSON writes a JSON representation of a group to an io.Writer
func (g *Group) WriteJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(g)
}

// WriteXML writes an XML representation of a group to an io.Writer
func (g *Group) WriteXML(w io.Writer) error {
	enc := xml.NewEncoder(w)
	return enc.Encode(g)
}

// newGroup creates a new Group structure
func newGroup(client *Client, id int32) *Group {
	return &Group{
		Node:     Node{ID: id},
		Children: make([]*Node, 0),
		client:   client,
	}
}

// parseGroup parses information about a group from a message received at /g_queryTree.
// It *does not* recursively query for child groups.
func parseGroup(msg osc.Message) (*Group, error) {
	// return an error if msg.Address is not right
	if msg.Address != gQueryTreeReply {
		return nil, fmt.Errorf("msg.Address should be %s, got %s", gQueryTreeReply, msg.Address)
	}
	// g_queryTree replies should have at least 3 arguments
	var (
		g       = &Group{}
		numArgs = len(msg.Arguments)
	)
	if numArgs < 3 {
		return nil, fmt.Errorf("expected 3 arguments for message, got %d", numArgs)
	}
	// get the id of the group this reply is for
	nodeID, err := msg.Arguments[0].ReadInt32()
	if err != nil {
		return nil, err
	}
	g.Node.ID = nodeID

	if err := g.getChildren(msg); err != nil {
		return nil, err
	}
	return g, nil
}

// getChildren gets all the children of a group.
func (g *Group) getChildren(msg osc.Message) error {
	numArgs := len(msg.Arguments)

	// initialize the children array
	numChildren, err := msg.Arguments[1].ReadInt32()
	if err != nil {
		return err
	}
	if numChildren < 0 {
		return fmt.Errorf("expected numChildren >= 0, got %d", numChildren)
	}
	g.Children = make([]*Node, numChildren)

	// get the childrens' ids
	var numControls, numSubChildren int32

	for i := 3; i < numArgs; {
		j := (3 * i) - 7

		nodeID, err := msg.Arguments[j].ReadInt32()
		if err != nil {
			return err
		}
		g.Children[i-3] = &Node{ID: nodeID}
		// get the number of children of this node
		// if -1 this is a synth, if >= 0 this is a group
		numSubChildren, err = msg.Arguments[j+1].ReadInt32()
		if err != nil {
			return err
		}
		if numSubChildren == -1 {
			// synth
			i += 3
			numControls, err = msg.Arguments[j+2].ReadInt32()
			if err != nil {
				return err
			}
			i += 1 + int(numControls*2)
		} else if numSubChildren >= 0 {
			// group
			i += 2
		}
	}
	return nil
}
