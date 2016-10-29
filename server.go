package sc

import (
	"fmt"
	"os/exec"
	"strconv"
)

// Server represents a running instance of scsynth.
type Server struct {
	*exec.Cmd

	Network string
	Port    int
}

// args gets the command line args to scsynth
func (s *Server) args() ([]string, error) {
	args := []string{}

	var (
		portArg = strconv.FormatInt(int64(s.Port), 10)
	)
	switch s.Network {
	default:
		return nil, fmt.Errorf("unrecognized network type: %s", s.Network)
	case "udp":
		args = append(args, "-u", portArg)
	case "tcp":
		args = append(args, "-t", portArg)
	}

	return args, nil
}

// Start starts a new instance of scsynth.
func (s *Server) Start() error {
	args, err := s.args()
	if err != nil {
		return err
	}
	s.Cmd = exec.Command(ServerPath, args...)
	return s.Cmd.Start()
}

// Stop stops a running server.
func (s *Server) Stop() error {
	return s.Process.Kill()
}
