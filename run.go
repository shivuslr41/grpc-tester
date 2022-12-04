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
		for i := range r.Data {
			var iface []interface{}
			err := json.Unmarshal([]byte(fmt.Sprint(r.Data[i])), &iface)
			if err != nil {
				log.Fatal(err)
			}
			for j := range iface {
				b, err := json.Marshal(iface[j])
				if err != nil {
					log.Fatal(err)
				}
				io.WriteString(writer, string(b))
			}
		}
	} else {
		for i := range r.Data {
			var iface interface{}
			err := json.Unmarshal([]byte(fmt.Sprint(r.Data[i])), &iface)
			if err != nil {
				log.Fatal(err)
			}
			b, err := json.Marshal(iface)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(writer, string(b))
		}
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
		return nil, err
	}

	stdoutReader, err := exe.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := exe.Start(); err != nil {
		return nil, err
	}
	// for now only check any error at start
	go func() {
		b, err := io.ReadAll(stderrReader)
		if err != nil {
			panic(err)
		}
		fmt.Println("------->", string(b))
	}()
	return stdoutReader, nil
}
