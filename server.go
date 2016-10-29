package sc

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

// DefaultServerPort is the default listening port for scsynth.
const DefaultServerPort = 57120

// Server represents a running instance of scsynth.
type Server struct {
	*exec.Cmd

	Network      string
	Port         int
	StartTimeout time.Duration
}

// args gets the command line args to scsynth
func (s *Server) args() ([]string, error) {
	// Get the port.
	if s.Port <= 0 {
		s.Port = DefaultServerPort
	}
	portArg := strconv.FormatInt(int64(s.Port), 10)

	// Create the args slice.
	args := []string{}

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
	if err := s.Cmd.Start(); err != nil {
		return err
	}

	// Wait until the server returns a status.
	_, err = NewClient(s.Network, "0.0.0.0:0", fmt.Sprintf("0.0.0.0:%d", s.Port), 5*time.Second)
	return err
}

// Stop stops a running server.
func (s *Server) Stop() error {
	return s.Process.Kill()
}
