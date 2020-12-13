const router = require('koa-router')();
const controller = require('../controllers/c-gets.js');

router.get('/', controller.home)

router.get('/404', controller.notfound)

router.get('/helloworld', controller.helloworld)

router.get('/image-list', controller.imageList)

module.exports = router