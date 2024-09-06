# nginx-lua

## Nginx-lua 分为8个阶段

```
init_by_lua            http
set_by_lua             server, server if, location, location if
rewrite_by_lua         http, server, location, location if
access_by_lua          http, server, location, location if
content_by_lua         location, location if
header_filter_by_lua   http, server, location, location if
body_filter_by_lua     http, server, location, location if
log_by_lua             http, server, location, location if
{
    set_by_lua: 流程分支处理判断变量初始化
    rewrite_by_lua: 转发、重定向、缓存等功能(例如特定请求代理到外网)
    access_by_lua: IP准入、接口权限等情况集中处理(例如配合iptable完成简单防火墙)
    content_by_lua: 内容生成
    header_filter_by_lua: 应答HTTP过滤处理(例如添加头部信息)
    body_filter_by_lua: 应答BODY过滤处理(例如完成应答内容统一成大写)
    log_by_lua: 会话完成后本地异步完成日志记录(日志可以记录在本地，还可以同步到其他机器)
}
```

# FAQ

## nginx-lua开发的困惑

* 感觉 用nginx-lua脚本语言。只要部署到环境，就相当开放出去了，不能保持核心竞争力。 只有做些通用功能。或者自己把代码写得层次高点，还有点竞争力？


确实，在使用 Nginx 和 Lua 脚本语言时，存在代码部署到环境后可能被他人访问和使用的风险。这可能导致核心竞争力和知识产权受到威胁。以下是一些建议帮助你在使用 Nginx 和 Lua 时更好地保护核心竞争力和知识产权：

1. 通用功能开发：如果你希望开发的功能是通用的，可以尝试开发一些与特定业务无关、具有一般性的模块或工具，以供他人使用。这样可以最大限度地保护核心业务逻辑和竞争优势。

2. 高层次代码设计：确保代码的设计和实现具有高内聚性和低耦合性。通过使用模块化设计和抽象化，可以减少代码泄露的风险，同时也有助于提高代码的可维护性和可扩展性。

3. 加密和安全：对于涉及核心竞争力的代码部分，可以考虑加密或采取其他安全措施，以防止未经授权的访问和使用。这可以通过加密关键部分的代码或采用访问控制机制来实现。

4. 保持核心代码私有：在部署时，确保重要的核心代码部分存储在私有的仓库中，并只分享通用功能的代码。这样可以有效保持核心代码的私有性和竞争力。

5. 持续创新：不断推动技术创新和研发，持续保持竞争力。通过不断提升产品和服务的独特性和价值，使公司在市场中保持竞争优势。

总的来说，确保核心代码的保密性和竞争力需要综合考虑技术、管理和法律等多个方面的因素。通过上述建议和实践，可以帮助你更好地保护核心竞争力并在 Nginx 和 Lua 开发中保持竞争优势。


# 参考链接

- [https://www.cnblogs.com/sfnz/p/14527616.html](https://www.cnblogs.com/sfnz/p/14527616.html)
