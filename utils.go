package tester

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/shivuslr41/grpc-tester/jq"
)

// stores extracted data
const file = "variables.json"

var variables = make(map[string]any)

// grpcurl tls configuration
func (l *Lister) tlsFlag() string {
	if l.TLS {
		return ""
	}
	return "--plaintext"
}

// grpcurl proto files and path configuration
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

// replaces file configs from global -G configs if provided
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

// removes empty vals from slice
func removeEmptyStrings(s []string) []string {
	var ss []string
	for i := range s {
		if s[i] != "" {
			ss = append(ss, s[i])
		}
	}
	return ss
}

// reads stderr from pipe
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

// prints error and exits
func printErrAndExit(err error) {
	fmt.Print(err)
	os.Exit(1)
}

// format JSON string into "jq" format
func (t *T) format(b []byte) error {
	str, err := jq.Format(string(b))
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), &t.Response)
}

// load extracted data from variables.json file to variables map
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

// save extracted result data to variables.json file from variables map
func save() error {
	b, err := json.MarshalIndent(variables, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, b, 0644)
}
