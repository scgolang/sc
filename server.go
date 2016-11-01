package sc

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// DefaultServerPort is the default listening port for scsynth.
const DefaultServerPort = 57120

// ErrNoScsynth happens when you try to start a SuperCollider
// server but do not have an scsynth executable in your PATH.
var ErrNoScsynth = errors.New("Please install scsynth somewhere in your PATH.")

// Server represents a running instance of scsynth.
type Server struct {
	*exec.Cmd

	Network      string
	Port         int
	StartTimeout time.Duration
}

// getServerPath gets the path to the scsynth executable.
func (s *Server) getServerPath() (string, error) {
	for _, file := range strings.Split(ServerPath, ":") {
		ok, err := isExecutable(file)
		if err != nil {
			return "", err
		}
		if ok {
			return file, nil
		}
	}

	path, hasPath := os.LookupEnv("PATH")
	if !hasPath {
		return "", ErrNoScsynth
	}

	for _, file := range strings.Split(path, ":") {
		ok, err := isExecutable(file)
		if err != nil {
			return "", err
		}
		if ok {
			return file, nil
		}
	}
	return "", ErrNoScsynth
}

// isExecutable returns true if the provided file is executable, false otherwise.
// It also returns any error that occurs while trying to read the file.
func isExecutable(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	mode := info.Mode()
	return ((mode & 0x01) | (mode & 0x08) | (mode & 0x40)) != 0, nil
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

	serverPath, err := s.getServerPath()
	if err != nil {
		return err
	}

	s.Cmd = exec.Command(serverPath, args...)
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
