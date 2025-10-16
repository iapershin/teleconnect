package cli

import (
	"context"

	"github.com/iapershin/teleconnect/internal/kube"
	"github.com/iapershin/teleconnect/internal/telep"
)

type ConnectCmdOptions struct {
	KubeconfigPath string
	Namespace      string
	Cluster        string
	AlsoProxy      string
}

func ConnectCmd(ctx context.Context, opts ConnectCmdOptions) error {
	kubeSettings, err := kube.GetKubeSettings(opts.KubeconfigPath, opts.Namespace, opts.Cluster)
	if err != nil {
		return err
	}

	telep.QuitSession(ctx)

	return telep.Connect(ctx, telep.ConnectOptions{
		Namespace: kubeSettings.Namespace,
		Cluster:   kubeSettings.Cluster,
		AlsoProxy: opts.AlsoProxy,
	})
}
