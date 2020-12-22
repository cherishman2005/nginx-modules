# koa2与express比较

* express包含了router；

* koa2更加轻量级，通过middleware中间件插入koa-router功能。

* Koa专注于核心中间件功能，设计显式地利用了async/await使异步代码可读性更高。

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



## express

express示例：
```
const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => res.send('Hello World!'))

app.listen(port, () => console.log(`Example app listening on port ${port}!`))
```

# 参考链接

- [KOA2框架原理解析和实现](https://juejin.cn/post/6844903709592256525)

- [如何选择正确的Node框架：Express，Koa还是Hapi？](https://blog.fundebug.com/2019/05/10/express-koa-hapi/)

- [https://juejin.cn/post/6844903968041091080](https://juejin.cn/post/68449039680410910800)

- [https://juejin.cn/post/6844903830979624974](https://juejin.cn/post/6844903830979624974)