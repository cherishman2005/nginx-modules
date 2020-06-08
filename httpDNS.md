# httpDNS 

## httpDNS比local DNS的优势

* 防劫持：由于绕过了运营商的LocalDNS，用户解析域名的请求通过Http协议直接透传到了HTTPDNS服务器IP上，用户在客户端的域名解析请求将不会遭受到域名解析被劫持的困扰。

* 精度调整: 由于运营商策略的多样性，其 Local DNS 的解析结果可能不是最近、最优的节点，HTTPDNS 能直接获取客户端 IP ，基于客户端 IP 获得最精准的解析结果，让客户端就近接入业务节点。

* 扩展性强: HTTPDNS提供可靠的域名解析服务，业务可将自有调度逻辑与HTTPDNS返回结果结合，实现更精细化的流量调度。

* 快速生效: 解决 “Local DNS 不遵循权威 TTL ，变更域名解析结果后全网无法即时生效”的问题。



## httpDNS解析示例

访问HTTPDNS服务时，一次请求只能解析一个域名。
请求示例：
示例1（默认来源IP）：http://203.107.1.33/100000/d?host=www.aliyun.com
示例2（指定来源IP）：http://203.107.1.33/100000/d?host=www.aliyun.com&ip=42.120.74.196
示例3（指定解析类型）：http://203.107.1.33/100000/d?host=www.aliyun.com&ip=219.242.0.1&query=4,6


```javascript
{
	"host": "www.aliyun.com",
	"ipsv6": ["2401:b180:1:60:0:0:0:6", "2401:b180:1:50:0:0:0:f"],
	"client_ip": "219.242.0.1",
	"ips": ["140.205.135.3"],
	"ttl": 298,
	"origin_ttl": 300
}
```

# FAQ

## https必须有域名才能申请吗？

不是的，申请SSL证书（HTTPS），需要域名或者公网IP才都是可以申请的。
目前支持公网IP或域名申请。 
