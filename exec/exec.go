package exec

import (
	"os/exec"
)

// creates a new exec.Cmd object to run a shell command.
func NewCMD(cmd string) *exec.Cmd { return exec.Command("/bin/sh", "-c", cmd) }
