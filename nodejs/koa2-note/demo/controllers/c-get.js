let controller = {};

controller.home = async (ctx) => {
    let html = 'Hello, koa2'
    ctx.body = html
}

controller.notfound = async (ctx) => {
    ctx.body = '404 page!'
}

controller.helloworld = async (ctx) => {
    ctx.body = 'helloworld page!'
}

module.exports = controller