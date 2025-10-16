package kube

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func writeTempKubeconfig(t *testing.T, config *api.Config) string {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "kubeconfig-*.yaml")
	assert.NoError(t, err)

	err = clientcmd.WriteToFile(*config, tmpFile.Name())
	assert.NoError(t, err)

	return tmpFile.Name()
}

func TestGetKubeSettings_WithDefaults(t *testing.T) {
	config := &api.Config{
		CurrentContext: "test-context",
		Contexts: map[string]*api.Context{
			"test-context": {
				Cluster:   "test-cluster",
				Namespace: "test-namespace",
			},
		},
		Clusters: map[string]*api.Cluster{
			"test-cluster": {
				Server: "https://dummy",
			},
		},
	}

	kubeconfigPath := writeTempKubeconfig(t, config)
	defer os.Remove(kubeconfigPath)

	settings, err := GetKubeSettings(kubeconfigPath, "", "")
	assert.NoError(t, err)
	assert.Equal(t, "test-namespace", settings.Namespace)
	assert.Equal(t, "test-cluster", settings.Cluster)
}

func TestGetKubeSettings_WithOverrides(t *testing.T) {
	config := &api.Config{
		CurrentContext: "other-context",
		Contexts: map[string]*api.Context{
			"other-context": {
				Cluster:   "other-cluster",
				Namespace: "other-namespace",
			},
		},
		Clusters: map[string]*api.Cluster{
			"other-cluster": {
				Server: "https://dummy",
			},
			"override-cluster": {
				Server: "https://override",
			},
		},
	}

	kubeconfigPath := writeTempKubeconfig(t, config)
	defer os.Remove(kubeconfigPath)

	settings, err := GetKubeSettings(kubeconfigPath, "override-namespace", "override-cluster")
	assert.NoError(t, err)
	assert.Equal(t, "override-namespace", settings.Namespace)
	assert.Equal(t, "override-cluster", settings.Cluster)
}

func TestGetKubeSettings_MissingCurrentContext(t *testing.T) {
	config := &api.Config{
		CurrentContext: "nonexistent-context",
		Contexts:       map[string]*api.Context{},
		Clusters:       map[string]*api.Cluster{},
	}

	kubeconfigPath := writeTempKubeconfig(t, config)
	defer os.Remove(kubeconfigPath)

	_, err := GetKubeSettings(kubeconfigPath, "", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "context")
}

func TestGetKubeSettings_ClusterNotFound(t *testing.T) {
	config := &api.Config{
		CurrentContext: "test-context",
		Contexts: map[string]*api.Context{
			"test-context": {
				Cluster:   "missing-cluster",
				Namespace: "some-namespace",
			},
		},
		Clusters: map[string]*api.Cluster{
			// missing-cluster intentionally not defined
		},
	}

	kubeconfigPath := writeTempKubeconfig(t, config)
	defer os.Remove(kubeconfigPath)

	_, err := GetKubeSettings(kubeconfigPath, "", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cluster")
}

func TestGetKubeSettings_DefaultNamespaceUsed(t *testing.T) {
	config := &api.Config{
		CurrentContext: "ctx",
		Contexts: map[string]*api.Context{
			"ctx": {
				Cluster: "cluster1",
				// Namespace intentionally left empty
			},
		},
		Clusters: map[string]*api.Cluster{
			"cluster1": {
				Server: "https://dummy",
			},
		},
	}

	kubeconfigPath := writeTempKubeconfig(t, config)
	defer os.Remove(kubeconfigPath)

	settings, err := GetKubeSettings(kubeconfigPath, "", "")
	assert.NoError(t, err)
	assert.Equal(t, "default", settings.Namespace)
}
