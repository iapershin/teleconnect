package executor

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type RunCommandOptions struct {
	Silent bool
}

func RunCommand(ctx context.Context, name string, args []string, opts RunCommandOptions) error {
	cmd := exec.CommandContext(ctx, name, args...)
	if opts.Silent {
		return runCommandSilent(cmd)
	}
	return runCommand(cmd)
}

func runCommand(cmd *exec.Cmd) error {
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout: %w", err)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr: %w", err)
	}
	var stderrBuf bytes.Buffer

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	go streamOutput(stdoutPipe, os.Stdout)

	go func() {
		io.Copy(&stderrBuf, stderrPipe)
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command failed: %s", stderrBuf.String())
	}

	return nil
}

func runCommandSilent(cmd *exec.Cmd) error {
	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command failed: %s", stderrBuf.String())
	}

	return nil
}

func streamOutput(r io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Fprintln(out, scanner.Text())
	}
}
