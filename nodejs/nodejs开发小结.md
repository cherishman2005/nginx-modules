# nodejs

nodejs和前端js的区别：

| 区别                      | nodejs                                  | 前端js                                  |
|:--------------------------|:--------------------------------------|:--------------------------------------|
|     | 基于服务端的 javascript                      | 基于浏览器端的 javascript                    |
|     | nodejs尽量不要打包，采用模块设计              | webpack等工具打包                            |

JavaScript是一门编程语言（脚本语言），而Node.js是一个平台。

* nodejs

  nodejs是一个基于Google V8引擎（js引擎的某一种）构建的JavaScript运行环境。

* 前端js

  基于浏览器端的 javascript。


## nodejs开发小结

### 注意内存泄露问题；

* 性能管理： CPU密集型功能采用c module；

### nodejs多进程原理？

#### nodejs后端服务多实例（多进程）
NodeJS程序是以单进程形式运行，32位机器上最多也只有1GB内存的实用权限（在64位机器上最大的内存权限扩大到1.7GB）。而目前绝大部分线上服务器的CPU都是多核并且至少16GB起，如此以来Node程序便无法充分发挥机器的潜力。同时NodeJS自己也意识到了这一点，所以它允许程序创建多个子进程用于运行多个实例。

有很多解决办法：
（1）一条机器启动多个nodejs进程，分别对应不同的端口；再用nginx进行代理；
（2）nodejs采用master-worker进程模式，形如nginx。—— 如采用pm2进行部署。
（3）采用nginx-Cluster模块。

负载均衡策略多借鉴nginx的设计思想；

# FAQ

## 64bit数据的支持。

为什么采用最多53位整型，而不是64位整型？这是因为考虑到大部分应用程序是Web应用，如果要和JavaScript打交道，由于JavaScript支持的最大整型就是53位，超过这个位数，JavaScript将丢失精度。因此，使用53位整数可以直接由JavaScript读取，而超过53位时，就必须转换成字符串才能保证JavaScript处理正确，这会给API接口带来额外的复杂度。

## MVC模式设计

nodejs后端开发经典代码：

- [https://github.com/bezkoder/nodejs-express-mysql](https://github.com/bezkoder/nodejs-express-mysql)

MVC模块化设计，便于维护。

## nodejs程序设计要考虑的问题。

假设现在需要你用NodeJS搭建一个http服务，借助express框架用不到10行的代码完成这项工作。不能说这么做是错的，但这样简易的程序是脆弱的：一旦部署上线之后，可能瞬间就被大量涌入的请求击垮，更不要提各种潜在的漏洞危险。退一步说，即使线上程序经过了这一关考验，

* 如果你要更新程序怎么办？不得不让用户中断访问一段时间？
  —— 实质做到微服务，在前面加上API网关（如nginx）就好办了。

后端服务务必要满足两条特性：
* 能容错（Fault tolerant）
* 可扩展（Scalability）

## nodejs经典框架

* Express

* KOA

  KOA框架由Express原班人马打造，它的核心是ES6的generator。KOA使用generator来实现中间件的流程控制，使用try/catch来增强异常处理，同时在KOA框架中你再也看不到复杂的callback回调了。KOA框架本身非常小，只打包了一些必要的功能，但是它本身通过良好的模块化组织，让开发人员可以按照自己的想法来实现一个扩展性非常好的应用。

* Egg

  Egg已经被用在阿里多条产品线（包括蚂蚁）上，已经证明它的安全和可靠性，可以放心用。



# 参考链接

- [https://leokongwq.github.io/2016/11/08/nodejs-gc.html](https://leokongwq.github.io/2016/11/08/nodejs-gc.html)

- [https://www.techug.com/post/10-tips-make-node-js-web-app-faster.html](https://www.techug.com/post/10-tips-make-node-js-web-app-faster.html)

- [https://www.zhihu.com/question/313414600](https://www.zhihu.com/question/313414600)

- [https://tech.meituan.com/2017/04/21/mt-leaf.html](https://tech.meituan.com/2017/04/21/mt-leaf.html)

- [https://www.liaoxuefeng.com/article/1280526512029729](https://www.liaoxuefeng.com/article/1280526512029729)

- [https://www.cnblogs.com/tugenhua0707/p/10776725.html](https://www.cnblogs.com/tugenhua0707/p/10776725.html)

- [https://github.com/bezkoder/nodejs-express-mysql](https://github.com/bezkoder/nodejs-express-mysql)

- [https://juejin.im/entry/6844903698880004110](https://juejin.im/entry/6844903698880004110)

- [node.js后端框架介绍](https://zhuanlan.zhihu.com/p/133666957)