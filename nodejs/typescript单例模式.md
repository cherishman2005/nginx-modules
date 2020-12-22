# 单例模式

传统的单例模式可以用来解决所有代码必须写到 class 中的问题：
```
class Singleton {
  private static instance: Singleton;
  private constructor() {
    // ..
  }

  public static getInstance() {
    if (!Singleton.instance) {
      Singleton.instance = new Singleton();
    }

    return Singleton.instance;
  }

  someMethod() {}
}

let someThing = new Singleton(); // Error: constructor of 'singleton' is private

let instacne = Singleton.getInstance(); // do some thing with the instance
```
然而，如果你不想延迟初始化，你可以使用 namespace 替代：
```
namespace Singleton {
  // .. 其他初始化的代码

  export function someMethod() {}
}

// 使用
Singleton.someMethod();
WARNING
```
单例只是全局的一个别称。

对大部分使用者来说，namespace 可以用模块来替代。
```
// someFile.ts
// ... any one time initialization goes here ...
export function someMethod() {}

// Usage
import { someMethod } from './someFile';
```
# 链接

- [https://jkchao.github.io/typescript-book-chinese/tips/singletonPatern.html](https://jkchao.github.io/typescript-book-chinese/tips/singletonPatern.html)