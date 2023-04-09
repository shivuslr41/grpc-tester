package tester

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/shivuslr41/grpc-tester/exec"
	"golang.org/x/sync/errgroup"
)

// Run execute grpcurl(calls grpc API) with the given configurations
func (r *Runner) Run(reader func(io.ReadCloser) error) error {
	// if global flag -G is used, ignores test file configs
	if GConf.Use {
		r.replaceGconf()
	}

	// prepare grpcurl cmd to call grpc server
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

	// create writer for passing payloads/requests
	writer, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	// create errgroup
	g := new(errgroup.Group)

	// send requests
	g.Go(func() error {
		return r.write(writer)
	})

	// create reader to catch any server errors
	stderrReader, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	// listen to errors in background
	g.Go(func() error {
		return readStdErr(stderrReader)
	})

	// create reader to collect server responses
	stdoutReader, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// listen for results in the background
	g.Go(func() error {
		return reader(stdoutReader)
	})

	// execute the grpcurl cmd
	if err := cmd.Start(); err != nil {
		return err
	}

	// wait for any error from error group
	if err = g.Wait(); err != nil {
		return err
	}
	return nil
}

// write sends/streams grpc requests to stdin for grpcurl
func (r *Runner) write(writer io.WriteCloser) error {
	defer writer.Close()
	// write to stdin based type of payload
	// if streamPayload then divide the requests and write into stdin,
	// else write without dividing them
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

// Execute starts run command
func (r *Runner) Execute() {
	r.print()
}
