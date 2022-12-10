package exec

import (
	"os/exec"
)

type Execute struct {
	Flags string
	cmd   *exec.Cmd
}

func getGrpcurlCmd(flags string) *exec.Cmd {
	return exec.Command("/bin/sh", "-c", "grpcurl "+flags)
}

func getCmd(cmd string) *exec.Cmd {
	return exec.Command("/bin/sh", "-c", cmd)
}

func NewExec() *Execute {
	return &Execute{}
}
