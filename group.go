package sc

import (
	"fmt"

	"github.com/scgolang/osc"
)

const (
	gQueryTree      = "/g_queryTree"
	gQueryTreeReply = "/g_queryTree.reply"
)

// Node is a node in a synth execution tree.
// It could be a synth node or a group.
type Node interface {
	Free() error
	ID() int32
}

// SynthNode is a node in a graph
type SynthNode struct {
	id int32
}

// Free frees the node.
func (g *SynthNode) Free() error {
	return nil
}

// ID returns the node ID.
func (g *SynthNode) ID() int32 {
	return g.id
}

// GroupNode is a group of nodes.
type GroupNode struct {
	Children []Node

	id     int32
	client *Client
}

// Free frees all the nodes in a group.
// TODO
func (g *GroupNode) Free() error {
	return nil
}

// FreeAll frees all the nodes in a group recursively.
// TODO
func (g *GroupNode) FreeAll() error {
	return nil
}

// ID returns the node ID.
func (g *GroupNode) ID() int32 {
	return g.id
}

// Synth adds a synth to a group
func (g *GroupNode) Synth(defName string, id, action int32, ctls map[string]float32) (*Synth, error) {
	return g.client.Synth(defName, id, action, g.id, ctls)
}

// Synths creates multiple synth nodes at once with an OSC bundle.
func (g *GroupNode) Synths(args []SynthArgs) error {
	for _, arg := range args {
		arg.Target = g.id
	}
	return g.client.Synths(args)
}

// newGroup creates a new Group structure
func newGroup(client *Client, id int32) *GroupNode {
	return &GroupNode{
		client:   client,
		id:       id,
		Children: []Node{},
	}
}

// parseGroup parses information about a group from a message received at /g_queryTree.
// It *does not* recursively query for child groups.
func parseGroup(msg osc.Message) (*GroupNode, error) {
	// g_queryTree replies should have at least 3 arguments
	numArgs := len(msg.Arguments)
	if numArgs < 3 {
		return nil, fmt.Errorf("expected 3 arguments for message, got %d", numArgs)
	}
	// get the id of the group this reply is for
	nodeID, err := msg.Arguments[1].ReadInt32()
	if err != nil {
		return nil, err
	}
	g := &GroupNode{id: nodeID}

	if err := g.getChildren(msg); err != nil {
		return nil, err
	}
	return g, nil
}

// getChildren gets all the children of a group.
func (g *GroupNode) getChildren(msg osc.Message) error {
	// initialize the children array
	numChildren, err := msg.Arguments[2].ReadInt32()
	if err != nil {
		return err
	}
	if numChildren < 0 {
		return fmt.Errorf("expected numChildren >= 0, got %d", numChildren)
	}
	g.Children = make([]Node, numChildren)
	return g.getChildrenR(msg, 0, 3)
}

func (g *GroupNode) getChildrenR(msg osc.Message, childIndex, startIndex int) error {
	nodeID, err := msg.Arguments[startIndex].ReadInt32()
	if err != nil {
		return err
	}
	g.Children[childIndex] = &GroupNode{id: nodeID}

	// get the number of children of this node
	// if -1 this is a synth, if >= 0 this is a group
	numSubChildren, err := msg.Arguments[startIndex+1].ReadInt32()
	if err != nil {
		return err
	}
	if numSubChildren == -1 {
		// synth
		_, err := msg.Arguments[startIndex+2].ReadInt32()
		if err != nil {
			return err
		}
	} else if numSubChildren >= 0 {
	}
	return nil
}
