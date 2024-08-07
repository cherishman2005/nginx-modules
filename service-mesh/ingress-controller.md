# ingress-controller

## 什么是 Ingress Controller？

Ingress Controller 是一种专用的负载均衡，用于管理进出 Kubernetes 集群的四层和七层流量。明确这一点后，我们来看看 NGINX 使用的术语（与业界所用术语基本相同）：

* 入向流量（Ingress traffic） – 进入 Kubernetes 集群的流量
* 出向流量（Egress traffic） – 离开 Kubernetes 集群的流量
* 南北向流量（North‑south traffic） – 进出 Kubernetes 集群的流量（也称为“入向到出向的流量”）
* 东西向流量（East‑west traffic） – 在 Kubernetes 集群内的服务之间移动的流量（也称为“从服务到服务的流量”）
* 服务网格（Service mesh） – 一种针对服务到服务的流量进行路由和保护的流量管理工具
