# koa2与express比较

* express包含了router；

* koa2更加轻量级，通过middleware中间件插入koa-router功能。

* Koa专注于核心中间件功能，设计显式地利用了async/await使异步代码可读性更高。

|        |   koa(Router = require('koa-router'))  |  express(假设不使用app.get之类的方法) |
| :---------- | :------ | :------ |
| 初始化      | const app = new koa() | const app = express() |
| 实例化路由   | const router = Router() | const router = express.Router() |
| app级别的中间件   | app.use | app.use |
| 路由级别的中间件   | router.get | router.get |
| 路由中间件挂载   | app.use(router.routes()) | app.use('/', router) |
| 监听端口   | app.listen(3000) | app.listen(3000) |

上表展示了二者的使用区别，从初始化就看出koa语法都是用的新标准。在挂载路由中间件上也有一定的差异性，这是因为二者内部实现机制的不同。其他都是大同小异。

## koa2
koa是一个基于node实现的一个新的web框架，它是由express框架的原班人马打造的。它的特点是优雅、简洁、表达力强、自由度高。它跟express相比，它是一个更轻量的node框架，因为它所有功能都通过插件实现，这种插拔式的架构设计模式，很符合unix哲学。


context对象：koa2应用上下文，封装了request和response。


koa2示例：
```
const Koa = require('koa');
const app = new Koa();

app.use(async ctx => {
  ctx.body = 'Hello World';
});

app.listen(3000);
```
### koa2-websocket

![koa2-websocket逻辑流程](/img/js/koa2-websocket.png)

把WebSocketServer绑定到同一个端口的关键代码是先获取koa创建的http.Server的引用，再根据http.Server创建WebSocketServer：

```
// koa app的listen()方法返回http.Server:
let server = app.listen(3000);

// 创建WebSocketServer:
let wss = new WebSocketServer({
    server: server
});
```

要始终注意，浏览器创建WebSocket时发送的仍然是标准的HTTP请求。无论是WebSocket请求，还是普通HTTP请求，都会被http.Server处理。具体的处理方式则是由koa和WebSocketServer注入的回调函数实现的。`WebSocketServer会首先判断请求是不是WS请求，如果是，它将处理该请求，如果不是，该请求仍由koa处理`。

所以，WS请求会直接由WebSocketServer处理，它根本不会经过koa，koa的任何middleware都没有机会处理该请求。


## express

express示例：
```
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => res.send('Hello World!'))

app.listen(port, () => console.log(`Example app listening on port ${port}!`))
```

# 小结

* github上开源的nodejs/javascript很多经典代码。非常值得借鉴和研究。
  * 多研究express，koa2，ws，socket.io源码。

* nodejs适合开发RTM/RTC接入服务系统，开发效率高。

# 参考链接

- [KOA2框架原理解析和实现](https://juejin.cn/post/6844903709592256525)

- [如何选择正确的Node框架：Express，Koa还是Hapi？](https://blog.fundebug.com/2019/05/10/express-koa-hapi/)

- [https://juejin.cn/post/6844903968041091080](https://juejin.cn/post/68449039680410910800)

- [https://juejin.cn/post/6844903830979624974](https://juejin.cn/post/6844903830979624974)

- [https://www.liaoxuefeng.com/wiki/1022910821149312/1103332447876608](https://www.liaoxuefeng.com/wiki/1022910821149312/1103332447876608)

- [https://zhuanlan.zhihu.com/p/94682749](https://zhuanlan.zhihu.com/p/94682749)