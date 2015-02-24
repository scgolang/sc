package sc

import (
	"encoding/binary"
	"io"
)

// ParamName represents a parameter name of a synthdef
type ParamName struct {
	Name  string `json:'name,omitempty'`
	Index int32   `json:'index,omitempty'`
}

func (p *ParamName) Write(w io.Writer) error {
	// FIXME
	// if we := p.Name.Write(w); we != nil {
	// 	return we
	// }
	return binary.Write(w, byteOrder, p.Index)
}

// ReadParamName reads a ParamName from an io.Reader
func ReadParamName(r io.Reader) (*ParamName, error) {
	name, err := ReadPstring(r)
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
