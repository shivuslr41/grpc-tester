package exec

import (
	"fmt"
	"os"
	"strings"
)

func (e *Execute) GetStdErr() {
	//TODO: handle error rather then sending directly to user's terminal
	e.cmd.Stderr = os.Stderr
}

func (e *Execute) GetStdout() ([]byte, error) {
	e.cmd = getCmd(e.Flags)
	return e.cmd.Output()
}

func (e *Execute) GetCombinedStdout() ([]string, error) {
	cmd := getGrpcurlCmd(e.Flags)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s%s", string(b), err.Error())
	}
	return removeEmptyStrings(strings.Split(string(b), "\n")), nil
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
