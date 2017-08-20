package utils

import (
	"os/exec"
)

// ExecCmd exec command
func ExecCmd(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
