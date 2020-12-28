# nodejs与go比较 

nodejs和go均擅长做web服务器

|        |   nodejs  |  go |
| :---------- | :------ | :------ |
| 内存限制      | 64 位操作系统约为 1.4G<BR>32 位操作系统约为 0.7G |  |



## nodejs和golang协程的不同

nodejs和golang都是支持协程的，从表现上来看，nodejs对于协程的支持在于async/await，golang对协程的支持在于goroutine。关于协程的话题，简单来说，可以看作是非抢占式的轻量级线程。

**nodejs await协程**
```
function sleep(ms){
  return new Promise((resolve,reject) => setTimeout(
     ()=>resolve(),ms
  ))
}

async function hello(){
  await sleep(3000)
  console.log("hello")
};

async function test(){
  await sleep(3000)
  console.log("world")
};

hello();
test();
```


**go协程**
```
package main

import (
  "fmt"
  "time"
)

func deferPrint(str string){
  time.Sleep(time.Second*2)
  fmt.Println(str)
}

func main(){
  go deferPrint("hello")
  go deferPrint("world")
  //如果主协程不阻塞，永远不会切换

  time.Sleep(time.Second*4)
}
```

### 线程支持

golang之所以要支持锁协程，我想是为了多线程支持。golang中可以启用多个线程并行执行相同数量的协程。

nodejs受限于v8的isolate机制，只能跑在单线程中。所有代码无法并行执行，无法处理计算密集型应用场景。

### 切换机制

nodejs使用await阻塞协程，手动切换线程控制权，node的协程是c++控制的，c++里写了这个函数可以被推入事件队列就能够用promise封装成协程

golang在协程阻塞时自动切换协程，所以在写golang的时候所有的代码可以都写同步代码，然后用go关键字去调用，golang的协程是自己规定的，所有函数在阻塞时都必须切换线程控制权

### 取返回值

* nodejs中async函数是能直接返回值的
* golang只能传递一个引用的channel

golang类似与小c++，最大的亮点就是使用协程管理多线程。

# 参考链接

- [https://www.cnblogs.com/kazetotori/p/6896653.html](https://www.cnblogs.com/kazetotori/p/6896653.html)

- [https://zhuanlan.zhihu.com/p/59295820](https://zhuanlan.zhihu.com/p/59295820)

- [https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/)

- [《深入浅出Node.js》-内存控制](https://lz5z.com/%E6%B7%B1%E5%85%A5%E6%B5%85%E5%87%BANode-js-%E5%86%85%E5%AD%98%E6%8E%A7%E5%88%B6/)
