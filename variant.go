package sc

import (
	"encoding/binary"
	"io"
)

// Variant
type variant struct {
	Name               Pstring   `json:'name,omitempty'`
	InitialParamValues []float32 `json:'initialParamValues'`
}

// readVariant read a Variant from an io.Reader
func readVariant(r io.Reader, numParams int32) (*variant, error) {
	name, err := readPstring(r)
	if err != nil {
		return nil, err
	}
	paramValues := make([]float32, numParams)
	for i := 0; int32(i) < numParams; i++ {
		err = binary.Read(r, byteOrder, &paramValues[i])
		if err != nil {
			return nil, err
		}
	}
	v := variant{*name, paramValues}
	return &v, nil
}
