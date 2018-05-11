// +build !windows

package flac

import (
	"context"
	"os/exec"
)

// Command creates command struct on unix systems
func Command(ctx context.Context, command string, args []string) *exec.Cmd {
	return exec.CommandContext(ctx, command, args...)
}
