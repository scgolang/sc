package sc

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
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
	Controls map[string]string
	DefName  string

	client *Client
	id     int32
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

	client *Client
	id     int32
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
		client: client,
		id:     id,
	}
}

// parseGroup parses information about a group from a reply to /g_queryTree.
func (c *Client) parseGroup(msg osc.Message) (*GroupNode, error) {
	n, _, err := c.parseNodeFrom(msg, 0)
	if err != nil {
		return nil, err
	}
	gn, ok := n.(*GroupNode)
	if !ok {
		return nil, errors.Wrap(err, "type assertion from Node to *GroupNode")
	}
	return gn, nil
}

// parseNodeFrom parses information about a group from a message received at /g_queryTree.
// It *does not* recursively query for child groups.
func (c *Client) parseNodeFrom(msg osc.Message, startIndex int) (Node, int, error) {
	argsConsumed := 0

	// We should have at least 2 arguments for nodes contained in the group.
	if numArgs := len(msg.Arguments); numArgs < 2 {
		return nil, 0, fmt.Errorf("expected 2 arguments for message, got %d", numArgs)
	}
	// get the id of the group this reply is for
	nodeID, err := msg.Arguments[startIndex].ReadInt32()
	if err != nil {
		return nil, 0, err
	}
	argsConsumed++

	numChildren, err := msg.Arguments[startIndex+1].ReadInt32()
	if err != nil {
		return nil, 0, err
	}
	argsConsumed++

	if numChildren < 0 {
		// Synth node.
		node, synthArgsConsumed, err := c.parseSynthNodeFrom(msg, nodeID, startIndex+argsConsumed)
		return node, argsConsumed + synthArgsConsumed, errors.Wrap(err, "parsing synth node")
	}
	g := &GroupNode{
		Children: make([]Node, numChildren),
		client:   c,
		id:       nodeID,
	}
	startIndex = startIndex + 2

	for i := 0; int32(i) < numChildren; i++ {
		node, childArgsConsumed, err := c.parseNodeFrom(msg, startIndex)
		if err != nil {
			return nil, 0, err
		}
		startIndex += childArgsConsumed
		argsConsumed += childArgsConsumed
		g.Children[i] = node
	}
	return g, argsConsumed, nil
}

func (c *Client) parseSynthNodeFrom(msg osc.Message, nodeID int32, startIndex int) (n Node, argsConsumed int, err error) {
	// Assume we are starting at the synthdef name.
	defName, err := msg.Arguments[startIndex+argsConsumed].ReadString()
	if err != nil {
		return nil, 0, errors.Wrap(err, "reading synthdef name")
	}
	argsConsumed++

	numControls, err := msg.Arguments[startIndex+argsConsumed].ReadInt32()
	if err != nil {
		return nil, 0, errors.Wrap(err, "reading number of synth controls")
	}
	argsConsumed++

	sn := &SynthNode{
		DefName:  defName,
		Controls: make(map[string]string, numControls),
		client:   c,
		id:       nodeID,
	}
	for i := 0; int32(i) < numControls; i++ {
		controlName, err := msg.Arguments[startIndex+argsConsumed].ReadString()
		if err != nil {
			return nil, 0, errors.Wrap(err, "reading synth control name")
		}
		argsConsumed++

		var cvstr string
		cvflt, err := msg.Arguments[startIndex+argsConsumed].ReadFloat32()
		if err == nil {
			cvstr = strconv.FormatFloat(float64(cvflt), 'f', -1, 32)
		} else {
			cvstr, err = msg.Arguments[startIndex+argsConsumed].ReadString()
			if err != nil {
				return nil, 0, errors.Wrap(err, "reading synth control value")
			}
		}
		argsConsumed++
		sn.Controls[controlName] = cvstr
	}
	return sn, argsConsumed, nil
}
