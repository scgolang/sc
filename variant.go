package sc

import (
	"encoding/binary"
	"io"
)

// variant
type Variant struct {
	Name               string    `json:'name,omitempty'`
	InitialParamValues []float32 `json:'initialParamValues'`
}

func (self *Variant) Write(w io.Writer) error {
	err := newPstring(self.Name).Write(w)
	if err != nil {
		return err
	}
	for _, v := range self.InitialParamValues {
		err = binary.Write(w, byteOrder, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// readVariant read a Variant from an io.Reader
func readVariant(r io.Reader, numParams int32) (*Variant, error) {
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
	v := Variant{name.String(), paramValues}
	return &v, nil
}
