# istio

Istio 扩展了 Kubernetes，以建立可编程、应用程序感知的网络。Istio 可与 Kubernetes 和传统工作负载配合使用，为复杂的部署带来标准、通用的`流量管理`、`可观测`和`安全性`。

选择所需的功能，Istio 会根据需要部署代理基础设施。使用零信任隧道实现四层性能和安全性，或添加强大的 Envoy 服务代理实现七层功能。

* 服务网格

使用 Istio 轻松、安全、可靠地构建云原生工作负载，带或不带 Sidecar。

## 流量管理

Istio 的流量路由规则可以让您很容易的控制服务之间的流量和 API 调用。 Istio 简化了服务级别属性的配置，比如熔断器、超时和重试，并且能轻松的设置重要的任务， 如 A/B 测试、金丝雀发布、基于流量百分比切分的分阶段发布等。它还提供了开箱即用的故障恢复特性， 有助于增强应用的健壮性，从而更好地应对被依赖的服务或网络发生故障的情况。

Istio 的流量管理模型源于和服务一起部署的 Envoy 代理。 网格内服务发送和接收的所有 data plane 流量都经由 Envoy 代理， 这让控制网格内的流量变得异常简单，而且不需要对服务做任何的更改。


# 参考链接

- [https://istio.io/latest/zh/docs/concepts/traffic-management/](https://istio.io/latest/zh/docs/concepts/traffic-management/)
