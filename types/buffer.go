package types

// Buffer is a client-side representation of an scsynth audio buffer
type Buffer interface {
	// Num returns the index of this buffer in
	// scsynth's global buffers array
	Num() int32
	// Gen generates data for a buffer using a routine
	// (see BufferRoutine constants)
	Gen(routine string, flags int, args ...float32) error
}
