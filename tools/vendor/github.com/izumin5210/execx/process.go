package execx

import (
	"os"
	"os/exec"

	"github.com/Songmu/wrapcommander"
)

type Process interface {
	Start() error
	Wait() <-chan *ExitStatus
	Terminate() error
	Kill() error
	Signal() os.Signal
}

type NewProcessFunc func(*exec.Cmd) Process

type process struct {
	cmd *exec.Cmd
	ex  ExitStatus
}

func (p *process) Start() error {
	return p.cmd.Start()
}

func (p *process) Wait() <-chan *ExitStatus {
	// https://github.com/Songmu/timeout/blob/v0.4.0/timeout.go#L185-L191
	ch := make(chan *ExitStatus)
	go func() {
		defer close(ch)
		err := p.cmd.Wait()

		if err != nil {
			st, _ := wrapcommander.ErrorToWaitStatus(err)
			p.ex.Code = wrapcommander.WaitStatusToExitCode(st)
			p.ex.Signaled = st.Signaled()
			p.ex.Err = err
			ch <- &p.ex
		}
	}()
	return ch
}

func (p *process) Kill() error {
	return p.cmd.Process.Kill()
}
