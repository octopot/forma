package execx

import (
	"context"
	"os"
	"os/exec"
)

func NewFakeNewProcessFunc(f func(ctx context.Context, cmd *exec.Cmd) error) NewProcessFunc {
	return func(cmd *exec.Cmd) Process {
		ctx, cancel := context.WithCancel(context.Background())
		errCh := make(chan error)

		return &FakeProcess{RunFunc: f, cmd: cmd, ctx: ctx, cancel: cancel, errCh: errCh}
	}
}

type FakeProcess struct {
	RunFunc func(ctx context.Context, cmd *exec.Cmd) error

	ctx    context.Context
	cancel func()
	cmd    *exec.Cmd
	errCh  chan error
}

func (p *FakeProcess) Start() error {
	go func() {
		defer p.cancel()
		defer close(p.errCh)
		err := p.RunFunc(p.ctx, p.cmd)
		if err != nil {
			p.errCh <- err
		}
	}()

	return nil
}

func (p *FakeProcess) Wait() <-chan *ExitStatus {
	ch := make(chan *ExitStatus)
	go func() {
		defer close(ch)
		if err := <-p.errCh; err != nil {
			ex := new(ExitStatus)
			ex.Code = 1
			ex.Err = err
			ch <- ex
		}
	}()
	return ch
}

func (p *FakeProcess) Terminate() error {
	p.cancel()
	return nil
}

func (p *FakeProcess) Kill() error {
	p.cancel()
	return nil
}

func (p *FakeProcess) Signal() os.Signal { return os.Interrupt }
