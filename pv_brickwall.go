package sc

// PVBrickWall clears bins above or below a cutoff point.
type PVBrickWall struct {
	// FFT Buffer.
	Buffer Input

	// Can range between -1 and +1.
	// If wipe == 0 then there is no effect.
	// If wipe > 0 then it acts like a high pass filter, clearing bins from the bottom up.
	// If wipe < 0 then it acts like a low pass filter, clearing bins from the top down.
	Wipe Input
}

func (bw *PVBrickWall) defaults() {
	if bw.Wipe == nil {
		bw.Wipe = C(0)
	}
}

// Rate creates a new ugen at a specific rate.
// If rate is an unsupported value this method will cause a runtime panic.
// If Buffer is nil this method will panic.
func (bw PVBrickWall) Rate(rate int8) Input {
	CheckRate(rate)
	if bw.Buffer == nil {
		panic("PVBrickWall requires a Buffer parameter")
	}
	(&bw).defaults()
	return NewInput("PV_BrickWall", rate, 0, 1, bw.Buffer, bw.Wipe)
}
