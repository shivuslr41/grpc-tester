package exec

import (
	"io"
)

func (e *Execute) StdinPipe() (io.WriteCloser, error) {
	cmd := getGrpcurlCmd(e.Flags)
	e.cmd = cmd
	return cmd.StdinPipe()
}

func (e *Execute) StdoutPipe() (io.ReadCloser, error) {
	return e.cmd.StdoutPipe()
}

func (e *Execute) StderrPipe() (io.ReadCloser, error) {
	return e.cmd.StderrPipe()
}

func (e *Execute) Start() error {
	return e.cmd.Start()
}
