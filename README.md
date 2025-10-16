# 🛰️ teleconnect

A lightweight CLI utility for quickly establishing a [Telepresence](https://www.telepresence.io/) connection to a Kubernetes namespace.
It reads your current context form `kubeconfig` and connects `Telepresence` to the corresponding namespace — making local debugging and development against remote clusters effortless.

---

## 🚀 Features

* Automatically detects your current Kubernetes context and namespace
* Optionally allows specifying:

  * A different cluster or namespace
  * A custom `kubeconfig` file
  * Additional CIDR ranges for proxying traffic (`--also-proxy`)
* Gracefully handles existing Telepresence sessions before reconnecting

---

## ⚙️ Usage

```bash
teleconnect [flags]
```

### Flags

| Flag           | Short | Description                                   | Default                           |
| -------------- | ----- | --------------------------------------------- | --------------------------------- |
| `--kubeconfig` | `-k`  | Path to kubeconfig file                       | `$KUBECONFIG` or `~/.kube/config` |
| `--namespace`  | `-n`  | Kubernetes namespace                          | From current context              |
| `--cluster`    | `-c`  | Kubernetes cluster name                       | From current context              |
| `--also-proxy` |       | Additional CIDR to proxy through Telepresence | `10.0.0.0/8` *(default)*          |

Example:

```bash
teleconnect -n dev -c staging --also-proxy 192.168.0.0/16
```



