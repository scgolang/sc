package gosc

// import (
// 	"encoding/binary"
// 	"os"
// )

type SynthDef interface {
	Send(addr NetAddr) error
	Write(path string) error
}

type testTone struct {
}

func (self *testTone) Send(addr NetAddr) error {
	return nil
}

func (self *testTone) Write(path string) error {
	return nil
}

func TestTone() SynthDef {
	tt := testTone{}
	return &tt
}
