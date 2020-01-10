// +build !plan9,!windows

package wrapcommander

import "syscall"

func waitStatusToExitCode(w syscall.WaitStatus) int {
	if w.Signaled() {
		return int(w.Signal()) + 128
	}
	return w.ExitStatus()
}
