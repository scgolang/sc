package sc

import (
	"encoding/binary"
	"io"
)

// ParamName represents a parameter name of a synthdef
type ParamName struct {
	Name  string `json:"name,omitempty"`
	Index int32  `json:"index"`
}

func (pn *ParamName) Write(w io.Writer) error {
	err := newPstring(pn.Name).Write(w)
	if err != nil {
		return err
	}
	return binary.Write(w, byteOrder, int32(pn.Index))
}

// readParamName reads a ParamName from an io.Reader
func readParamName(r io.Reader) (*ParamName, error) {
	name, err := readPstring(r)
	if err != nil {
		return nil, err
	}
	var idx int32
	err = binary.Read(r, byteOrder, &idx)
	if err != nil {
		return nil, err
	}
	pn := ParamName{name.String(), idx}
	return &pn, nil
}
