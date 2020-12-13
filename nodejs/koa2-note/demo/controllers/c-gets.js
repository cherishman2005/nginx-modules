
exports.home = async (ctx) => {
    let html = 'Hello, koa2'
    ctx.body = html
}

exports.notfound = async (ctx) => {
    ctx.body = '404 page!'
}

exports.helloworld = async (ctx) => {
    ctx.body = 'helloworld page!'
}

exports.imageList = async (ctx) => {
    let images = [];
    for (let i = 0; i < 10; i++) {
        let data = { id: i, name: '阿武', imgUrl: 'http://127.0.0.1/helloworld'};
        images.push(data);
    }
    ctx.body = JSON.stringify(images);
}