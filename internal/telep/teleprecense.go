package telep

import (
	"context"
	"fmt"

	"github.com/iapershin/teleconnect/internal/executor"
)

const (
	telepresenceBinary = "telepresence"
)

type ConnectOptions struct {
	Namespace string
	Cluster   string
	AlsoProxy string
}

func Connect(ctx context.Context, opts ConnectOptions) error {
	connectArgs := []string{
		"connect",
		"--context", opts.Cluster,
		"--mapped-namespaces", opts.Namespace,
		"-n", opts.Namespace,
		"--also-proxy", opts.AlsoProxy,
	}
	if err := executor.RunCommand(ctx, telepresenceBinary, connectArgs...); err != nil {
		return fmt.Errorf("failed to start telepresence session: %w", err)
	}
	fmt.Println("\033[1;32mTelepresence connected successfully\033[0m")
	return nil
}

func QuitSession(ctx context.Context) error {
	_ = executor.RunCommand(ctx, telepresenceBinary, "quit", "-s")
	return nil
}
