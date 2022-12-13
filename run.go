package tester

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/shivuslr41/grpc-tester/exec"
	"golang.org/x/sync/errgroup"
)

func (r *Runner) Run(reader func(io.ReadCloser) error) error {
	cmd := exec.NewCMD(
		fmt.Sprintf(
			"grpcurl %s %s %s -d @ %s %s",
			r.GrpcurlFlags,
			r.tlsFlag(),
			r.protoFlag(),
			r.Server,
			r.Endpoint,
		),
	)

	writer, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	g := new(errgroup.Group)

	g.Go(func() error {
		return r.write(writer)
	})

	stderrReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	g.Go(func() error {
		return readStdErr(stderrReader)
	})

	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	g.Go(func() error {
		return reader(stdoutReader)
	})

	if err := cmd.Start(); err != nil {
		return err
	}

	if err = g.Wait(); err != nil {
		return err
	}
	return nil
}

func (r *Runner) write(writer io.WriteCloser) error {
	defer writer.Close()
	if r.StreamPayload {
		data := r.Data
		if !r.testerCall {
			err := json.Unmarshal([]byte(r.Data[0].(string)), &data)
			if err != nil {
				return err
			}
		}
		for i := range data {
			b, err := json.Marshal(data[i])
			if err != nil {
				return err
			}
			io.WriteString(writer, string(b))
		}
	} else {
		data := fmt.Sprint(r.Data[0])
		if r.testerCall {
			b, err := json.Marshal(r.Data[0])
			if err != nil {
				return err
			}
			data = string(b)
		}
		io.WriteString(writer, data)
	}
	return nil
}

func (r *Runner) Execute() {
	r.print()
}
