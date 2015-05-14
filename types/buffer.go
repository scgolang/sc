package types

// Buffer is a client-side representation of an scsynth audio buffer
type Buffer interface {
	// Num returns the index of this buffer in
	// scsynth's global buffers array
	Num() int32
}
