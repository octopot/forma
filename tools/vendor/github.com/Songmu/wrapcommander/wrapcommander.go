package wrapcommander

import (
	"os"
	"os/exec"
	"syscall"
)

// exit statuses are same with GNU coreutils
const (
	ExitNormal            = 0
	ExitUnknownErr        = 125
	ExitCommandNotInvoked = 126
	ExitCommandNotFound   = 127
)

// IsPermission is alias of os.IsPermission for now
func IsPermission(err error) bool {
	return os.IsPermission(err)
}

// IsNotExist is alias of os.IsNotExist for now
func IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

// IsNotFoundInPATH returns a boolean indicating whether the error is known to
// report that a prog is not found in $PATH.
// ex. "exec: \"prog\": executable file not found in $PATH"
func IsNotFoundInPATH(err error) bool {
	e, ok := err.(*exec.Error)
	return ok && e.Err == exec.ErrNotFound
}

// IsExecFormatError returns a boolean indicating whether the error is known to
// report that format of an executable file is invalid.
// ex. "fork/exec ./prog: exec format error"
func IsExecFormatError(err error) bool {
	e, ok := err.(*os.PathError)
	return ok && e.Err == syscall.ENOEXEC
}

// IsInvoked returns a boolean indicating whether the error is known to report
// that the command is invoked or not.
func IsInvoked(err error) bool {
	if err == nil {
		return true
	}
	_, ok := err.(*exec.ExitError)
	return ok
}

// ErrorToWaitStatus try to convert error into syscall.WaitStatus
func ErrorToWaitStatus(err error) (syscall.WaitStatus, bool) {
	if e, ok := err.(*exec.ExitError); ok {
		st, ok := e.Sys().(syscall.WaitStatus)
		return st, ok
	}
	var zero syscall.WaitStatus
	return zero, false
}

// WaitStatusToExitCode converts WaitStatus to ExitCode
func WaitStatusToExitCode(st syscall.WaitStatus) int {
	return waitStatusToExitCode(st)
}

// ResolveExitCode retruns a int as command exit code from an error.
func ResolveExitCode(err error) int {
	if err == nil {
		return ExitNormal
	}
	if !IsInvoked(err) {
		switch {
		case IsPermission(err), IsExecFormatError(err):
			return ExitCommandNotInvoked
		case IsNotExist(err), IsNotFoundInPATH(err):
			return ExitCommandNotFound
		default:
			return ExitUnknownErr
		}
	}
	if status, ok := ErrorToWaitStatus(err); ok {
		return WaitStatusToExitCode(status)
	}
	return -1
}

// SeparateArgs separates command line arguments for wrapper command.
func SeparateArgs(args []string) ([]string, []string) {
	optsArgs := []string{}
	cmdArgs := []string{}
	for i, v := range args {
		if v == "--" && i+1 < len(args) {
			cmdArgs = args[i+1:]
			break
		}
		optsArgs = append(optsArgs, v)
	}
	if len(cmdArgs) <= 0 {
		cmdArgs, optsArgs = optsArgs, []string{}
	}
	return optsArgs, cmdArgs
}
