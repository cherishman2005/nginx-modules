# nodejs

V8堆内存限制

内存在服务端本来就是一个寸土寸金的东西，在 V8 中限制 64 位的机器大约 1.4GB，32 位机器大约为 0.7GB。因此，对于一些大内存的操作需谨慎否则超出 V8 内存限制将会造成进程退出。

## nodejs gc回收

内存泄漏识别

在 Node.js 环境里提供了 process.memoryUsage 方法用来查看当前进程内存使用情况，单位为字节

* rss（resident set size）：RAM 中保存的进程占用的内存部分，包括代码本身、栈、堆。
* heapTotal：堆中总共申请到的内存量。
* heapUsed：堆中目前用到的内存量，判断内存泄漏我们主要以这个字段为准。
* external： V8 引擎内部的 C++ 对象占用的内存。

### global.gc()示例

（1）global.gc();

```
/**
 * 单位为字节格式为 MB 输出
 */
const format = function (bytes) {
    return (bytes / 1024 / 1024).toFixed(2) + ' MB';
};

/**
 * 封装 print 方法输出内存占用信息 
 */
const print = function() {
    const memoryUsage = process.memoryUsage();

    console.log(JSON.stringify({
        rss: format(memoryUsage.rss),
        heapTotal: format(memoryUsage.heapTotal),
        heapUsed: format(memoryUsage.heapUsed),
        external: format(memoryUsage.external),
    }));
}


// example.js
function Quantity(num) {
    if (num) {
        return new Array(num * 1024 * 1024);
    }

    return num;
}

function Fruit(name, quantity) {
    this.name = name
    this.quantity = new Quantity(quantity)
}

let apple = new Fruit('apple');
print();
let banana = new Fruit('banana', 20);
print();

delete banana;
banana = null;
print();

global.gc();
print();
```

（2）运行
手动在代码中操作 GC (不推荐)
node --expose-gc example.js
运行示例：
```
{"rss":"30.55 MB","heapTotal":"6.23 MB","heapUsed":"3.64 MB","external":"0.01 MB"}
{"rss":"190.56 MB","heapTotal":"166.25 MB","heapUsed":"163.68 MB","external":"0.01 MB"}
{"rss":"190.56 MB","heapTotal":"166.25 MB","heapUsed":"163.69 MB","external":"0.01 MB"}
{"rss":"30.83 MB","heapTotal":"8.73 MB","heapUsed":"3.23 MB","external":"0.01 MB"}
```

【疑问】new Array()开辟了20M，为啥新增160M？


### overflow

（1）overflow.js
```
// overflow.js
const format = function (bytes) {
    return (bytes / 1024 / 1024).toFixed(2) + ' MB';
};

const print = function() {
    const memoryUsage = process.memoryUsage();
    console.log(`heapTotal: ${format(memoryUsage.heapTotal)}, heapUsed: ${format(memoryUsage.heapUsed)}`);
}

const total = [];
setInterval(function() {
    total.push(new Array(20 * 1024 * 1024)); // 大内存占用
    print();
}, 1000)
```

（2）运行
```
heapTotal: 166.25 MB, heapUsed: 163.67 MB
heapTotal: 326.26 MB, heapUsed: 323.72 MB
heapTotal: 486.27 MB, heapUsed: 483.73 MB
heapTotal: 648.78 MB, heapUsed: 643.26 MB
heapTotal: 808.79 MB, heapUsed: 803.26 MB
heapTotal: 968.80 MB, heapUsed: 963.27 MB
heapTotal: 1128.82 MB, heapUsed: 1123.23 MB
heapTotal: 1288.83 MB, heapUsed: 1283.23 MB

<--- Last few GCs --->

[13057:0x2c80070]    10270 ms: Mark-sweep 1283.2 (1286.8) -> 1283.2 (1285.8) MB, 426.9 / 0.0 ms  (average mu = 0.456, current mu = 0.000) last resort GC in old space requested
[13057:0x2c80070]    10725 ms: Mark-sweep 1283.2 (1285.8) -> 1283.2 (1285.8) MB, 454.8 / 0.0 ms  (average mu = 0.284, current mu = 0.000) last resort GC in old space requested


<--- JS stacktrace --->

==== JS stack trace =========================================

    0: ExitFrame [pc: 0x151b50f5be1d]
    1: StubFrame [pc: 0x151b50f8bb71]
Security context: 0x08fff9e9e6e1 <JSObject>
    2: _onTimeout [0x27725c34ef41] [/home/zhangbiwu/nginx/openresty/nginx/html/nodejs/overflow.js:13] [bytecode=0x1e72739da149 offset=28](this=0x27725c34efe1 <Timeout map = 0x2f3b9d251fb9>)
    3: ontimeout(aka ontimeout) [0x27725c369d69] [timers.js:436] [bytecode=0x1e72739d9e29 offset=83](this=0x0ec5389026f1 <undefined>...

FATAL ERROR: CALL_AND_RETRY_LAST Allocation failed - JavaScript heap out of memory
```

内存泄露。

# 小结

* 全局变量

  * 未声明的变量或挂在全局 global 下的变量不会自动回收，将会常驻内存直到进程退出才会被释放，除非通过 delete 或 重新赋值为 undefined/null 解决之间的引用关系，才会被回收。关于全局变量上面举的几个例子中也有说明。

* 闭包

* 慎将内存做为缓存
  * 启动多个进程或部署在多台机器会造成每个进程都会保存一份，显然是资源的浪费，最好是通过 Redis 做共享。

* 模块私有变量内存永驻

* 事件重复监听
  在 Node.js 中对一个事件重复监听则会报如下错误，实际上使用的 EventEmitter 类，该类包含一个 listeners 数组，默认为 10 个监听器超出这个数则会报警如下所示，用于发现内存泄漏，也可以通过 emitter.setMaxListeners() 方法为指定的 EventEmitter 实例修改限制。

* 其他注意事项
  * 在使用定时器 setInterval 时记的使用对应的 clearInterval 进行清除，因为 setInterval 执行完之后会返回一个值且不会自动释放。