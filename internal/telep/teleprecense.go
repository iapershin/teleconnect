package telep

import (
	"context"
	"fmt"

	"github.com/iapershin/teleconnect/internal/executor"
	"github.com/iapershin/teleconnect/internal/log"
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
	log.LogSuccess("Telepresence connected successfully")
	return nil
}

func QuitSession(ctx context.Context) error {
	_ = executor.RunCommand(ctx, telepresenceBinary, "quit", "-s")
	return nil
}
