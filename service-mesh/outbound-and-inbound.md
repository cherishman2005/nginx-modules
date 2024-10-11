# outbound-and-inbound

* Inbound与Outbound概念

![image](https://github.com/user-attachments/assets/833e6690-df3a-49a9-9c3b-1585038ee41d)

**什么是 bound**

* bound: 字面意为边界。有人译为站（名词）。而在现实的 k8s + istio 环境中，可以理解为 pod / service

* inbound: 有人译为入站。而在现实的 k8s + istio 环境中，可以理解为流量从 pod 外部进入 pod。即服务的被调用流量

* outbound: 有人译为出站。而在现实的 k8s + istio 环境中，可以理解为流量从 pod 内部输出到 pod 外部。

需要注意的是，对于同一个调用请求。他可以是调用者 service 的 outbound，同时也是被调用者的 inbound。如上图。
