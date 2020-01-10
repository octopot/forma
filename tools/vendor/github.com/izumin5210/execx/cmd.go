package execx

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"sync"
	"time"

	"github.com/Songmu/wrapcommander"
)

var (
	ErrUnimplemented = errors.New("unimplemented")
	ErrNotStarted    = errors.New("command not started")
)

type Cmd struct {
	*exec.Cmd
	*Config
	ctx context.Context
	p   Process
}

func (c *Cmd) Run() error {
	if err := c.Start(); err != nil {
		return err
	}
	return c.Wait()
}

func (c *Cmd) Start() error {
	c.p = c.NewProcessFunc(c.Cmd)
	if err := c.p.Start(); err != nil {
		return &ExitStatus{
			Code: wrapcommander.ResolveExitCode(err),
			Err:  err,
		}
	}
	return nil
}

func (c *Cmd) Wait() error {
	if c.p == nil {
		return ErrNotStarted
	}

	// https://github.com/Songmu/timeout/blob/v0.4.0/timeout.go#L132-L174
	killCh := make(chan struct{}, 2)

	done := make(chan struct{})
	defer close(done)

	exitCh := c.p.Wait()
	var killOnce, termOnce sync.Once
	var killed bool

	for {
		select {
		case ex, ok := <-exitCh:
			if ok {
				ex.Killed = killed
				ex.Timeout = c.ctx.Err() == context.DeadlineExceeded
				ex.Canceled = c.ctx.Err() == context.Canceled
				return ex
			}
			return nil

		case <-killCh:
			killOnce.Do(func() {
				c.handleError(c.p.Kill())
				killed = true
			})

		case <-c.ctx.Done():
			termOnce.Do(func() {
				c.handleError(c.p.Terminate())

				go func() {
					select {
					case <-done:
						return
					case <-time.After(c.GracePeriod):
						killCh <- struct{}{}
					}
				}()
			})
		}
	}
}

func (c *Cmd) CombinedOutput() ([]byte, error) {
	buf := new(bytes.Buffer)
	c.Stdout = buf
	c.Stderr = buf
	err := c.Run()
	return buf.Bytes(), err
}

func (c *Cmd) Output() ([]byte, error) {
	buf := new(bytes.Buffer)
	c.Stdout = buf
	err := c.Run()
	return buf.Bytes(), err
}

func (c *Cmd) handleError(err error) {
	if err != nil {
		c.ErrorLog.Print(err)
	}
}
