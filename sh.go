// Package sh provides tools for shelling out commands
//
//
package sh

import (
	"os/exec"
	"syscall"
	"time"
)

// Error type returned by WaitCodeTimeout in the event a process timed out
type TimeoutError struct {
	Killed bool
}

func (e TimeoutError) Error() string {
	if e.Killed {
		return "process timed out"
	} else {
		return "process timed out; failed to kill process"
	}
}

// getCode gets the exit code from the returned command of (*exec.Cmd).Wait()
//
// If no error is present, returns 0, nil
// If an exit code is present, returns code, nil
// If no exit code is present, returns -1, original error
func getCode(e error) (int, error) {
	if e != nil {
		if exitError, ok := e.(*exec.ExitError); ok {
			exitCode := exitError.Sys().(syscall.WaitStatus).ExitStatus()
			return exitCode, nil
		} else {
			return -1, e
		}
	}

	return 0, nil
}

// Get a new *exec.Cmd to run a bash string
func NewBash(bash string) *exec.Cmd {
	return exec.Command("bash", "-c", bash)
}

// Get a new *exec.Cmd to run a program
func NewCmd(prog string, args ...string) *exec.Cmd {
	return exec.Command(prog, args...)
}

// Create and immediately run a bash string
func RunBash(bash string) (*exec.Cmd, error) {
	c := NewBash(bash)
	if err := c.Start(); err != nil {
		return nil, err
	}
	return c, nil
}

// Immediately run a program
func RunCmd(prog string, args ...string) (*exec.Cmd, error) {
	c := NewCmd(prog, args...)
	if err := c.Start(); err != nil {
		return nil, err
	}
	return c, nil
}

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

	return getCode(c.Wait())
}

// WaitCodeTimeout gets the exit code for a running command, attempting to kill the command after timeout
//
// If the command times out, it will return TimeoutError
// Otherwise, behaves as WaitCode
func WaitCodeTimeout(c *exec.Cmd, e error, timeout time.Duration) (int, error) {
	if c == nil {
		return -1, e
	}

	done := make(chan error, 1)
	go func() { done <- c.Wait() }()

	select {
	case <-time.After(timeout):
		timeoutErr := TimeoutError{false}
		if err := c.Process.Kill(); err == nil {
			timeoutErr.Killed = true
		}

		return -1, timeoutErr
	case err := <-done:
		return getCode(err)
	}
}
