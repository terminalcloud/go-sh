// Package sh provides tools for shelling out commands
//
//
package sh

import (
	"os/exec"
	"syscall"
)

// WaitCode gets the exit code for a running command
//
// Normally returns exit code, nil
//
// Error Conditions:
// WaitCode takes err, so that it can be used on the result of Run{Cmd/Bash}.
// All failures to run the process manifest as an exit code of -1, with additional error info
func WaitCode(c *exec.Cmd, e error) (int, error) {
	if c == nil {
		return -1, e
	}
	if err := c.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode := exitError.Sys().(syscall.WaitStatus).ExitStatus()
			return exitCode, nil
		} else {
			return -1, err
		}
	}
	return 0, nil
}

func RunBash(bash string) (*exec.Cmd, error) {
	c := exec.Command("bash", "-c", bash)
	if err := c.Start(); err != nil {
		return nil, err
	}
	return c, nil
}

func RunCmd(prog string, args ...string) (*exec.Cmd, error) {
	c := exec.Command(prog, args...)
	if err := c.Start(); err != nil {
		return nil, err
	}
	return c, nil
}
