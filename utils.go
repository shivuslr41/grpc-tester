package tester

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/shivuslr41/grpc-tester/jq"
)

const file = "variables.json"

var variables = make(map[string]any)

func (l *Lister) tlsFlag() string {
	if l.TLS {
		return ""
	}
	return "--plaintext"
}

func (l *Lister) protoFlag() string {
	if l.ProtoPath != "" {
		return fmt.Sprintf(
			"--import-path %s --proto %s",
			l.ProtoPath,
			l.ProtoFile,
		)
	}
	return ""
}

func (r *Runner) replaceGconf() {
	if GConf.Server != "" {
		r.Server = GConf.Server
	}
	if GConf.Endpoint != "" {
		r.Endpoint = GConf.Endpoint
	}
	if GConf.ProtoPath != "" {
		r.ProtoPath = GConf.ProtoPath
	}
	if GConf.ProtoFile != "" {
		r.ProtoFile = GConf.ProtoFile
	}
	r.StreamPayload = GConf.StreamPayload
	r.TLS = GConf.TLS
}

func removeEmptyStrings(s []string) []string {
	var ss []string
	for i := range s {
		if s[i] != "" {
			ss = append(ss, s[i])
		}
	}
	return ss
}

func readStdErr(rc io.ReadCloser) error {
	b, err := io.ReadAll(rc)
	if err != nil {
		return err
	}
	if len(b) != 0 {
		return fmt.Errorf("%s", string(b))
	}
	return nil
}

func printErrAndExit(err error) {
	fmt.Print(err)
	os.Exit(1)
}

func (t *T) format(b []byte) error {
	str, err := jq.Format(string(b))
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), &t.Response)
}

func load() error {
	b, err := os.ReadFile(file)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil
		}
		return err
	}
	return json.Unmarshal(b, &variables)
}

func save() error {
	b, err := json.MarshalIndent(variables, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, b, 0644)
}
