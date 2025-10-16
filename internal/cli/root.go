package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	defaultAlsoProxyCIDR = "10.10.0.0/14"
)

func Execute() error {
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
			return ConnectCmd(cmd.Context(), ConnectCmdOptions{
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

	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
