package kube

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

const (
	defaultNamespace = "default"
)

type KubeSettings struct {
	Namespace string
	Cluster   string
}

func GetKubeSettings(kubeconfigPath, namespace, cluster string) (*KubeSettings, error) {
	config, err := loadKubeConfig(kubeconfigPath)
	if err != nil {
		return nil, err
	}

	currentContextName := config.CurrentContext
	kubeContext, ok := config.Contexts[currentContextName]
	if !ok {
		return nil, fmt.Errorf("context %q not found in kubeconfig", currentContextName)
	}

	if cluster == "" {
		cluster = kubeContext.Cluster
	}

	if _, ok := config.Clusters[cluster]; !ok {
		return nil, fmt.Errorf("cluster %q not found in kubeconfig", cluster)
	}

	if namespace == "" {
		namespace = kubeContext.Namespace
		if namespace == "" {
			namespace = defaultNamespace
		}
	}

	return &KubeSettings{
		Namespace: namespace,
		Cluster:   cluster,
	}, nil
}

func getKubeConfigPath(providedPath string) (string, error) {
	if providedPath != "" {
		absPath, err := filepath.Abs(providedPath)
		if err != nil {
			return "", fmt.Errorf("unable to determine absolute path for kubeconfig: %w", err)
		}
		return absPath, nil
	}
	return clientcmd.RecommendedHomeFile, nil
}

func loadKubeConfig(kubeconfigPath string) (*api.Config, error) {
	path, err := getKubeConfigPath(kubeconfigPath)
	if err != nil {
		return nil, err
	}
	config, err := clientcmd.LoadFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load kubeconfig: %w", err)
	}
	return config, nil
}
