package execx

import "context"

func Command(cmd string, args ...string) *Cmd {
	return New().Command(cmd, args...)
}

func CommandContext(ctx context.Context, cmd string, args ...string) *Cmd {
	return New().CommandContext(ctx, cmd, args...)
}
