package sc

import (
	"fmt"
)

// NetAddr represents a UDP network address
type NetAddr struct {
	addr string
	port int
}

func (self *NetAddr) Addr() string {
	return self.addr
}

func (self *NetAddr) Port() int {
	return self.port
}

func (self *NetAddr) Network() string {
	return "udp"
}

func (self *NetAddr) String() string {
	return fmt.Sprintf("%s:%d", self.addr, self.port)
}

func NewAddr(addr string, port int) *NetAddr {
	a := NetAddr{addr, port}
	return &a
}
