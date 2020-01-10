// +build plan9 windows

package wrapcommander

import "syscall"

func waitStatusToExitCode(w syscall.WaitStatus) int {
	return w.ExitStatus()
}
