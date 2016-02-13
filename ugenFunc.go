package sc

// UgenFunc is a func that is used to define a ugen graph.
type UgenFunc func(p Params) Ugen

// Play corresponds to http://doc.sccode.org/Classes/Function.html#-play.
// It wraps the provided UgenFunc in a synthdef,
// sends this synthdef to a server instance with DefaultClient,
// then creates a new synth from the synthdef.
func Play(c *Client, f UgenFunc) error {
	// To implement this we need a DefaultServer and
	// a way to generate the names of the temp synthdefs.
	// If the ugen node returned by f is not Out,
	// wrap the node in an Out ugen with bus 0.
	return nil
}
