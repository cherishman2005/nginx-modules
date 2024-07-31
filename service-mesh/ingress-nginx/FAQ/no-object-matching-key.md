# Error obtaining Endpoints for Service "default/whoami": no object matching key "default/whoami" in local store

这个错误信息表明 Kubernetes 集群在尝试获取名为 "whoami" 的服务（Service）的端点（Endpoints）时失败了，因为在本地存储中没有找到与键 "default/whoami" 匹配的对象。这通常意味着以下几个可能的问题：

服务不存在：在 "default" 命名空间中，名为 "whoami" 的服务可能没有被创建。你需要检查是否已经创建了该服务。可以使用以下命令来查看 "default" 命名空间中的所有服务：

```bash
kubectl get services --namespace=default
```
或者，如果你正在使用默认的命名空间（即没有指定命名空间），可以直接使用：

```bash
kubectl get services
```
命名空间错误：服务可能存在于不同的命名空间中。确保你查找的命名空间是正确的。如果你不确定服务在哪个命名空间，可以使用以下命令列出所有命名空间中的服务：

```bash
kubectl get services --all-namespaces
```
权限问题：如果你没有足够的权限来查看服务，也可能导致这个错误。确保你的 Kubernetes 角色和角色绑定允许你访问该服务。

集群状态问题：在极少数情况下，集群的状态可能不一致或出现问题，导致某些资源无法正确显示。尝试重新启动你的 kubectl 客户端或检查集群的健康状态。

网络问题：虽然这不太可能是直接导致这个错误的原因，但网络问题可能会阻止 kubectl 正确与 Kubernetes API 服务器通信。

如果确认服务确实存在但仍然出现这个错误，可以尝试以下步骤来进一步诊断问题：

查看服务定义：使用 kubectl describe service whoami --namespace=default 来查看服务的详细信息和状态。

检查端点：使用 kubectl get endpoints whoami --namespace=default 来查看服务的端点信息。如果服务没有关联的端点，可能是因为服务选择的后端 Pod 没有正确运行或标签不匹配。

检查 Pod：查看与 "whoami" 服务相关联的 Pod 是否正在运行且标签正确。使用 kubectl get pods --selector=app=whoami --namespace=default（假设你的 Pod 使用了 app=whoami 这样的标签）。

如果以上步骤都不能解决问题，可能需要更深入地检查 Kubernetes 集群的配置和日志。
