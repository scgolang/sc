package sc

/*

Our UgenGraph's are constructed of constant nodes and
ugen nodes.
Constant nodes have no children whereas ugen nodes can
have children, and the children can themselves be either
constants or ugens.
To convert a UgenGraph to a synthdef we must flatten the
tree to the scsynth binary synthdef format.
This is basically flattening the tree into a list of
constants and a list of ugens.

*/
