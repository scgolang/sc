package sc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
)

// Write writes a binary representation of a synthdef to an io.Writer.
// The binary representation written by this method is
// the data that scsynth expects at its /d_recv endpoint.
func (def *Synthdef) Write(w io.Writer) error {
	if err := def.writeHead(w); err != nil {
		return err
	}
	// write synthdef name
	name := newPstring(def.Name)
	if err := name.Write(w); err != nil {
		return err
	}
	if err := def.writeConstants(w); err != nil {
		return err
	}
	if err := def.writeParams(w); err != nil {
		return err
	}
	// write number of ugens
	if err := binary.Write(w, byteOrder, int32(len(def.Ugens))); err != nil {
		return err
	}
	// write ugens
	for _, u := range def.Ugens {
		if err := u.Write(w); err != nil {
			return err
		}
	}
	// write number of variants
	if err := binary.Write(w, byteOrder, int16(len(def.Variants))); err != nil {
		return err
	}
	// write variants
	for _, v := range def.Variants {
		if err := v.Write(w); err != nil {
			return err
		}
	}
	return nil
}

// writeHead writes everything leading up to the synthdef name
// to an io.Writer.
func (def *Synthdef) writeHead(w io.Writer) error {
	written, err := w.Write(bytes.NewBufferString(synthdefStart).Bytes())
	if err != nil {
		return err
	}
	if written != len(synthdefStart) {
		return fmt.Errorf("Could not write synthdef")
	}
	// write synthdef version
	if err := binary.Write(w, byteOrder, int32(synthdefVersion)); err != nil {
		return err
	}
	// write number of synthdefs
	if err := binary.Write(w, byteOrder, int16(1)); err != nil {
		return err
	}
	return nil
}

// writeConstants writes the number of constants and the values
// to an io.Writer.
func (def *Synthdef) writeConstants(w io.Writer) error {
	// write number of constants
	if err := binary.Write(w, byteOrder, int32(len(def.Constants))); err != nil {
		return err
	}
	// write constant values
	for _, constant := range def.Constants {
		if err := binary.Write(w, byteOrder, constant); err != nil {
			return err
		}
	}
	return nil
}

// writeParams writes the number of synthdef params,
// the initial param values, and the param names
// to an io.Writer.
func (def *Synthdef) writeParams(w io.Writer) error {
	// write number of params
	if err := binary.Write(w, byteOrder, int32(len(def.ParamNames))); err != nil {
		return err
	}
	// write initial param values
	for _, val := range def.InitialParamValues {
		if err := binary.Write(w, byteOrder, val); err != nil {
			return err
		}
	}
	// write number of param names
	if err := binary.Write(w, byteOrder, int32(len(def.ParamNames))); err != nil {
		return err
	}
	// write param names
	for _, p := range def.ParamNames {
		if err := p.Write(w); err != nil {
			return err
		}
	}
	return nil
}

// WriteJSON writes a json-formatted representation of a
// synthdef to an io.Writer.
func (def *Synthdef) WriteJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(def)
}

// WriteXML writes an xml-formatted representation of a synthdef
// to an io.Writer.
func (def *Synthdef) WriteXML(w io.Writer) error {
	enc := xml.NewEncoder(w)
	return enc.Encode(def)
}
