# nginx-modules

## nginx-consul完美组合：

	（1）服务发现API网关；
	（2）后端微服务架构：负载均衡；
	
## nginx
	（1）http/https API网关：限流，黑白名单，接入控制；—— 成熟的lua插件，C module；
	（2）http2;
	（3）http CDN代理缓存；
	（4）upsteram静态负载均衡；（4层，7层负载均衡）；
	（5）c module插件开发，做点播/直播的转码转封装框架；

## consul
	（1）microservice服务发现；
	（2）key/value配置中心；
	（3）health健康检测；
	（4）与nginx结合，实现（4层/7层）动态负载均衡；
