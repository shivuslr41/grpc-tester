package tester

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/shivuslr41/grpc-tester/exec"
)

func (r *Runner) Print() {
	reader, err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
	var out interface{}
	decoder := json.NewDecoder(reader)
	for decoder.More() {
		err = decoder.Decode(&out)
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}
	// time.Sleep(2 * time.Second)
	// TODO: wait for amy errors
}

func (r *Runner) write(writer io.WriteCloser) {
	defer writer.Close()
	if r.StreamPayload {
		data := r.Data
		if !r.testerCall {
			err := json.Unmarshal([]byte(r.Data[0].(string)), &data)
			if err != nil {
				log.Fatal(err)
			}
		}
		for i := range data {
			b, err := json.Marshal(data[i])
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(writer, string(b))
		}
	} else {
		data := fmt.Sprint(r.Data[0])
		if r.testerCall {
			b, err := json.Marshal(r.Data[0])
			if err != nil {
				panic(err)
			}
			data = string(b)
		}
		io.WriteString(writer, data)
	}
}

func (r *Runner) Run() (io.ReadCloser, error) {
	exe := exec.NewExec()
	exe.Flags = fmt.Sprintf("%s %s %s -d @ %s %s", r.GrpcurlFlags, r.tlsFlag(), r.protoFlag(), r.Server, r.Endpoint)
	writer, err := exe.StdinPipe()
	if err != nil {
		return nil, err
	}
	// TODO: error handling
	// explore errgroup
	go r.write(writer)

	stderrReader, err := exe.StderrPipe()
	if err != nil {
		fmt.Println("stderr", err)
		return nil, err
	}

	stdoutReader, err := exe.StdoutPipe()
	if err != nil {
		fmt.Println("stdout err", err)
		return nil, err
	}

	if err := exe.Start(); err != nil {
		fmt.Println("start err", err)
		return nil, err
	}
	// for now only check any error at start
	go func() {
		b, err := io.ReadAll(stderrReader)
		if err != nil {
			panic(err)
		}
		fmt.Println("stderr ------->", string(b))
	}()
	return stdoutReader, nil
}
