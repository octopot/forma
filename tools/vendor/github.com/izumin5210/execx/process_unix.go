// +build !windows

package execx

import (
	"os"
	"os/exec"
	"syscall"
)

func NewOSProcess(cmd *exec.Cmd) Process {
	// https://github.com/Songmu/timeout/blob/v0.4.0/timeout_unix.go#L14-L19
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}
	return &process{
		cmd: cmd,
	}
}

func (p *process) Terminate() error {
	// https://github.com/Songmu/timeout/blob/v0.4.0/timeout_unix.go#L21-L35
	sig := p.Signal()
	syssig, ok := sig.(syscall.Signal)
	if !ok {
		return p.cmd.Process.Signal(sig)
	}
	err := syscall.Kill(-p.cmd.Process.Pid, syssig)
	if err != nil {
		return err
	}
	if syssig != syscall.SIGKILL && syssig != syscall.SIGCONT {
		return syscall.Kill(-p.cmd.Process.Pid, syscall.SIGCONT)
	}
	return nil
}

func (p *process) Signal() os.Signal { return syscall.SIGTERM }
