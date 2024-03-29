# 什么是 API 网关？

API 网关 是与实际后端服务或数据连接的流量管理器，并针对 API 调用应用策略、身份验证和一般访问控制来保护有价值的数据。API 网关是您控制对后端系统和服务的访问的方式，并且它旨在优化外部客户端和后端服务之间的通信，为您的客户提供无缝体验。API 网关可确保服务的可扩展性和高可用性。它负责将请求路由到相应的服务，然后向请求者发送回复。API 网关在您的数据和 API 之间维护安全连接，并管理您公司内外的 API 流量和请求，包括负载均衡。网关针对 API 调用应用策略、身份验证和一般访问控制，以保护有价值的数据。API 网关接收来自客户端的所有 API 调用，并利用请求路由、合成和协议转换将它们路由到正确的微服务。

![image](https://user-images.githubusercontent.com/17688273/227445770-cc45f485-41c8-460c-84fb-e48beb648f71.png)

使用 API 网关的主要原因之一是因为它能够调用多个后端服务并聚合结果。客户不必为每项服务发送请求，而是可以将请求发送到 API 网关，然后由 API 网关将请求传递给相关服务。此外，API 网关还可以替代 “一刀切” 风格的 API。API 网关还可以为每个客户端公开不同的 API，这在当今不断变化的环境中是必不可少的。


## 为什么要使用 API 网关？

今天，大多数企业 API 都是使用网关进行部署的。由于微服务使用的增加，API 网关的使用也越来越多。微服务 允许将应用程序解构为几个松散耦合的服务，因为每个微服务都需要自己的功能。微服务使开发、部署和维护应用程序的不同功能变得更加容易，但它们也会使客户更难以快速、安全地访问应用程序。API 网关是此问题的解决方案。网关不是让客户单独请求访问每项微服务，而是请求的单一入口点，它将请求分配给相应的服务，收集结果并将其传递给请求者。使用 API 网关的主要原因，此函数被称为开发人员的路由。例如，API 网关可帮助贵公司管理来自诸如优步等移动应用程序和谷歌地图等后端应用程序的调用所产生的流量。

API 网关对 成功的 API 管理 至关重要。作为连接您的客户与服务的主要代理，该网关支持重要的管理和安全功能，包括身份验证、指标收集、输入验证和响应转换。

### 身份验证
API 网关可用于对 API 调用进行身份验证。这样，即使客户需要从多个服务访问数据，他们也只需在网关上进行一次身份验证。这可以减少延迟，并确保身份验证过程在整个应用程序中保持一致。与使用护照验证您的身份或签证证明您可以在特定国家工作的方式相似，API 网关为消费者提供了多种身份验证和访问 API 资源的方式。网关可以使用众多开放标准之一来确定消费者的身份或有效性（即 OAuth、JWT 令牌、API 密匙、HTTP BasIC/ 摘要、SAML 等），也可以使用非标准方法在消息标头或有效负载中查找凭据。API 网关还可以调用其他系统来验证身份，就像警方可以访问犯罪数据库一样。此外，就像机场的海关一样，API 网关还可以检查传入的 API 消费者中的威胁。他们可以使用 API 防火墙、内容验证和消息完整性检查，包括确定 API 是否被篡改。API 网关还可以将传入 API 的风险评估委派给第三方应用程序，以便作出决定。

### 指标收集
由于所有请求都通过 API 网关，因此它是收集数据分析的理想场所。例如，API 网关可以测量用户正在发出的请求数量或向每个微服务转发的请求数量。API 网关也可用于限制请求。如果用户发送的请求过多，则可以对网关进行编程以拒绝请求，而不是将它们传递给其中一个微服务。

### 输入验证
输入验证是 API 网关，确保所有客户请求都包含完成请求所需的信息，并以正确的格式提供。如果出现差错，网关将拒绝该请求。如果包含所有必要信息，网关将请求路由到负责检索请求信息的微服务。

### 响应转换
响应转换是 API 网关的重要功能。它充当信息的 “转换者”。例如，如果您的后端服务放弃了 XML 中的答案，但请求者需要在 JSON 中使用，则网关将自动处理此问题。不同的应用程序和用户通常需要访问不同的信息。例如，移动应用程序需要的数据通常少于其网络应用程序，因此网关可以对请求提供正确的响应。来自内部用户的请求可能会在响应中包含更多数据。其中一些数据需要安全保护，然后才能响应外部用户的类似请求，这是网关的工作。

## API 网关的优势

将微服务打包在一起并通过 API 网关访问，可以安全、更快、更轻松地访问您的服务。使用 API 网关为数字商务提供了无数额外的优势，包括：

* 通过单一界面方法使您的 API 和后端系统更加安全
* 使用安全和访问控制、限制、路由、中介和 SLA 管理的可扩展策略，让您完全控制 API 执行环境。
* 为您的服务和应用程序用户编写更简单的代码
* 由于来回电话减少，随时间推移，延迟也减少
* 更快、更轻松地访问所有微服务
* 减少每个微服务或负载平衡的工作负载
* 全面的指标收集

API 网关的其他优势是：它隐藏了应用程序与请求者或客户端的分区方式，客户端不再需要知道所有个体服务的位置，并且无论使用何种代码，它都能为每个请求提供最佳 API。

应用程序的成功可能取决于功能强大的 API 网关。网关可确保服务的出色性能、高可用性和可扩展性。
