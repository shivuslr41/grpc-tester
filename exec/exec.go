package exec

import (
	"os/exec"
)

func NewCMD(cmd string) *exec.Cmd { return exec.Command("/bin/sh", "-c", cmd) }
