package sc

// Node defines a node in a ugen graph.
// The only thing common to all nodes is that they
// have an index.
// Note that for a given ugen graph the indices
// for constants and the indices for ugens
// are independent of each other.
type Node interface {
	Index() int32
}
