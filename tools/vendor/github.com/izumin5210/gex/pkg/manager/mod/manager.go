package mod

import (
	"context"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"

	"github.com/izumin5210/gex/pkg/manager"
)

const (
	// A space-separated list of -flag=value settings to apply
	// to go commands by default, when the given flag is known by
	// the current command. Each entry must be a standalone flag.
	// Because the entries are space-separated, flag values must
	// not contain spaces. Flags listed on the command line
	// are applied after this list and therefore override it.
	//
	// Details: https://tip.golang.org/cmd/go/#hdr-Environment_variables
	GoFlagsEnv = "GOFLAGS"
)

// NewManager creates a manager.Interface instance to build tools vendored with Modules.
func NewManager(executor manager.Executor) manager.Interface {
	return &managerImpl{
		executor: executor,
	}
}

type managerImpl struct {
	executor manager.Executor
}

func (m *managerImpl) Add(ctx context.Context, pkgs []string, verbose bool) error {
	args := []string{"get"}
	if verbose {
		args = append(args, "-v")
	}
	args = append(args, pkgs...)
	return errors.WithStack(m.executor.Exec(ctx, "go", args...))
}

func (m *managerImpl) Build(ctx context.Context, binPath, pkg string, verbose bool) error {
	args := []string{"build", "-o", binPath}
	if verbose {
		args = append(args, "-v")
	}
	args = append(args, pkg)
	return errors.WithStack(m.executor.Exec(ctx, "go", args...))
}

func (m *managerImpl) Sync(ctx context.Context, verbose bool) error {
	args := []string{"mod", "tidy"}
	if verbose {
		args = append(args, "-v")
	}
	if err := errors.WithStack(m.executor.Exec(ctx, "go", args...)); err != nil {
		return err
	}
	if m.Vendor() {
		args := []string{"mod", "vendor"}
		if verbose {
			args = append(args, "-v")
		}
		return errors.WithStack(m.executor.Exec(ctx, "go", args...))
	}
	return nil
}

func (m *managerImpl) Vendor() bool {
	const (
		vendor = "vendor"
	)
	var (
		mod  string
		set  = flag.NewFlagSet(GoFlagsEnv, flag.ContinueOnError)
		args = strings.Split(strings.TrimSpace(os.Getenv(set.Name())), " ")
	)
	set.SetOutput(ioutil.Discard)
	set.StringVar(&mod, "mod", "", "module download mode to use: readonly or vendor")
	_ = set.Parse(args)
	return mod == vendor
}
