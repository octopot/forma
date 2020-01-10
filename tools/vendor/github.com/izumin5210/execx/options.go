package execx

import (
	"context"
	"os/exec"
	"time"
)

var (
	DefaultGracePeriod    time.Duration  = 30 * time.Second
	DefaultErrorLog       Logger         = new(nopLogger)
	DefaultNewProcessFunc NewProcessFunc = NewOSProcess
)

type Option func(*Config)

func defaultConfig() *Config {
	return &Config{
		GracePeriod:    DefaultGracePeriod,
		ErrorLog:       DefaultErrorLog,
		NewProcessFunc: DefaultNewProcessFunc,
	}
}

type Config struct {
	GracePeriod time.Duration

	NewProcessFunc NewProcessFunc

	ErrorLog Logger
}

func (c *Config) apply(opts []Option) {
	for _, f := range opts {
		f(c)
	}
}

func WithGracePeriod(d time.Duration) Option {
	return func(c *Config) { c.GracePeriod = d }
}

func WithNewProcessFunc(f NewProcessFunc) Option {
	return func(c *Config) { c.NewProcessFunc = f }
}

func WithFakeProcess(f func(context.Context, *exec.Cmd) error) Option {
	return WithNewProcessFunc(NewFakeNewProcessFunc(f))
}

func WithErrorLog(l Logger) Option {
	return func(c *Config) { c.ErrorLog = l }
}
