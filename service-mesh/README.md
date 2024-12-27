# service-mesh

技术选型

- [技术选型](./技术选型.md)

* serivce mesh架构图

![serivce-mesh架构图](./serivce-mesh-control-plane.png)

**服务网格工作流程**

1. 控制平面将整个网格中的服务配置推送到所有节点的 sidecar 代理中。
2. Sidecar 代理将服务请求路由到目的地址，根据中的参数判断是到生产环境、测试环境还是 staging 环境中的服务（服务可能同时部署在这三个环境中），是路由到本地环境还是公有云环境？所有的这些路由信息可以动态配置，可以是全局配置也可以为某些服务单独配置。
3. 当 sidecar 确认了目的地址后，将流量发送到相应服务发现端点，在 Kubernetes 中是 service，然后 service 会将服务转发给后端的实例。
4. Sidecar 根据它观测到最近请求的延迟时间，选择出所有应用程序的实例中响应最快的实例。
5. Sidecar 将请求发送给该实例，同时记录响应类型和延迟数据。
6. 如果该实例挂了、不响应了或者进程不工作了，sidecar 将把请求发送到其他实例上重试。
7. 如果该实例持续返回 error，sidecar 会将该实例从负载均衡池中移除，稍后再周期性得重试。
8. 如果请求的截止时间已过，sidecar 主动失败该请求，而不是再次尝试添加负载。
9. Sidecar 以 metric 和分布式追踪的形式捕获上述行为的各个方面，这些追踪信息将发送到集中 metric 系统。

`服务网格并没有给我们带来新功能，它是用于解决其他工具已经解决过的问题，只不过这次是在云原生的 Kubernetes 环境下的实现。`

## nginx

### ingress-nginx

![image](https://github.com/user-attachments/assets/9d29d1fa-5315-4afc-a340-8a85b312b4b6)


## apisix

![image](https://github.com/user-attachments/assets/4b37d6a1-8b68-408a-9fde-2944e4f8a37c)


### 服务网格

未来五到十年，基于云原生模式架构下的服务网格架构开始崭露头角。APISIX 也提前开始锁定赛道，通过调研和技术分析后，APISIX 已经支持了 xDS 协议，APISIX Mesh 就此诞生，在服务网格领域 APISIX 也拥有了一席之地。

![image](https://github.com/user-attachments/assets/df373852-6cf8-4a9d-b206-d4b6e852216d)

* sidecar边车模式

![image](https://github.com/user-attachments/assets/f9c22793-2dcf-44cc-8ff4-3d3656d7e22d)


## bfe

- [bfe](./bfe.md)

# 小结

* 在某些技术领域，不停炒作概念，反而技术还是原来的技术。-- 旧酒换新瓶。
  * 不要被表象迷惑。

# 参考链接

- [Nginx Ingress、ALB Ingress和MSE Ingress对比](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/comparison-among-nginx-ingresses-alb-ingresses-and-mse-ingresses-1?spm=a2c4g.11186623.0.0.2b571a01qX91wa)

- [https://www.taloflow.ai/guides/comparisons/apisix-vs-tyk-api-gateway](https://www.taloflow.ai/guides/comparisons/apisix-vs-tyk-api-gateway)

- [https://jimmysong.io/kubernetes-handbook/usecases/service-mesh.html](https://jimmysong.io/kubernetes-handbook/usecases/service-mesh.html)
