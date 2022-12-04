package exec

import (
	"os/exec"
)

type Execute struct {
	Flags string
}

func getGrpcurlCmd(flags string) *exec.Cmd {
	return exec.Command("/bin/sh", "-c", "grpcurl "+flags)
}

func NewExec() *Execute {
	return &Execute{}
}
