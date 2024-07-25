
# minikube addons enable ingress 获取到deploy.ymal

在 Minikube 中，addons 是预配置的服务和工具，可以通过 Minikube 命令行工具轻松启用或禁用。当你使用 minikube addons enable ingress 命令时，Minikube 会自动为你设置 Ingress 控制器，但它通常不会直接提供一个 deploy.yaml 文件，因为这个控制器是通过 Minikube 的内部机制来管理的。

但是，如果你想查看或获取与 Minikube 的 Ingress 控制器相关的 YAML 配置，你可以尝试以下方法：

查看 Minikube 的文档：

Minikube 的文档可能包含了关于如何启用 Ingress 控制器以及相关的 YAML 配置的信息。但是，这通常只会是概念性的，而不是具体的 YAML 文件。

检查 Minikube 生成的资源：

一旦你启用了 Ingress 控制器，你可以使用 kubectl 命令来查看 Minikube 创建的所有资源。例如，你可以使用以下命令来查看所有的 Deployments：

```bash
kubectl get deployments --all-namespaces
```
然后，你可以针对特定的 Deployment（比如与 Ingress 控制器相关的 Deployment）使用 kubectl describe 或 kubectl get -o yaml 命令来获取更详细的信息：

```bash
kubectl -n kube-system get deployment ingress-nginx-controller -o yaml
```
注意：这里的 ingress-nginx-controller 是示例名称，实际名称可能会有所不同，具体取决于你启用的 Ingress 控制器类型和版本。

从 Helm Chart 获取：

如果 Minikube 使用 Helm 来管理 Ingress 控制器（这是常见的做法），你可以从相关的 Helm Chart 中获取 YAML 文件。首先，你需要找到与你的 Minikube 版本和配置相匹配的 Helm Chart 版本。然后，你可以下载 Chart 并查看其中的 YAML 文件。但是，请注意，直接应用这些 YAML 文件可能不会与 Minikube 的内部机制完全兼容。

自定义 YAML：

如果你需要自定义 Ingress 控制器的配置，你可以从现有的 YAML 文件中获取灵感，并创建自己的 YAML 文件来部署 Ingress 控制器。这通常涉及从 Helm Chart 或其他来源获取基础配置，并根据你的需求进行修改。

总之，虽然 Minikube 不会直接提供一个与 minikube addons enable ingress 命令相关的 deploy.yaml 文件，但你可以通过查看 Minikube 创建的资源、从 Helm Chart 中获取灵感或创建自己的 YAML 文件来实现类似的功能。
