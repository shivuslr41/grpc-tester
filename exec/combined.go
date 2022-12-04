package exec

import (
	"fmt"
	"strings"
)

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
