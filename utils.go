package tester

import (
	"fmt"
	"io"
	"os"
)

func (l *Lister) tlsFlag() string {
	var tls string
	if !l.TLS {
		tls = "--plaintext"
	}
	return tls
}

func (l *Lister) protoFlag() string {
	var protoOption string
	if l.ProtoPath != "" {
		protoOption = fmt.Sprintf(
			"--import-path %s --proto %s",
			l.ProtoPath,
			l.ProtoFile,
		)
	}
	return protoOption
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

func readErr(rc io.ReadCloser) error {
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
