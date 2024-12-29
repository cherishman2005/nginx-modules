# go-micro

## go-micro的主要功能
* 服务发现：自动服务注册和名称解析。服务发现是微服务开发的核心。当服务A需要与服务B通话时，它需要该服务的位置。默认发现机制是多播DNS（mdns），一种零配置系统。您可以选择使用SWIM协议为p2p网络设置八卦，或者为弹性云原生设置设置consul

* 负载均衡：基于服务发现构建的客户端负载均衡。一旦我们获得了服务的任意数量实例的地址，我们现在需要一种方法来决定要路由到哪个节点。我们使用随机散列负载均衡来提供跨服务的均匀分布，并在出现问题时重试不同的节点

* 消息编码：基于内容类型的动态消息编码。客户端和服务器将使用编解码器和内容类型为您无缝编码和解码Go类型。可以编码任何种类的消息并从不同的客户端发送。客户端和服务器默认处理此问题。这包括默认的protobuf和json

* 请求/响应：基于RPC的请求/响应，支持双向流。我们提供了同步通信的抽象。对服务的请求将自动解决，负载平衡，拨号和流式传输。启用tls时，默认传输为http / 1.1或http2

* Async Messaging：PubSub是异步通信和事件驱动架构的一流公民。事件通知是微服务开发的核心模式。启用tls时，默认消息传递是点对点http / 1.1或http2

* 可插拔接口：Go Micro为每个分布式系统抽象使用Go接口，因此，这些接口是可插拔的，并允许Go Micro与运行时无关，可以插入任何基础技术

* 插件地址：https://github.com/micro/go-plugins

## go-micro通信流程

* Server监听客户端的调用，和Brocker推送过来的信息进行处理。并且Server端需要向Register注册自己的存在或消亡，这样Client才能知道自己的状态

* Register服务的注册的发现，Client端从Register中得到Server的信息，然后每次调用都根据算法选择一个的Server进行通信，当然通信是要经过编码/解码，选择传输协议等一系列过程的

* 如果有需要通知所有的Server端可以使用Brocker进行信息的推送，Brocker 信息队列进行信息的接收和发布

## go-micro核心接口

go-micro之所以可以高度订制和他的框架结构是分不开的，go-micro由8个关键的interface组成，每一个interface都可以根据自己的需求重新实现，这8个主要的inteface也构成了go-micro的框架结构

![image](https://github.com/user-attachments/assets/c25c2e18-b120-47d8-9ab4-bcd732932919)

# 参考链接

- [go-micro](https://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/GoMicro%E5%85%A5%E9%97%A8.html)
