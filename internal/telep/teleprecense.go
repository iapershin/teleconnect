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
	err := executor.RunCommand(ctx,
		telepresenceBinary,
		connectArgs,
		executor.RunCommandOptions{},
	)
	if err != nil {
		return fmt.Errorf("failed to start telepresence session: %w", err)
	}
	log.LogSuccess("Telepresence connected successfully")
	return nil
}

func QuitSession(ctx context.Context) error {
	_ = executor.RunCommand(ctx,
		telepresenceBinary,
		[]string{"quit", "-s"},
		executor.RunCommandOptions{Silent: true},
	)
	return nil
}

func Version(ctx context.Context) error {
	err := executor.RunCommand(ctx,
		telepresenceBinary,
		[]string{"version"},
		executor.RunCommandOptions{Silent: true},
	)
	if err != nil {
		return fmt.Errorf("unable to find telepresence binary: %w\n install telepresence or check your PATH", err)
	}
	return nil
}
