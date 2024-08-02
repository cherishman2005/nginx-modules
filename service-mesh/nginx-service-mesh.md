# nginx service-mesh

为 Kubernetes 服务提供企业级的可用性、安全防护以及可视化

在 Kubernetes 集群中减少复杂性，延长正常运行时间，并大规模且更好地获取关于 service 健康状况和性能状况的洞察信息。

![image](https://github.com/user-attachments/assets/f05c6032-4db2-4bdf-8591-d92eb6f6a8c5)


## 为什么要使用 NGINX Service Mesh？

### 降低复杂性

![image](https://github.com/user-attachments/assets/0809927e-8818-4031-9644-cc56c5f94edf)

* 简化运维

NGINX Service Mesh 专为提供一致、安全、低延迟的 service 连接而构建：

一个易于部署和使用的 Kubernetes 原生工具，用于在 Kubernetes 集群内集中控制和管理 service 之间的通信。

横跨所有混合云和多云环境，使用相同的数据平面和控制平面。

与 NGINX Ingress Controller 紧密且无缝集成，实现应用在集群内部、外部和之间的统一连接。

### 增加正常运行时间

![image](https://github.com/user-attachments/assets/01b650ab-f788-4158-bae3-ae80bd308c3e)

* 无间断地交付应用

NGINX Service Mesh 通过以下方式确保关键业务后端服务的可用性：

* 借助动态更新目标服务实例实现高级七层（HTTP、gRPC）和四层（TCP、UDP）负载均衡
* 借助灰度发布和蓝绿发布，避免在部署后端服务的新版本时出现宕机
* 通过事务重试功能，实现速率限制和断路器连接模式，以防止在拓扑变化、极高的请求速率以及服务故障期间发生连接超时和错误

### 强化防护

![image](https://github.com/user-attachments/assets/8030bc9a-604e-404e-a543-d8574fca9959)

* 缓解网络安全威胁

在 Kubernetes 集群中集成强大的集中式安全防护控制，覆盖分布式服务：
* 使用内置或外部证书颁发机构（CA）以及自动化证书生命周期管理，借助 SPIFFE 和 SPIRE 运行时，通过已经验证的服务身份来管理和强制执行授权访问
* 应用访问控制策略来控制与特定源和目标端点之间的通信
* 通过 mTLS（双向传输层安全性）认证和加密，确保 service 之间通信的安全性

### 提升可视化

![image](https://github.com/user-attachments/assets/4ed845ea-1cb7-401c-9cc5-734e58c11d7b)

* 获取更佳的见解

通过 NGINX Service Mesh 监控和报告的 100 多个细粒度指标，实现更好的针对应用健康状态和性能状态的可视化：

* 在问题影响到客户之前发现问题，减少宕机和停机时间
* 快速找到服务故障的根本原因，简化和优化故障排除流程
* 通过与您喜欢的生态系统工具相集成，实现收集、监控和分析指标，包括OpenTelemetry、Grafana、Prometheus 和 Jaeger 等工具

### 更快地发布应用

* 加快发布速度

借助 NGINX Service Mesh 更快、更轻松地交付应用：
* 在复杂的分布式微服务环境中，简化和优化 service 之间的通信
* 集中精力在后端 service 中实现核心业务能力和功能
* 将安全防护和其他非功能性要求卸载到平台层

# 参考链接

- [https://www.nginx-cn.net/products/nginx-service-mesh/](https://www.nginx-cn.net/products/nginx-service-mesh/)
