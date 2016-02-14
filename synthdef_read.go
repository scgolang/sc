package sc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// ReadSynthdef reads a synthdef from an io.Reader
func ReadSynthdef(r io.Reader) (*Synthdef, error) {
	// read the type
	if err := readSynthdefType(r); err != nil {
		return nil, err
	}
	// read version
	if err := readSynthdefVersion(r); err != nil {
		return nil, err
	}
	// read number of synth defs
	if err := readNumberOfSynthdefs(r); err != nil {
		return nil, err
	}
	// read synthdef name
	defName, err := readPstring(r)
	if err != nil {
		return nil, err
	}
	// read constants
	constants, err := readSynthdefConstants(r)
	if err != nil {
		return nil, err
	}
	// read param initial values and names
	initialValues, paramNames, err := readSynthdefParams(r)
	if err != nil {
		return nil, err
	}
	// read ugens
	ugens, err := readSynthdefUgens(r)
	if err != nil {
		return nil, err
	}
	// read variants
	variants, err := readSynthdefVariants(r, len(initialValues))
	if err != nil {
		return nil, err
	}
	synthDef := Synthdef{
		Name:               defName.String(),
		Constants:          constants,
		InitialParamValues: initialValues,
		ParamNames:         paramNames,
		Ugens:              ugens,
		Variants:           variants,
		seen:               []Ugen{},
	}
	return &synthDef, nil
}

// readSynthdefType reads the first 4 bytes of a synthdef file
// and returns an error if it isn't a supported type.
// Otherwise it returns nil.
func readSynthdefType(r io.Reader) error {
	startLen := len(synthdefStart)
	start := make([]byte, startLen)
	read, err := r.Read(start)
	if err != nil {
		return err
	}
	if read != startLen {
		return fmt.Errorf("Only read %d bytes of synthdef file", read)
	}
	if actual := bytes.NewBuffer(start).String(); actual != synthdefStart {
		return fmt.Errorf("synthdef started with %s instead of %s", actual, synthdefStart)
	}
	return nil
}

// readSynthdefVersion reads the version of a synthdef file
// and returns an error if it is an unsupported version.
// Otherwise it returns nil.
func readSynthdefVersion(r io.Reader) error {
	var version int32
	if err := binary.Read(r, byteOrder, &version); err != nil {
		return err
	}
	if version != synthdefVersion {
		return fmt.Errorf("bad synthdef version %d", version)
	}
	return nil
}

// readNumberOfSynthdefs reads the number of synthdefs in
// a particular synthdef file.
// We only support 1 synthdef per file, so if this number is not 1
// this func will return an error.
func readNumberOfSynthdefs(r io.Reader) error {
	var numDefs int16
	if err := binary.Read(r, byteOrder, &numDefs); err != nil {
		return err
	}
	if numDefs != 1 {
		return fmt.Errorf("multiple synthdefs not supported")
	}
	return nil
}

// readSynthdefConstants reads the constants of a synthdef.
func readSynthdefConstants(r io.Reader) ([]float32, error) {
	// read number of constants
	var numConstants int32
	if err := binary.Read(r, byteOrder, &numConstants); err != nil {
		return nil, err
	}
	// read constants
	constants := make([]float32, numConstants)
	for i := 0; i < int(numConstants); i++ {
		if err := binary.Read(r, byteOrder, &constants[i]); err != nil {
			return nil, err
		}
	}
	return constants, nil
}

// readSynthdefParams reads the initial param values and param names
// of a synthdef.
func readSynthdefParams(r io.Reader) ([]float32, []ParamName, error) {
	// read number of parameters
	var numParams int32
	if err := binary.Read(r, byteOrder, &numParams); err != nil {
		return nil, nil, err
	}
	// read initial parameter values
	initialValues := make([]float32, numParams)
	for i := 0; i < int(numParams); i++ {
		if err := binary.Read(r, byteOrder, &initialValues[i]); err != nil {
			return nil, nil, err
		}
	}
	// read number of parameter names
	var numParamNames int32
	if err := binary.Read(r, byteOrder, &numParamNames); err != nil {
		return nil, nil, err
	}
	// read param names
	paramNames := make([]ParamName, numParamNames)
	for i := 0; int32(i) < numParamNames; i++ {
		pn, err := readParamName(r)
		if err != nil {
			return nil, nil, err
		}
		paramNames[i] = *pn
	}
	return initialValues, paramNames, nil
}

// readSynthdefUgens reads the ugens of a synthdef.
func readSynthdefUgens(r io.Reader) ([]*ugen, error) {
	// read number of ugens
	var numUgens int32
	if err := binary.Read(r, byteOrder, &numUgens); err != nil {
		return nil, err
	}
	// read ugens
	ugens := make([]*ugen, numUgens)
	for i := 0; int32(i) < numUgens; i++ {
		ugen, err := readugen(r)
		if err != nil {
			return nil, err
		}
		ugens[i] = ugen
	}
	return ugens, nil
}

// readSynthdefVariants reads the variants of a synthdef.
func readSynthdefVariants(r io.Reader, numParams int) ([]*Variant, error) {
	// read number of variants
	var numVariants int16
	if err := binary.Read(r, byteOrder, &numVariants); err != nil {
		return nil, err
	}
	// read variants
	variants := make([]*Variant, numVariants)
	for i := 0; int16(i) < numVariants; i++ {
		v, err := readVariant(r, int32(numParams))
		if err != nil {
			return nil, err
		}
		variants[i] = v
	}
	return variants, nil
}
