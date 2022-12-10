package exec

import (
	"io"
)

func (e *Execute) StdinPipe() (io.WriteCloser, error) {
	e.cmd = getGrpcurlCmd(e.Flags)
	return e.cmd.StdinPipe()
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
