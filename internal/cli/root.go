package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	defaultAlsoProxyCIDR = "10.0.0.0/8"
)

func Execute(ctx context.Context) error {
	var (
		namespace      string
		cluster        string
		kubeconfigPath string
		alsoProxy      string
	)

	var rootCmd = &cobra.Command{
		Use:   "teleconnect",
		Short: "Establish a Telepresence connection",
		Long: `Establish a Telepresence connection to a specified Kubernetes namespace.
If no namespace is specified, the current context's namespace from kubeconfig is used.
Optionally, you can switch kubectl contexts based on a provided suffix.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ConnectCmd(ctx, ConnectCmdOptions{
				KubeconfigPath: kubeconfigPath,
				Namespace:      namespace,
				Cluster:        cluster,
				AlsoProxy:      alsoProxy,
			})
		},
	}

	rootCmd.PersistentFlags().StringVarP(&kubeconfigPath, "kubeconfig", "k", "", "Path to the kubeconfig file (default: $KUBECONFIG or ~/.kube/config)")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Namespace (current context from kubeconfig by default)")
	rootCmd.PersistentFlags().StringVarP(&cluster, "cluster", "c", "", "Cluster name (current context from kubeconfig by default)")
	rootCmd.PersistentFlags().StringVarP(&alsoProxy, "also-proxy", "", "", fmt.Sprintf("Also proxy CIDR (default: %s)", defaultAlsoProxyCIDR))

	if alsoProxy == "" {
		alsoProxy = defaultAlsoProxyCIDR
	}

	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		return err
	}

	return nil
}
