package sc

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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

const ServerReadyMessage = "server ready"

// Start starts a new instance of scsynth.
// If the server doesn't print a line containing ServerReadyMessage
// within the timeout then ErrTimeout is returned.
func (s *Server) Start(timeout time.Duration) (io.ReadCloser, io.ReadCloser, error) {
	args, err := s.args()
	if err != nil {
		return nil, nil, err
	}

	serverPath, err := s.getServerPath()
	if err != nil {
		return nil, nil, err
	}

	s.Cmd = exec.Command(serverPath, args...)
	stdout, err := s.Cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderr, err := s.Cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}
	if err := s.Cmd.Start(); err != nil {
		return nil, nil, err
	}

	// Wait until the server prints a ready message.
	var (
		scanner = bufio.NewScanner(stdout)
		start   = time.Now()
	)
	for i := 0; scanner.Scan(); i++ {
		if time.Now().Sub(start) > timeout {
			return nil, nil, ErrTimeout
		}
		if strings.Index(scanner.Text(), ServerReadyMessage) == -1 {
			continue
		} else {
			break
		}
	}
	return stdout, stderr, scanner.Err()
}

// Stop stops a running server.
func (s *Server) Stop() error {
	return s.Process.Kill()
}
