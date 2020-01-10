package execx

import (
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

func NewOSProcess(cmd *exec.Cmd) Process {
	// https://github.com/Songmu/timeout/blob/v0.4.0/timeout_windows.go#L9-L16
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			CreationFlags: syscall.CREATE_UNICODE_ENVIRONMENT | 0x00000200,
		}
	}
	return &process{
		cmd: cmd,
	}
}

func (p *process) Terminate() error {
	// https://github.com/mattn/goreman/blob/v0.3.4/proc_windows.go#L16-L55
	dll, err := windows.LoadDLL("kernel32.dll")
	if err != nil {
		return err
	}
	defer dll.Release()

	pid := p.cmd.Process.Pid

	f, err := dll.FindProc("AttachConsole")
	if err != nil {
		return err
	}
	r1, _, err := f.Call(uintptr(pid))
	if r1 == 0 && err != syscall.ERROR_ACCESS_DENIED {
		return err
	}

	f, err = dll.FindProc("SetConsoleCtrlHandler")
	if err != nil {
		return err
	}
	r1, _, err = f.Call(0, 1)
	if r1 == 0 {
		return err
	}
	f, err = dll.FindProc("GenerateConsoleCtrlEvent")
	if err != nil {
		return err
	}
	r1, _, err = f.Call(windows.CTRL_BREAK_EVENT, uintptr(pid))
	if r1 == 0 {
		return err
	}
	r1, _, err = f.Call(windows.CTRL_C_EVENT, uintptr(pid))
	if r1 == 0 {
		return err
	}
	return nil
}

func (p *process) Signal() os.Signal { return os.Interrupt }
