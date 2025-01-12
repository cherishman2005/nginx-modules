# Service-Mesh与SDN的关系

Service Mesh 是下一代 SDN 的概念，但它并不是 SDN 的继任者。Service Mesh 和 SDN 都是网络编排技术，但它们的应用范围和实现方式有所不同。

Service Mesh 是一种专门用于微服务架构的网络编排技术，它将服务之间的通信管理和控制与服务的业务逻辑分离开来，使得服务的部署和管理更加灵活和简单。Service Mesh 的核心组件是代理（Proxy），它将服务之间的网络通信转发到代理上，代理再将请求转发到目标服务。这样可以实现服务之间的网络通信的监控、控制和管理。

而 SDN 是一种软件定义网络技术，它将网络控制层和数据层分离开来，使得网络的控制和数据传输可以灵活地进行编排和管理。SDN 的核心组件是控制器（Controller），它通过控制器来实现网络的编排和管理。

因此，Service Mesh 和 SDN 的应用范围和实现方式不同，它们并不是继任者和前任的关系，而是两种不同的网络编排技术。

![image](https://github.com/user-attachments/assets/947fe113-588e-472c-ae1d-69fcc4c63ec9)

SDN和Service Mesh出于网络协议中的不同层次

Service Mesh可以借鉴SDN的架构来解决微服务系统的服务通信的相关问题。

![image](https://github.com/user-attachments/assets/88df8369-6d26-4c2d-9e4c-e14da2acbe90)


# 参考链接

- [https://cloud.tencent.com/developer/article/2063485](https://cloud.tencent.com/developer/article/2063485)

