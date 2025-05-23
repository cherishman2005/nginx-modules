# ingress-controller

## 什么是 Ingress Controller？

Ingress Controller 是一种专用的负载均衡，用于管理进出 Kubernetes 集群的四层和七层流量。明确这一点后，我们来看看 NGINX 使用的术语（与业界所用术语基本相同）：

* 入向流量（Ingress traffic） – 进入 Kubernetes 集群的流量
* 出向流量（Egress traffic） – 离开 Kubernetes 集群的流量
* 南北向流量（North‑south traffic） – 进出 Kubernetes 集群的流量（也称为“入向到出向的流量”）
* 东西向流量（East‑west traffic） – 在 Kubernetes 集群内的服务之间移动的流量（也称为“从服务到服务的流量”）
* 服务网格（Service mesh） – 一种针对服务到服务的流量进行路由和保护的流量管理工具


## Nginx Ingress Controller实现原理

Nginx Ingress Controller是一个集控制平面和数据平面于一体的实现方案。每个Pod下有一个Controller进程，同时也包含Nginx相关进程。

![image](https://github.com/user-attachments/assets/07139883-23a5-4161-9473-3cbe10a7cfb6)

# 参考链接

- [Nginx Ingress运维相关知识点](https://www.alibabacloud.com/help/zh/ack/ack-managed-and-ack-dedicated/user-guide/nginx-ingress-operation-and-maintenance-related-knowledge-points)
