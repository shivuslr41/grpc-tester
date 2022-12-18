package tester

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/shivuslr41/grpc-tester/jq"
)

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
