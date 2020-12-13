const router = require('koa-router')();
const controller = require('../controllers/c-get.js');

router.get('/', controller.home)

router.get('/404', controller.notfound)

router.get('/helloworld', controller.helloworld)

module.exports = router