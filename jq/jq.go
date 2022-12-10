package jq

import (
	"fmt"
	"time"

	"github.com/shivuslr41/grpc-tester/exec"
)

func Format(j string) string {
	exe := exec.NewExec()
	exe.Flags = fmt.Sprintf("echo '%s' | jq -s '.'", j)
	b, err := exe.GetStdout()
	exe.GetStdErr()
	time.Sleep(2 * time.Second)
	if err != nil {
		panic(err)
	}
	return string(b)
}
