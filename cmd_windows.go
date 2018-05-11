// +build windows

package flac

import (
	"context"
	"os/exec"
	"syscall"
)

// Command creates command struct on windows systems.
// NOTE: It suppresses the window (useful for embedding in gui applications)
func Command(ctx context.Context, command string, args []string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
