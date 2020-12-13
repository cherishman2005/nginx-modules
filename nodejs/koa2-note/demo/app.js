const Koa = require('koa');
//const Router = require('koa-router')
const config = require('./config/default.js');
const app = new Koa();

// 加载路由中间件
app.use(require('./routers/routes.js').routes())

app.listen(config.port);

console.log(`listening on port ${config.port}`)
